// File generated by specialize. Do not edit.

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

package reflectx

import (
	"fmt"
	"reflect"
)

// Generated arity-specialized Func implementations to avoid runtime temporary
// slices. Code that knows the arity can potentially avoid that overhead in
// addition to avoiding the reflection call overhead.
//
// We force FuncXxY to implement Func as well to have a single registry. Concrete
// Func implementations should implement the most specific supported variant. This
// also allows more specific intermediate interfaces, such as Func2xbool, to be added
// later.

type Func0x0 interface {
	Func
	Call0x0()
}

type shimFunc0x0 struct {
	inner Func
}

func (c *shimFunc0x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc0x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc0x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc0x0) Call0x0() {
	ret := c.inner.Call([]interface{}{})
	_ = ret
	return
}

func ToFunc0x0(c Func) Func0x0 {
	if c.Type().NumIn() != 0 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 0 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func0x0); ok {
		return sc
	}
	return &shimFunc0x0{inner: c}
}

func MakeFunc0x0(fn interface{}) Func0x0 {
	return ToFunc0x0(MakeFunc(fn))
}

type Func0x1 interface {
	Func
	Call0x1() interface{}
}

type shimFunc0x1 struct {
	inner Func
}

func (c *shimFunc0x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc0x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc0x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc0x1) Call0x1() interface{} {
	ret := c.inner.Call([]interface{}{})
	_ = ret
	return ret[0]
}

func ToFunc0x1(c Func) Func0x1 {
	if c.Type().NumIn() != 0 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 0 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func0x1); ok {
		return sc
	}
	return &shimFunc0x1{inner: c}
}

func MakeFunc0x1(fn interface{}) Func0x1 {
	return ToFunc0x1(MakeFunc(fn))
}

type Func0x2 interface {
	Func
	Call0x2() (interface{}, interface{})
}

type shimFunc0x2 struct {
	inner Func
}

func (c *shimFunc0x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc0x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc0x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc0x2) Call0x2() (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc0x2(c Func) Func0x2 {
	if c.Type().NumIn() != 0 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 0 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func0x2); ok {
		return sc
	}
	return &shimFunc0x2{inner: c}
}

func MakeFunc0x2(fn interface{}) Func0x2 {
	return ToFunc0x2(MakeFunc(fn))
}

type Func0x3 interface {
	Func
	Call0x3() (interface{}, interface{}, interface{})
}

type shimFunc0x3 struct {
	inner Func
}

func (c *shimFunc0x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc0x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc0x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc0x3) Call0x3() (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc0x3(c Func) Func0x3 {
	if c.Type().NumIn() != 0 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 0 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func0x3); ok {
		return sc
	}
	return &shimFunc0x3{inner: c}
}

func MakeFunc0x3(fn interface{}) Func0x3 {
	return ToFunc0x3(MakeFunc(fn))
}

type Func1x0 interface {
	Func
	Call1x0(interface{})
}

type shimFunc1x0 struct {
	inner Func
}

func (c *shimFunc1x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc1x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc1x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc1x0) Call1x0(arg0 interface{}) {
	ret := c.inner.Call([]interface{}{arg0})
	_ = ret
	return
}

func ToFunc1x0(c Func) Func1x0 {
	if c.Type().NumIn() != 1 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 1 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func1x0); ok {
		return sc
	}
	return &shimFunc1x0{inner: c}
}

func MakeFunc1x0(fn interface{}) Func1x0 {
	return ToFunc1x0(MakeFunc(fn))
}

type Func1x1 interface {
	Func
	Call1x1(interface{}) interface{}
}

type shimFunc1x1 struct {
	inner Func
}

func (c *shimFunc1x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc1x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc1x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc1x1) Call1x1(arg0 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0})
	_ = ret
	return ret[0]
}

func ToFunc1x1(c Func) Func1x1 {
	if c.Type().NumIn() != 1 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 1 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func1x1); ok {
		return sc
	}
	return &shimFunc1x1{inner: c}
}

func MakeFunc1x1(fn interface{}) Func1x1 {
	return ToFunc1x1(MakeFunc(fn))
}

type Func1x2 interface {
	Func
	Call1x2(interface{}) (interface{}, interface{})
}

type shimFunc1x2 struct {
	inner Func
}

func (c *shimFunc1x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc1x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc1x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc1x2) Call1x2(arg0 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc1x2(c Func) Func1x2 {
	if c.Type().NumIn() != 1 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 1 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func1x2); ok {
		return sc
	}
	return &shimFunc1x2{inner: c}
}

