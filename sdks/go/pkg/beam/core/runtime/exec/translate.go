// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/coder"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/graphx"
	v1pb "github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/graphx/v1"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/protox"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/stringx"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/internal/errors"
	fnpb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/fnexecution_v1"
	pipepb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/pipeline_v1"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// TODO(lostluck): 2018/05/28 Extract these from the canonical enums in beam_runner_api.proto
const (
	urnDataSource                          = "beam:runner:source:v1"
	urnDataSink                            = "beam:runner:sink:v1"
	urnPerKeyCombinePre                    = "beam:transform:combine_per_key_precombine:v1"
	urnPerKeyCombineMerge                  = "beam:transform:combine_per_key_merge_accumulators:v1"
	urnPerKeyCombineExtract                = "beam:transform:combine_per_key_extract_outputs:v1"
	urnPerKeyCombineConvert                = "beam:transform:combine_per_key_convert_to_accumulators:v1"
	urnPairWithRestriction                 = "beam:transform:sdf_pair_with_restriction:v1"
	urnSplitAndSizeRestrictions            = "beam:transform:sdf_split_and_size_restrictions:v1"
	urnProcessSizedElementsAndRestrictions = "beam:transform:sdf_process_sized_element_and_restrictions:v1"
)

// UnmarshalPlan converts a model bundle descriptor into an execution Plan.
func UnmarshalPlan(desc *fnpb.ProcessBundleDescriptor) (*Plan, error) {
	b, err := newBuilder(desc)
	if err != nil {
		return nil, err
	}
	for id, transform := range desc.GetTransforms() {
		if transform.GetSpec().GetUrn() != urnDataSource {
			continue
		}
		if len(transform.GetOutputs()) != 1 {
			return nil, errors.Errorf("expected one output from DataSource, got %v", transform.GetOutputs())
		}

		port, cid, err := unmarshalPort(transform.GetSpec().GetPayload())
		if err != nil {
			return nil, err
		}

		u := &DataSource{UID: b.idgen.New()}
		u.Coder, err = b.coders.Coder(cid) // Expected to be windowed coder
		if err != nil {
			return nil, err
		}
		if !coder.IsW(u.Coder) {
			return nil, errors.Errorf("unwindowed coder %v on DataSource %v: %v", cid, id, u.Coder)
		}

		// There's only a single pair in this map, but a for loop range statement
		// is the easiest way to extract it, so this loop will iterate only once.
		for key, pid := range transform.GetOutputs() {
			u.SID = StreamID{PtransformID: id, Port: port}
			u.Name = key

			u.Out, err = b.makePCollection(pid)
			if err != nil {
				return nil, err
			}
			// Elide the PCollection Node for DataSources
			// DataSources can get byte samples directly, and can handle CoGBKs.
			// Copying the PCollection here is fine, as the PCollection will never
			// have used it's mutex yet.
			u.PCol = *u.Out.(*PCollection)
			u.Out = u.PCol.Out
			b.units = b.units[:len(b.units)-1]
		}

		b.units = append(b.units, u)
	}
	return b.build()
}

type builder struct {
	desc   *fnpb.ProcessBundleDescriptor
	coders *graphx.CoderUnmarshaller

	prev map[string]int      // PCollectionID -> #incoming
	succ map[string][]linkID // PCollectionID -> []linkID

	windowing map[string]*window.WindowingStrategy
	nodes     map[string]*PCollection // PCollectionID -> Node (cache)
	links     map[linkID]Node         // linkID -> Node (cache)

	units []Unit // result
	idgen *GenID
}

// linkID represents an incoming data link to an Node.
type linkID struct {
	to    string // TransformID
	input int    // input index. If > 0, it's a side input.
}

