package graphs

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

import (
	"fmt"

	"github.com/wealdtech/graph"
)

type UndirectedGraph struct {
	nodes         map[int64]graph.Node
	edges         map[int64]map[int64]graph.Edge
	graphDefaults map[interface{}]interface{}
	nodeDefaults  map[interface{}]interface{}
	edgeDefaults  map[interface{}]interface{}
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		nodes:         make(map[int64]graph.Node),
		edges:         make(map[int64]map[int64]graph.Edge),
		graphDefaults: make(map[interface{}]interface{}),
		nodeDefaults:  make(map[interface{}]interface{}),
		edgeDefaults:  make(map[interface{}]interface{}),
	}
}

func (g *UndirectedGraph) HasNode(nid int64) bool {
	_, ok := g.nodes[nid]
	return ok
}

func (g *UndirectedGraph) HasEdge(aid, bid int64) bool {
	_, ok := g.edges[aid][bid]
	return ok
}

func (g *UndirectedGraph) Node(nid int64) graph.Node {
	return g.nodes[nid]
}

func (g *UndirectedGraph) Nodes() []graph.Node {
	nodes := make([]graph.Node, 0, len(g.nodes))
	for _, node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *UndirectedGraph) Edges(nid int64) []graph.Edge {
	node := g.Node(nid)
	edges := make([]graph.Edge, 0, len(g.nodes))
	if node != nil {
		for _, edge := range g.edges[nid] {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *UndirectedGraph) ConnectedNodes(nid int64, distance int64) []graph.Node {
	// TODO handle distances other than 1
	if distance != 1 {
		panic("Distance not supported")
	}
	nodes := make([]graph.Node, 0, len(g.edges[nid]))
	for nid, _ := range g.edges[nid] {
		nodes = append(nodes, g.nodes[nid])
	}
	return nodes
}

func (g *UndirectedGraph) Edge(aid, bid int64) graph.Edge {
	return g.edges[aid][bid]
}

func (g *UndirectedGraph) GraphDefaults() *map[interface{}]interface{} {
	return &g.graphDefaults
}

func (g *UndirectedGraph) GraphDefault(key interface{}) interface{} {
	return g.graphDefaults[key]
}

func (g *UndirectedGraph) SetGraphDefaults(defaults map[interface{}]interface{}) {
	g.graphDefaults = defaults
}

func (g *UndirectedGraph) NodeDefaults() *map[interface{}]interface{} {
	return &g.nodeDefaults
}

func (g *UndirectedGraph) NodeDefault(key interface{}) interface{} {
	return g.nodeDefaults[key]
}

func (g *UndirectedGraph) SetNodeDefaults(defaults map[interface{}]interface{}) {
	g.nodeDefaults = defaults
}

func (g *UndirectedGraph) EdgeDefaults() *map[interface{}]interface{} {
	return &g.edgeDefaults
}

func (g *UndirectedGraph) EdgeDefault(key interface{}) interface{} {
	return g.edgeDefaults[key]
}

func (g *UndirectedGraph) SetEdgeDefaults(defaults map[interface{}]interface{}) {
	g.edgeDefaults = defaults
}

// NodeManager
func (g *UndirectedGraph) AddNode(node graph.Node) error {
	if g.HasNode(node.Id()) {
		return fmt.Errorf("Node with ID %v already exists", node.Id())
	}
	g.nodes[node.Id()] = node
	g.edges[node.Id()] = make(map[int64]graph.Edge)
	return nil
}

func (g *UndirectedGraph) RemoveNode(nid int64) graph.Node {
	node := g.Node(nid)
	delete(g.nodes, nid)
	// Delete associated edges
	for bid, _ := range g.edges[nid] {
		delete(g.edges[bid], nid)
	}
	delete(g.edges, nid)
	return node
}

// EdgeManager
// AddEdge adds an edge to a graph
func (g *UndirectedGraph) AddEdge(edge graph.Edge) error {
	if !g.HasNode(edge.From()) {
		return fmt.Errorf("Unknown edge start %v", edge.From())
	}
	if !g.HasNode(edge.To()) {
		return fmt.Errorf("Unknown edge end %v", edge.To())
	}
	if g.HasEdge(edge.From(), edge.To()) {
		return fmt.Errorf("Edge from %v to %v already exists", edge.From(), edge.To())
	}
	g.edges[edge.From()][edge.To()] = edge
	if edge.From() != edge.To() {
		g.edges[edge.To()][edge.From()] = edge
	}
	return nil
}

func (g *UndirectedGraph) RemoveEdge(aid, bid int64) graph.Edge {
	edge := g.Edge(aid, bid)
	delete(g.edges[aid], bid)
	if aid != bid {
		delete(g.edges[bid], aid)
	}
	return edge
}