func MakeFunc1x2(fn interface{}) Func1x2 {
	return ToFunc1x2(MakeFunc(fn))
}

type Func1x3 interface {
	Func
	Call1x3(interface{}) (interface{}, interface{}, interface{})
}

type shimFunc1x3 struct {
	inner Func
}

func (c *shimFunc1x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc1x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc1x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc1x3) Call1x3(arg0 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc1x3(c Func) Func1x3 {
	if c.Type().NumIn() != 1 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 1 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func1x3); ok {
		return sc
	}
	return &shimFunc1x3{inner: c}
}

func MakeFunc1x3(fn interface{}) Func1x3 {
	return ToFunc1x3(MakeFunc(fn))
}

type Func2x0 interface {
	Func
	Call2x0(interface{}, interface{})
}

type shimFunc2x0 struct {
	inner Func
}

func (c *shimFunc2x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc2x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc2x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc2x0) Call2x0(arg0, arg1 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1})
	_ = ret
	return
}

func ToFunc2x0(c Func) Func2x0 {
	if c.Type().NumIn() != 2 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 2 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func2x0); ok {
		return sc
	}
	return &shimFunc2x0{inner: c}
}

func MakeFunc2x0(fn interface{}) Func2x0 {
	return ToFunc2x0(MakeFunc(fn))
}

type Func2x1 interface {
	Func
	Call2x1(interface{}, interface{}) interface{}
}

type shimFunc2x1 struct {
	inner Func
}

func (c *shimFunc2x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc2x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc2x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc2x1) Call2x1(arg0, arg1 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1})
	_ = ret
	return ret[0]
}

func ToFunc2x1(c Func) Func2x1 {
	if c.Type().NumIn() != 2 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 2 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func2x1); ok {
		return sc
	}
	return &shimFunc2x1{inner: c}
}

func MakeFunc2x1(fn interface{}) Func2x1 {
	return ToFunc2x1(MakeFunc(fn))
}

type Func2x2 interface {
	Func
	Call2x2(interface{}, interface{}) (interface{}, interface{})
}

type shimFunc2x2 struct {
	inner Func
}

func (c *shimFunc2x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc2x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc2x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc2x2) Call2x2(arg0, arg1 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc2x2(c Func) Func2x2 {
	if c.Type().NumIn() != 2 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 2 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func2x2); ok {
		return sc
	}
	return &shimFunc2x2{inner: c}
}

func MakeFunc2x2(fn interface{}) Func2x2 {
	return ToFunc2x2(MakeFunc(fn))
}

type Func2x3 interface {
	Func
	Call2x3(interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc2x3 struct {
	inner Func
}

func (c *shimFunc2x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc2x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc2x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc2x3) Call2x3(arg0, arg1 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc2x3(c Func) Func2x3 {
	if c.Type().NumIn() != 2 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 2 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func2x3); ok {
		return sc
	}
	return &shimFunc2x3{inner: c}
}

func MakeFunc2x3(fn interface{}) Func2x3 {
	return ToFunc2x3(MakeFunc(fn))
}

type Func3x0 interface {
	Func
	Call3x0(interface{}, interface{}, interface{})
}

type shimFunc3x0 struct {
	inner Func
}

func (c *shimFunc3x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc3x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc3x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc3x0) Call3x0(arg0, arg1, arg2 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2})
	_ = ret
	return
}

func ToFunc3x0(c Func) Func3x0 {
	if c.Type().NumIn() != 3 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 3 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func3x0); ok {
		return sc
	}
	return &shimFunc3x0{inner: c}
}

func MakeFunc3x0(fn interface{}) Func3x0 {
	return ToFunc3x0(MakeFunc(fn))
}

type Func3x1 interface {
	Func
	Call3x1(interface{}, interface{}, interface{}) interface{}
}

type shimFunc3x1 struct {
	inner Func
}

func (c *shimFunc3x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc3x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc3x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc3x1) Call3x1(arg0, arg1, arg2 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2})
	_ = ret
	return ret[0]
}

func ToFunc3x1(c Func) Func3x1 {
	if c.Type().NumIn() != 3 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 3 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func3x1); ok {
		return sc
	}
	return &shimFunc3x1{inner: c}
}

func MakeFunc3x1(fn interface{}) Func3x1 {
	return ToFunc3x1(MakeFunc(fn))
}

