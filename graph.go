package graph

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

type Node interface {
	Attributed

	Id() int64
}

type Edge interface {
	Attributed

	From() int64
	To() int64
}

type Graph interface {
	GraphDefaults() *map[interface{}]interface{}

	GraphDefault(interface{}) interface{}

	SetGraphDefaults(map[interface{}]interface{})

	NodeDefaults() *map[interface{}]interface{}

	NodeDefault(interface{}) interface{}

	SetNodeDefaults(map[interface{}]interface{})

	EdgeDefaults() *map[interface{}]interface{}

	EdgeDefault(interface{}) interface{}

	SetEdgeDefaults(map[interface{}]interface{})

	HasNode(nid int64) bool

	HasEdge(aid, bid int64) bool

	Node(nid int64) Node

	Nodes() []Node

	// ConnectedNodes reurns the nodes connected to a given node within
	// a given distance
	ConnectedNodes(aid, distance int64) []Node

	Edge(aid, bid int64) Edge

	Edges(nid int64) []Edge
}

type NodeManager interface {
	AddNode(node Node) error
	RemoveNode(nid int64) Node
}

type EdgeManager interface {
	AddEdge(edge Edge) error
	RemoveEdge(aid, bid int64) []Edge
}