func newBuilder(desc *fnpb.ProcessBundleDescriptor) (*builder, error) {
	// Preprocess graph structure to allow insertion of Multiplex,
	// Flatten and Discard.

	prev := make(map[string]int)      // PCollectionID -> #incoming
	succ := make(map[string][]linkID) // PCollectionID -> []linkID

	for id, transform := range desc.GetTransforms() {
		if len(transform.GetSubtransforms()) > 0 {
			continue // ignore composites
		}

		input := unmarshalKeyedValues(transform.GetInputs())
		for i, from := range input {
			succ[from] = append(succ[from], linkID{id, i})
		}
		output := unmarshalKeyedValues(transform.GetOutputs())
		for _, to := range output {
			prev[to]++
		}
	}

	b := &builder{
		desc:   desc,
		coders: graphx.NewCoderUnmarshaller(desc.GetCoders()),

		prev: prev,
		succ: succ,

		windowing: make(map[string]*window.WindowingStrategy),
		nodes:     make(map[string]*PCollection),
		links:     make(map[linkID]Node),

		idgen: &GenID{},
	}
	return b, nil
}

func (b *builder) build() (*Plan, error) {
	return NewPlan(b.desc.GetId(), b.units)
}

func (b *builder) makeWindowingStrategy(id string) (*window.WindowingStrategy, error) {
	if w, exists := b.windowing[id]; exists {
		return w, nil
	}

	ws, ok := b.desc.GetWindowingStrategies()[id]
	if !ok {
		return nil, errors.Errorf("windowing strategy %v not found", id)
	}
	wfn, err := unmarshalWindowFn(ws.GetWindowFn())
	if err != nil {
		return nil, err
	}
	w := &window.WindowingStrategy{Fn: wfn}
	b.windowing[id] = w
	return w, nil
}

func unmarshalWindowFn(wfn *pipepb.FunctionSpec) (*window.Fn, error) {
	switch urn := wfn.GetUrn(); urn {
	case graphx.URNGlobalWindowsWindowFn:
		return window.NewGlobalWindows(), nil

	case graphx.URNFixedWindowsWindowFn:
		var payload pipepb.FixedWindowsPayload
		if err := proto.Unmarshal(wfn.GetPayload(), &payload); err != nil {
			return nil, err
		}
		size, err := ptypes.Duration(payload.GetSize())
		if err != nil {
			return nil, err
		}
		return window.NewFixedWindows(size), nil

	case graphx.URNSlidingWindowsWindowFn:
		var payload pipepb.SlidingWindowsPayload
		if err := proto.Unmarshal(wfn.GetPayload(), &payload); err != nil {
			return nil, err
		}
		period, err := ptypes.Duration(payload.GetPeriod())
		if err != nil {
			return nil, err
		}
		size, err := ptypes.Duration(payload.GetSize())
		if err != nil {
			return nil, err
		}
		return window.NewSlidingWindows(period, size), nil

	case graphx.URNSessionsWindowFn:
		var payload pipepb.SessionWindowsPayload
		if err := proto.Unmarshal(wfn.GetPayload(), &payload); err != nil {
			return nil, err
		}
		gap, err := ptypes.Duration(payload.GetGapSize())
		if err != nil {
			return nil, err
		}
		return window.NewSessions(gap), nil

	default:
		return nil, errors.Errorf("unsupported window type: %v", urn)
	}
}

func unmarshalAndMakeWindowMapping(wmfn *pipepb.FunctionSpec) (WindowMapper, error) {
	switch urn := wmfn.GetUrn(); urn {
	case graphx.URNWindowMappingGlobal:
		return &windowMapper{wfn: window.NewGlobalWindows()}, nil
	case graphx.URNWindowMappingFixed:
		var payload pipepb.FixedWindowsPayload
		if err := proto.Unmarshal(wmfn.GetPayload(), &payload); err != nil {
			return nil, err
		}
		size, err := ptypes.Duration(payload.GetSize())
		if err != nil {
			return nil, err
		}
		return &windowMapper{wfn: window.NewFixedWindows(size)}, nil
	case graphx.URNWindowMappingSliding:
		var payload pipepb.SlidingWindowsPayload
		if err := proto.Unmarshal(wmfn.GetPayload(), &payload); err != nil {
			return nil, err
		}
		period, err := ptypes.Duration(payload.GetPeriod())
		if err != nil {
			return nil, err
		}
		size, err := ptypes.Duration(payload.GetSize())
		if err != nil {
			return nil, err
		}
		return &windowMapper{wfn: window.NewSlidingWindows(period, size)}, nil
	default:
		return nil, fmt.Errorf("unsupported window mapping fn URN %v", urn)
	}
}