type Func3x2 interface {
	Func
	Call3x2(interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc3x2 struct {
	inner Func
}

func (c *shimFunc3x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc3x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc3x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc3x2) Call3x2(arg0, arg1, arg2 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc3x2(c Func) Func3x2 {
	if c.Type().NumIn() != 3 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 3 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func3x2); ok {
		return sc
	}
	return &shimFunc3x2{inner: c}
}

func MakeFunc3x2(fn interface{}) Func3x2 {
	return ToFunc3x2(MakeFunc(fn))
}

type Func3x3 interface {
	Func
	Call3x3(interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc3x3 struct {
	inner Func
}

func (c *shimFunc3x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc3x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc3x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc3x3) Call3x3(arg0, arg1, arg2 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc3x3(c Func) Func3x3 {
	if c.Type().NumIn() != 3 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 3 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func3x3); ok {
		return sc
	}
	return &shimFunc3x3{inner: c}
}

func MakeFunc3x3(fn interface{}) Func3x3 {
	return ToFunc3x3(MakeFunc(fn))
}

type Func4x0 interface {
	Func
	Call4x0(interface{}, interface{}, interface{}, interface{})
}

type shimFunc4x0 struct {
	inner Func
}

func (c *shimFunc4x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc4x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc4x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc4x0) Call4x0(arg0, arg1, arg2, arg3 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3})
	_ = ret
	return
}

func ToFunc4x0(c Func) Func4x0 {
	if c.Type().NumIn() != 4 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 4 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func4x0); ok {
		return sc
	}
	return &shimFunc4x0{inner: c}
}

func MakeFunc4x0(fn interface{}) Func4x0 {
	return ToFunc4x0(MakeFunc(fn))
}

type Func4x1 interface {
	Func
	Call4x1(interface{}, interface{}, interface{}, interface{}) interface{}
}

type shimFunc4x1 struct {
	inner Func
}

func (c *shimFunc4x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc4x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc4x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc4x1) Call4x1(arg0, arg1, arg2, arg3 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3})
	_ = ret
	return ret[0]
}

func ToFunc4x1(c Func) Func4x1 {
	if c.Type().NumIn() != 4 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 4 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func4x1); ok {
		return sc
	}
	return &shimFunc4x1{inner: c}
}

func MakeFunc4x1(fn interface{}) Func4x1 {
	return ToFunc4x1(MakeFunc(fn))
}

type Func4x2 interface {
	Func
	Call4x2(interface{}, interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc4x2 struct {
	inner Func
}

func (c *shimFunc4x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc4x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc4x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc4x2) Call4x2(arg0, arg1, arg2, arg3 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc4x2(c Func) Func4x2 {
	if c.Type().NumIn() != 4 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 4 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func4x2); ok {
		return sc
	}
	return &shimFunc4x2{inner: c}
}

func MakeFunc4x2(fn interface{}) Func4x2 {
	return ToFunc4x2(MakeFunc(fn))
}

type Func4x3 interface {
	Func
	Call4x3(interface{}, interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc4x3 struct {
	inner Func
}

func (c *shimFunc4x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc4x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc4x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc4x3) Call4x3(arg0, arg1, arg2, arg3 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc4x3(c Func) Func4x3 {
	if c.Type().NumIn() != 4 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 4 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func4x3); ok {
		return sc
	}
	return &shimFunc4x3{inner: c}
}

func MakeFunc4x3(fn interface{}) Func4x3 {
	return ToFunc4x3(MakeFunc(fn))
}

type Func5x0 interface {
	Func
	Call5x0(interface{}, interface{}, interface{}, interface{}, interface{})
}

type shimFunc5x0 struct {
	inner Func
}

func (c *shimFunc5x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc5x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc5x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc5x0) Call5x0(arg0, arg1, arg2, arg3, arg4 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4})
	_ = ret
	return
}

func ToFunc5x0(c Func) Func5x0 {
	if c.Type().NumIn() != 5 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 5 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func5x0); ok {
		return sc
	}
	return &shimFunc5x0{inner: c}
}

func MakeFunc5x0(fn interface{}) Func5x0 {
	return ToFunc5x0(MakeFunc(fn))
}

type Func5x1 interface {
	Func
	Call5x1(interface{}, interface{}, interface{}, interface{}, interface{}) interface{}
}

type shimFunc5x1 struct {
	inner Func
}

func (c *shimFunc5x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc5x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc5x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc5x1) Call5x1(arg0, arg1, arg2, arg3, arg4 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4})
	_ = ret
	return ret[0]
}

func ToFunc5x1(c Func) Func5x1 {
	if c.Type().NumIn() != 5 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 5 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func5x1); ok {
		return sc
	}
	return &shimFunc5x1{inner: c}
}

