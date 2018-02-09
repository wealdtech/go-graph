package nodes

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

type SimpleNode struct {
	id    int64
	attrs map[interface{}]interface{}
}

func NewSimpleNode(nid int64) *SimpleNode {
	return &SimpleNode{
		id:    nid,
		attrs: make(map[interface{}]interface{}),
	}
}

func (n *SimpleNode) Id() int64 {
	return n.id
}

func (n *SimpleNode) Attributes() *map[interface{}]interface{} {
	return &n.attrs
}

func (n *SimpleNode) Attribute(key interface{}) interface{} {
	return n.attrs[key]
}

func (n *SimpleNode) SetAttributes(attrs map[interface{}]interface{}) {
	n.attrs = attrs
}

func (n *SimpleNode) SetAttribute(key, value interface{}) {
	n.attrs[key] = value
}