func (b *builder) makePCollections(out []string) ([]Node, error) {
	var ret []Node
	for _, o := range out {
		n, err := b.makePCollection(o)
		if err != nil {
			return nil, err
		}
		// This is the cleanest place to do this check and filtering,
		// since DataSinks don't know their inputs, due to the construction
		// call stack.
		// A Source->Sink is both uncommon and inefficent, with the Source eliding the
		// collection anyway.
		// TODO[BEAM-6374): Properly handle the multiplex and flatten cases.
		// Right now we just stop datasink collection.
		switch out := n.Out.(type) {
		case *DataSink:
			// We don't remove the PCollection from units here, since we
			// want to ensure it's included in snapshots.
			out.PCol = n
			ret = append(ret, out)
		default:
			ret = append(ret, n)
		}
	}
	return ret, nil
}

func (b *builder) makeCoderForPCollection(id string) (*coder.Coder, *coder.WindowCoder, error) {
	col, ok := b.desc.GetPcollections()[id]
	if !ok {
		return nil, nil, errors.Errorf("pcollection %v not found", id)
	}
	c, err := b.coders.Coder(col.CoderId)
	if err != nil {
		return nil, nil, err
	}
	if coder.IsW(c) {
		// TODO(herohde) 3/16/2018: remove potential WindowedValue from Dataflow.
		// However, windowing strategies are not yet passed through, so the main
		// path always gives us GlobalWindows.

		return coder.SkipW(c), c.Window, nil
	}

	ws, ok := b.desc.GetWindowingStrategies()[col.GetWindowingStrategyId()]
	if !ok {
		return nil, nil, errors.Errorf("windowing strategy %v not found", id)
	}
	wc, err := b.coders.WindowCoder(ws.GetWindowCoderId())
	if err != nil {
		return nil, nil, err
	}
	return c, wc, nil
}

func (b *builder) makePCollection(id string) (*PCollection, error) {
	if n, exists := b.nodes[id]; exists {
		return n, nil
	}

	list := b.succ[id]

	var u Node
	switch len(list) {
	case 0:
		// Discard.

		u = &Discard{UID: b.idgen.New()}

	case 1:
		out, err := b.makeLink(id, list[0])
		if err != nil {
			return nil, err
		}
		return b.newPCollectionNode(id, out)
	default:
		// Multiplex.

		out, err := b.makeLinks(id, list)
		if err != nil {
			return nil, err
		}
		u = &Multiplex{UID: b.idgen.New(), Out: out}
	}

	if count := b.prev[id]; count > 1 {
		// Guard node with Flatten, if needed.

		b.units = append(b.units, u)
		u = &Flatten{UID: b.idgen.New(), N: count, Out: u}
	}
	b.units = append(b.units, u)
	return b.newPCollectionNode(id, u)
}

func (b *builder) newPCollectionNode(id string, out Node) (*PCollection, error) {
	ec, _, err := b.makeCoderForPCollection(id)
	if err != nil {
		return nil, err
	}
	u := &PCollection{UID: b.idgen.New(), Out: out, PColID: id, Coder: ec, Seed: rand.Int63()}
	b.nodes[id] = u
	b.units = append(b.units, u)
	return u, nil
}

func (b *builder) makeLinks(from string, ids []linkID) ([]Node, error) {
	var ret []Node
	for _, id := range ids {
		n, err := b.makeLink(from, id)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}
	return ret, nil
}

