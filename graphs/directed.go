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

	"github.com/wealdtech/go-graph"
)

type DirectedGraph struct {
	nodes         map[int64]graph.Node
	edges         map[int64]map[int64]graph.Edge
	graphDefaults map[interface{}]interface{}
	nodeDefaults  map[interface{}]interface{}
	edgeDefaults  map[interface{}]interface{}
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		nodes:         make(map[int64]graph.Node),
		edges:         make(map[int64]map[int64]graph.Edge),
		graphDefaults: make(map[interface{}]interface{}),
		nodeDefaults:  make(map[interface{}]interface{}),
		edgeDefaults:  make(map[interface{}]interface{}),
	}
}

func (g *DirectedGraph) HasNode(nid int64) bool {
	_, ok := g.nodes[nid]
	return ok
}

func (g *DirectedGraph) HasEdge(aid, bid int64) bool {
	_, ok := g.edges[aid][bid]
	return ok
}

func (g *DirectedGraph) Node(nid int64) graph.Node {
	return g.nodes[nid]
}

func (g *DirectedGraph) Nodes() []graph.Node {
	nodes := make([]graph.Node, 0, len(g.nodes))
	for _, node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

// Edges returns all edges originating at this node
func (g *DirectedGraph) Edges(nid int64) []graph.Edge {
	node := g.Node(nid)
	edges := make([]graph.Edge, 0, len(g.nodes))
	if node != nil {
		for _, edge := range g.edges[nid] {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *DirectedGraph) ConnectedNodes(nid int64, distance int64) []graph.Node {
	nidMap := make(map[int64]bool)
	g.connectedNodes(nid, distance, &nidMap)
	nodes := make([]graph.Node, 0, len(nidMap))
	for nid := range nidMap {
		nodes = append(nodes, g.nodes[nid])
	}
	return nodes
}

func (g *DirectedGraph) connectedNodes(nid int64, distance int64, nidMap *map[int64]bool) {
	(*nidMap)[nid] = true
	if distance > 0 {
		for neighbourid, _ := range g.edges[nid] {
			(*nidMap)[neighbourid] = true
			g.connectedNodes(neighbourid, distance-1, nidMap)
		}
	}
}

func (g *DirectedGraph) Edge(aid, bid int64) graph.Edge {
	return g.edges[aid][bid]
}

func (g *DirectedGraph) GraphDefaults() *map[interface{}]interface{} {
	return &g.graphDefaults
}

func (g *DirectedGraph) GraphDefault(key interface{}) interface{} {
	return g.graphDefaults[key]
}

func (g *DirectedGraph) SetGraphDefaults(defaults map[interface{}]interface{}) {
	g.graphDefaults = defaults
}

func (g *DirectedGraph) NodeDefaults() *map[interface{}]interface{} {
	return &g.nodeDefaults
}

func (g *DirectedGraph) NodeDefault(key interface{}) interface{} {
	return g.nodeDefaults[key]
}

func (g *DirectedGraph) SetNodeDefaults(defaults map[interface{}]interface{}) {
	g.nodeDefaults = defaults
}

func (g *DirectedGraph) EdgeDefaults() *map[interface{}]interface{} {
	return &g.edgeDefaults
}

func (g *DirectedGraph) EdgeDefault(key interface{}) interface{} {
	return g.edgeDefaults[key]
}

func (g *DirectedGraph) SetEdgeDefaults(defaults map[interface{}]interface{}) {
	g.edgeDefaults = defaults
}

// NodeManager
func (g *DirectedGraph) AddNode(node graph.Node) error {
	if g.HasNode(node.Id()) {
		return fmt.Errorf("Node with ID %v already exists", node.Id())
	}
	g.nodes[node.Id()] = node
	g.edges[node.Id()] = make(map[int64]graph.Edge)
	return nil
}

func (g *DirectedGraph) RemoveNode(nid int64) graph.Node {
	node := g.Node(nid)
	delete(g.nodes, nid)
	// Delete edges that start at this node
	for bid, _ := range g.edges[nid] {
		delete(g.edges[bid], nid)
	}
	// Delete edges that terminate at this node
	for aid, _ := range g.edges {
		for bid, _ := range g.edges[aid] {
			if bid == nid {
				delete(g.edges[aid], bid)
			}
		}
	}
	delete(g.edges, nid)
	return node
}

// EdgeManager
// AddEdge adds an edge to a graph
func (g *DirectedGraph) AddEdge(edge graph.Edge) error {
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
	return nil
}

func (g *DirectedGraph) RemoveEdge(aid, bid int64) graph.Edge {
	edge := g.Edge(aid, bid)
	delete(g.edges[aid], bid)
	return edge
}