func MakeFunc5x1(fn interface{}) Func5x1 {
	return ToFunc5x1(MakeFunc(fn))
}

type Func5x2 interface {
	Func
	Call5x2(interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc5x2 struct {
	inner Func
}

func (c *shimFunc5x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc5x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc5x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc5x2) Call5x2(arg0, arg1, arg2, arg3, arg4 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc5x2(c Func) Func5x2 {
	if c.Type().NumIn() != 5 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 5 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func5x2); ok {
		return sc
	}
	return &shimFunc5x2{inner: c}
}

func MakeFunc5x2(fn interface{}) Func5x2 {
	return ToFunc5x2(MakeFunc(fn))
}

type Func5x3 interface {
	Func
	Call5x3(interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc5x3 struct {
	inner Func
}

func (c *shimFunc5x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc5x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc5x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc5x3) Call5x3(arg0, arg1, arg2, arg3, arg4 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc5x3(c Func) Func5x3 {
	if c.Type().NumIn() != 5 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 5 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func5x3); ok {
		return sc
	}
	return &shimFunc5x3{inner: c}
}

func MakeFunc5x3(fn interface{}) Func5x3 {
	return ToFunc5x3(MakeFunc(fn))
}

type Func6x0 interface {
	Func
	Call6x0(interface{}, interface{}, interface{}, interface{}, interface{}, interface{})
}

type shimFunc6x0 struct {
	inner Func
}

func (c *shimFunc6x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc6x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc6x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc6x0) Call6x0(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5})
	_ = ret
	return
}

func ToFunc6x0(c Func) Func6x0 {
	if c.Type().NumIn() != 6 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 6 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func6x0); ok {
		return sc
	}
	return &shimFunc6x0{inner: c}
}

func MakeFunc6x0(fn interface{}) Func6x0 {
	return ToFunc6x0(MakeFunc(fn))
}

type Func6x1 interface {
	Func
	Call6x1(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}
}

type shimFunc6x1 struct {
	inner Func
}

func (c *shimFunc6x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc6x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc6x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc6x1) Call6x1(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5})
	_ = ret
	return ret[0]
}

func ToFunc6x1(c Func) Func6x1 {
	if c.Type().NumIn() != 6 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 6 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func6x1); ok {
		return sc
	}
	return &shimFunc6x1{inner: c}
}

func MakeFunc6x1(fn interface{}) Func6x1 {
	return ToFunc6x1(MakeFunc(fn))
}

type Func6x2 interface {
	Func
	Call6x2(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc6x2 struct {
	inner Func
}

func (c *shimFunc6x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc6x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc6x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc6x2) Call6x2(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc6x2(c Func) Func6x2 {
	if c.Type().NumIn() != 6 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 6 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func6x2); ok {
		return sc
	}
	return &shimFunc6x2{inner: c}
}

func MakeFunc6x2(fn interface{}) Func6x2 {
	return ToFunc6x2(MakeFunc(fn))
}

type Func6x3 interface {
	Func
	Call6x3(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc6x3 struct {
	inner Func
}

func (c *shimFunc6x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc6x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc6x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc6x3) Call6x3(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc6x3(c Func) Func6x3 {
	if c.Type().NumIn() != 6 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 6 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func6x3); ok {
		return sc
	}
	return &shimFunc6x3{inner: c}
}

func MakeFunc6x3(fn interface{}) Func6x3 {
	return ToFunc6x3(MakeFunc(fn))
}

type Func7x0 interface {
	Func
	Call7x0(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{})
}

type shimFunc7x0 struct {
	inner Func
}

func (c *shimFunc7x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc7x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc7x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc7x0) Call7x0(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6})
	_ = ret
	return
}

func ToFunc7x0(c Func) Func7x0 {
	if c.Type().NumIn() != 7 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 7 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func7x0); ok {
		return sc
	}
	return &shimFunc7x0{inner: c}
}

func MakeFunc7x0(fn interface{}) Func7x0 {
	return ToFunc7x0(MakeFunc(fn))
}

type Func7x1 interface {
	Func
	Call7x1(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}
}

type shimFunc7x1 struct {
	inner Func
}

func (c *shimFunc7x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc7x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc7x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc7x1) Call7x1(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6})
	_ = ret
	return ret[0]
}

func ToFunc7x1(c Func) Func7x1 {
	if c.Type().NumIn() != 7 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 7 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func7x1); ok {
		return sc
	}
	return &shimFunc7x1{inner: c}
}