func (b *builder) makeLink(from string, id linkID) (Node, error) {
	if n, ok := b.links[id]; ok {
		return n, nil
	}

	// Process all incoming links for the edge and cache them. It thus doesn't matter
	// which exact link triggers the Node generation. The link caching is only needed
	// to process ParDo side inputs.

	transform := b.desc.GetTransforms()[id.to]
	urn := transform.GetSpec().GetUrn()
	payload := transform.GetSpec().GetPayload()

	// TODO(herohde) 1/25/2018: do we need to handle composites?

	out, err := b.makePCollections(unmarshalKeyedValues(transform.GetOutputs()))
	if err != nil {
		return nil, err
	}

	var u Node
	switch urn {
	case graphx.URNParDo,
		urnPerKeyCombinePre,
		urnPerKeyCombineMerge,
		urnPerKeyCombineExtract,
		urnPerKeyCombineConvert,
		urnPairWithRestriction,
		urnSplitAndSizeRestrictions,
		urnProcessSizedElementsAndRestrictions:
		var data string
		var sides map[string]*pipepb.SideInput
		switch urn {
		case graphx.URNParDo,
			urnPairWithRestriction,
			urnSplitAndSizeRestrictions,
			urnProcessSizedElementsAndRestrictions:
			var pardo pipepb.ParDoPayload
			if err := proto.Unmarshal(payload, &pardo); err != nil {
				return nil, errors.Wrapf(err, "invalid ParDo payload for %v", transform)
			}
			data = string(pardo.GetDoFn().GetPayload())
			sides = pardo.GetSideInputs()
		case urnPerKeyCombinePre, urnPerKeyCombineMerge, urnPerKeyCombineExtract, urnPerKeyCombineConvert:
			var cmb pipepb.CombinePayload
			if err := proto.Unmarshal(payload, &cmb); err != nil {
				return nil, errors.Wrapf(err, "invalid CombinePayload payload for %v", transform)
			}
			data = string(cmb.GetCombineFn().GetPayload())
		default:
			// TODO(herohde) 12/4/2017: we see DoFns directly with Dataflow. Handle that
			// case here, for now, so that the harness can use this logic.

			data = string(payload)
		}

		// TODO(herohde) 1/28/2018: Once Dataflow's fully off the old way,
		// we can simply switch on the ParDo DoFn URN directly.

		var tp v1pb.TransformPayload
		if err := protox.DecodeBase64(data, &tp); err != nil {
			return nil, errors.Wrapf(err, "invalid transform payload for %v", transform)
		}

		switch tpUrn := tp.GetUrn(); tpUrn {
		case graphx.URNDoFn:
			op, fn, _, in, _, err := graphx.DecodeMultiEdge(tp.GetEdge())
			if err != nil {
				return nil, err
			}

			switch op {
			case graph.ParDo:
				dofn, err := graph.AsDoFn(fn, graph.MainUnknown)
				if err != nil {
					return nil, err
				}
				switch urn {
				case urnPairWithRestriction:
					u = &PairWithRestriction{UID: b.idgen.New(), Fn: dofn, Out: out[0]}
				case urnSplitAndSizeRestrictions:
					u = &SplitAndSizeRestrictions{UID: b.idgen.New(), Fn: dofn, Out: out[0]}
				default:
					n := &ParDo{UID: b.idgen.New(), Fn: dofn, Inbound: in, Out: out}
					n.PID = transform.GetUniqueName()

					input := unmarshalKeyedValues(transform.GetInputs())
					for i := 1; i < len(input); i++ {
						// TODO(BEAM-3305) Handle ViewFns for side inputs

						ec, wc, err := b.makeCoderForPCollection(input[i])
						if err != nil {
							return nil, err
						}

						sidePB, ok := sides[indexToInputId(i)]
						if !ok {
							return nil, fmt.Errorf("missing side input info for collection %v", input[i])
						}

						mapper, err := unmarshalAndMakeWindowMapping(sidePB.GetWindowMappingFn())
						if err != nil {
							return nil, err
						}

						sid := StreamID{
							Port:         Port{URL: b.desc.GetStateApiServiceDescriptor().GetUrl()},
							PtransformID: id.to,
						}
						sideInputID := fmt.Sprintf("i%v", i) // SideInputID (= local id, "iN")
						side := NewSideInputAdapter(sid, sideInputID, coder.NewW(ec, wc), mapper)
						n.Side = append(n.Side, side)
					}
					u = n
					if urn == urnProcessSizedElementsAndRestrictions {
						u = &ProcessSizedElementsAndRestrictions{PDo: n, TfId: id.to}
					} else if dofn.IsSplittable() {
						u = &SdfFallback{PDo: n}
					}
				}

			case graph.Combine:
				cn := &Combine{UID: b.idgen.New(), Out: out[0]}
				cn.Fn, err = graph.AsCombineFn(fn)
				if err != nil {
					return nil, err
				}
				cn.UsesKey = typex.IsKV(in[0].Type)

				cn.PID = transform.GetUniqueName()

				switch urn {
				case urnPerKeyCombinePre:
					inputs := unmarshalKeyedValues(transform.GetInputs())
					if len(inputs) != 1 {
						return nil, errors.Errorf("unexpected sideinput to combine: got %d, want 1", len(inputs))
					}
					ec, wc, err := b.makeCoderForPCollection(inputs[0])
					if err != nil {
						return nil, err
					}
					if !coder.IsKV(ec) {
						return nil, errors.Errorf("unexpected non-KV coder PCollection input to combine: %v", ec)
					}
					u = &LiftedCombine{Combine: cn, KeyCoder: ec.Components[0], WindowCoder: wc}
				case urnPerKeyCombineMerge:
					ma := &MergeAccumulators{Combine: cn}
					if eo, ok := ma.Out.(*PCollection).Out.(*ExtractOutput); ok {
						// Strip PCollections from between MergeAccumulators and ExtractOutputs
						// as it's a synthetic PCollection.
						b.units = b.units[:len(b.units)-1]
						ma.Out = eo
					}
					u = ma
				case urnPerKeyCombineExtract:
					u = &ExtractOutput{Combine: cn}
				case urnPerKeyCombineConvert:
					u = &ConvertToAccumulators{Combine: cn}
				default: // For unlifted combines
					u = cn
				}
			default:
				panic(fmt.Sprintf("Opcode should be one of ParDo or Combine, but it is: %v", op))
			}

		case graphx.URNIterableSideInputKey:
			u = &FixedKey{UID: b.idgen.New(), Key: []byte(iterableSideInputKey), Out: out[0]}

		case graphx.URNInject:
			c, _, err := b.makeCoderForPCollection(from)
			if err != nil {
				return nil, err
			}
			if !coder.IsKV(c) {
				return nil, errors.Errorf("unexpected inject coder: %v", c)
			}
			valCoder := c.Components[1]
			// JIRA BEAM-12438 - an extra LP coder can get added here, but isn't added
			// on decode. Strip them until we get a better fix.
			if valCoder.Kind == coder.LP {
				// strip unexpected length prefix coder.
				valCoder = valCoder.Components[0]
			}
			u = &Inject{UID: b.idgen.New(), N: (int)(tp.GetInject().GetN()), ValueEncoder: MakeElementEncoder(valCoder), Out: out[0]}

		case graphx.URNExpand:
			var pid string
			for _, id := range transform.GetOutputs() {
				pid = id
			}
			c, _, err := b.makeCoderForPCollection(pid)
			if err != nil {
				return nil, err
			}
			if !coder.IsCoGBK(c) {
				return nil, errors.Errorf("unexpected expand coder: %v", c)
			}

			var decoders []ElementDecoder
			for _, dc := range c.Components[1:] {
				decoders = append(decoders, MakeElementDecoder(dc))
			}
			// Strip PCollections from Expand nodes, as CoGBK metrics are handled by
			// the DataSource that preceeds them.
			trueOut := out[0].(*PCollection).Out
			b.units = b.units[:len(b.units)-1]
			u = &Expand{UID: b.idgen.New(), ValueDecoders: decoders, Out: trueOut}

		case graphx.URNReshuffleInput:
			_, w, err := b.makeCoderForPCollection(from)
			if err != nil {
				return nil, err
			}
			preservedCoderID := tp.GetReshuffle().GetCoderId()
			pc, err := unmarshalReshuffleCoders(preservedCoderID, tp.GetReshuffle().GetCoderPayloads())
			if err != nil {
				return nil, err
			}
			u = &ReshuffleInput{UID: b.idgen.New(), Seed: rand.Int63(), Coder: coder.NewW(pc, w), Out: out[0]}

		case graphx.URNReshuffleOutput:
			var pid string
			// There's only one output PCollection, and iterating through the map
			// is the only way to extract it.
			for _, id := range transform.GetOutputs() {
				pid = id
			}
			_, w, err := b.makeCoderForPCollection(pid)
			if err != nil {
				return nil, err
			}
			preservedCoderID := tp.GetReshuffle().GetCoderId()
			pc, err := unmarshalReshuffleCoders(preservedCoderID, tp.GetReshuffle().GetCoderPayloads())
			if err != nil {
				return nil, err
			}
			u = &ReshuffleOutput{UID: b.idgen.New(), Coder: coder.NewW(pc, w), Out: out[0]}

		default:
			return nil, errors.Errorf("unexpected payload: %v", &tp)
		}

	case graphx.URNWindow:
		var wp pipepb.WindowIntoPayload
		if err := proto.Unmarshal(payload, &wp); err != nil {
			return nil, errors.Wrapf(err, "invalid WindowInto payload for %v", transform)
		}
		wfn, err := unmarshalWindowFn(wp.GetWindowFn())
		if err != nil {
			return nil, err
		}
		u = &WindowInto{UID: b.idgen.New(), Fn: wfn, Out: out[0]}

	case graphx.URNFlatten:
		u = &Flatten{UID: b.idgen.New(), N: len(transform.Inputs), Out: out[0]}

		// Use the same flatten instance for all the inputs links to this transform.
		for i := 0; i < len(transform.Inputs); i++ {
			b.links[linkID{id.to, i}] = u
		}

	case urnDataSink:
		port, cid, err := unmarshalPort(payload)
		if err != nil {
			return nil, err
		}

		sink := &DataSink{UID: b.idgen.New()}
		sink.SID = StreamID{PtransformID: id.to, Port: port}
		sink.Coder, err = b.coders.Coder(cid) // Expected to be windowed coder
		if err != nil {
			return nil, err
		}
		if !coder.IsW(sink.Coder) {
			return nil, errors.Errorf("unwindowed coder %v on DataSink %v: %v", cid, id, sink.Coder)
		}
		u = sink

	default:
		panic(fmt.Sprintf("Unexpected transform URN: %v", urn))
	}

	b.links[id] = u
	b.units = append(b.units, u)
	return u, nil
}

