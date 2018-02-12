package edges

// Copyright © 2018 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

type DirectedEdge struct {
	from  int64
	to    int64
	attrs map[interface{}]interface{}
}

func NewDirectedEdge(from, to int64) *DirectedEdge {
	return &DirectedEdge{
		from:  from,
		to:    to,
		attrs: make(map[interface{}]interface{}),
	}
}

func (e *DirectedEdge) From() int64 {
	return e.from
}

func (e *DirectedEdge) To() int64 {
	return e.to
}

func (e *DirectedEdge) Attributes() *map[interface{}]interface{} {
	return &e.attrs
}

func (e *DirectedEdge) Attribute(key interface{}) interface{} {
	return e.attrs[key]
}

func (e *DirectedEdge) SetAttributes(attrs map[interface{}]interface{}) {
	e.attrs = attrs
}

func (e *DirectedEdge) SetAttribute(key, value interface{}) {
	e.attrs[key] = value
}