func MakeFunc7x1(fn interface{}) Func7x1 {
	return ToFunc7x1(MakeFunc(fn))
}

type Func7x2 interface {
	Func
	Call7x2(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc7x2 struct {
	inner Func
}

func (c *shimFunc7x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc7x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc7x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc7x2) Call7x2(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc7x2(c Func) Func7x2 {
	if c.Type().NumIn() != 7 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 7 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func7x2); ok {
		return sc
	}
	return &shimFunc7x2{inner: c}
}

func MakeFunc7x2(fn interface{}) Func7x2 {
	return ToFunc7x2(MakeFunc(fn))
}

type Func7x3 interface {
	Func
	Call7x3(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc7x3 struct {
	inner Func
}

func (c *shimFunc7x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc7x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc7x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc7x3) Call7x3(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc7x3(c Func) Func7x3 {
	if c.Type().NumIn() != 7 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 7 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func7x3); ok {
		return sc
	}
	return &shimFunc7x3{inner: c}
}

func MakeFunc7x3(fn interface{}) Func7x3 {
	return ToFunc7x3(MakeFunc(fn))
}

type Func8x0 interface {
	Func
	Call8x0(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{})
}

type shimFunc8x0 struct {
	inner Func
}

func (c *shimFunc8x0) Name() string {
	return c.inner.Name()
}

func (c *shimFunc8x0) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc8x0) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc8x0) Call8x0(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	_ = ret
	return
}

func ToFunc8x0(c Func) Func8x0 {
	if c.Type().NumIn() != 8 || c.Type().NumOut() != 0 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 8 inputs and 0 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func8x0); ok {
		return sc
	}
	return &shimFunc8x0{inner: c}
}

func MakeFunc8x0(fn interface{}) Func8x0 {
	return ToFunc8x0(MakeFunc(fn))
}

type Func8x1 interface {
	Func
	Call8x1(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}
}

type shimFunc8x1 struct {
	inner Func
}

func (c *shimFunc8x1) Name() string {
	return c.inner.Name()
}

func (c *shimFunc8x1) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc8x1) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc8x1) Call8x1(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) interface{} {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	_ = ret
	return ret[0]
}

func ToFunc8x1(c Func) Func8x1 {
	if c.Type().NumIn() != 8 || c.Type().NumOut() != 1 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 8 inputs and 1 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func8x1); ok {
		return sc
	}
	return &shimFunc8x1{inner: c}
}

func MakeFunc8x1(fn interface{}) Func8x1 {
	return ToFunc8x1(MakeFunc(fn))
}

type Func8x2 interface {
	Func
	Call8x2(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{})
}

type shimFunc8x2 struct {
	inner Func
}

func (c *shimFunc8x2) Name() string {
	return c.inner.Name()
}

func (c *shimFunc8x2) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc8x2) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc8x2) Call8x2(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) (interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	_ = ret
	return ret[0], ret[1]
}

func ToFunc8x2(c Func) Func8x2 {
	if c.Type().NumIn() != 8 || c.Type().NumOut() != 2 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 8 inputs and 2 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func8x2); ok {
		return sc
	}
	return &shimFunc8x2{inner: c}
}

func MakeFunc8x2(fn interface{}) Func8x2 {
	return ToFunc8x2(MakeFunc(fn))
}

type Func8x3 interface {
	Func
	Call8x3(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) (interface{}, interface{}, interface{})
}

type shimFunc8x3 struct {
	inner Func
}

func (c *shimFunc8x3) Name() string {
	return c.inner.Name()
}

func (c *shimFunc8x3) Type() reflect.Type {
	return c.inner.Type()
}

func (c *shimFunc8x3) Call(args []interface{}) []interface{} {
	return c.inner.Call(args)
}

func (c *shimFunc8x3) Call8x3(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) (interface{}, interface{}, interface{}) {
	ret := c.inner.Call([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	_ = ret
	return ret[0], ret[1], ret[2]
}

func ToFunc8x3(c Func) Func8x3 {
	if c.Type().NumIn() != 8 || c.Type().NumOut() != 3 {
		panic(fmt.Sprintf("Incompatible func type: got func %v with %v inputs and %v outputs, want 8 inputs and 3 outputs", c.Type(), c.Type().NumIn(), c.Type().NumOut()))
	}
	if sc, ok := c.(Func8x3); ok {
		return sc
	}
	return &shimFunc8x3{inner: c}
}

func MakeFunc8x3(fn interface{}) Func8x3 {
	return ToFunc8x3(MakeFunc(fn))
}