func unmarshalReshuffleCoders(mainID string, payloads map[string][]byte) (*coder.Coder, error) {
	m := map[string]*pipepb.Coder{}
	for id, v := range payloads {
		pc := &pipepb.Coder{}
		if err := proto.Unmarshal(v, pc); err != nil {
			return nil, err
		}
		m[id] = pc
	}
	um := graphx.NewCoderUnmarshaller(m)
	return um.Coder(mainID)
}

// unmarshalKeyedValues converts a map {"i1": "b", ""i0": "a"} into an ordered list of
// of values: {"a", "b"}. If the keys are not in the expected format, the returned
// list does not guarantee any order, but will respect ordered values.
func unmarshalKeyedValues(m map[string]string) []string {
	if len(m) == 0 {
		return nil
	}
	if len(m) == 1 && stringx.Keys(m)[0] == "bogus" {
		return nil // Ignore special bogus node for legacy Dataflow.
	}

	// (1) Compute index. If generated by the marshaller, we have
	// a "iN" name that directly indicates the position.

	ordered := make(map[int]string)
	var unordered []string

	for key := range m {
		if i, err := inputIdToIndex(key); err == nil {
			if i < len(m) {
				ordered[i] = key
				continue
			} // else: out-of-range index.
		} // else: not in "iN" form.

		unordered = append(unordered, key)
	}

	// (2) Impose order, to the extent present, on values.

	ret := make([]string, len(m))
	k := 0
	for i := 0; i < len(ret); i++ {
		if key, ok := ordered[i]; ok {
			ret[i] = m[key]
		} else {
			ret[i] = m[unordered[k]]
			k++
		}
	}
	return ret
}

// inputIdToIndex converts a local input ID for a transform into an index. Use
// this to avoid relying on format details for input IDs.
//
// Currently, expects IDs in the format "iN" where N is the index. If the ID is
// in an invalid form, returns an error.
func inputIdToIndex(id string) (int, error) {
	if !strings.HasPrefix(id, "i") {
		return 0, errors.New("invalid input ID format")
	}
	return strconv.Atoi(strings.TrimPrefix(id, "i"))
}

// indexToInputId converts an index into a local input ID for a transform. Use
// this to avoid relying on format details for input IDs.
func indexToInputId(i int) string {
	return "i" + strconv.Itoa(i)
}

func unmarshalPort(data []byte) (Port, string, error) {
	var port fnpb.RemoteGrpcPort
	if err := proto.Unmarshal(data, &port); err != nil {
		return Port{}, "", err
	}
	return Port{
		URL: port.GetApiServiceDescriptor().GetUrl(),
	}, port.CoderId, nil
}
