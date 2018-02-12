package dot

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
	"bytes"
	"fmt"
	"sort"

	"github.com/wealdtech/go-graph"
	"github.com/wealdtech/go-graph/graphs"
)

func Marshal(g graph.Graph) []byte {
	var buffer bytes.Buffer
	var directed bool
	switch g.(type) {
	case *graphs.UndirectedGraph:
		directed = false
	case *graphs.DirectedGraph:
		directed = true
	}

	if directed {
		buffer.WriteString("digraph g {\n")
	} else {
		buffer.WriteString("graph g {\n")
	}

	graphDefaults(g, &buffer)
	nodeDefaults(g, &buffer)
	edgeDefaults(g, &buffer)
	nodeEntities(g, &buffer, directed) // nodeEntities writes out edges as well

	buffer.WriteString("}")
	return buffer.Bytes()
}

func graphDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.GraphDefaults()) > 0 {
		buffer.WriteString("  graph")
		writeAttrs(g, g.GraphDefaults(), buffer)
		buffer.WriteString(";\n")
	}
}

func nodeDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.NodeDefaults()) > 0 {
		buffer.WriteString("  node")
		writeAttrs(g, g.NodeDefaults(), buffer)
		buffer.WriteString(";\n")
	}
}

func edgeDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.EdgeDefaults()) > 0 {
		buffer.WriteString("  edge")
		writeAttrs(g, g.EdgeDefaults(), buffer)
		buffer.WriteString(";\n")
	}
}

func nodeEntities(g graph.Graph, buffer *bytes.Buffer, directed bool) {
	var sortedNodeKeys []int64
	for _, node := range g.Nodes() {
		sortedNodeKeys = append(sortedNodeKeys, node.Id())
	}
	sort.Slice(sortedNodeKeys, func(i, j int) bool { return sortedNodeKeys[i] < sortedNodeKeys[j] })
	for i := range sortedNodeKeys {
		node := g.Node(sortedNodeKeys[i])
		buffer.WriteString(fmt.Sprintf("  %d", node.Id()))
		writeAttrs(g, node.Attributes(), buffer)
		buffer.WriteString(";\n")
		nodeEdges(g, node.Id(), buffer, directed)
	}
}

func nodeEdges(g graph.Graph, nid int64, buffer *bytes.Buffer, directed bool) {
	var sortedEdgeKeys []int64
	for _, edge := range g.Edges(nid) {
		if edge.From() == nid {
			sortedEdgeKeys = append(sortedEdgeKeys, edge.To())
		}
	}
	sort.Slice(sortedEdgeKeys, func(i, j int) bool { return sortedEdgeKeys[i] < sortedEdgeKeys[j] })
	for j := range sortedEdgeKeys {
		edge := g.Edge(nid, sortedEdgeKeys[j])
		if directed {
			buffer.WriteString(fmt.Sprintf("  %d -> %d", edge.From(), edge.To()))
		} else {
			buffer.WriteString(fmt.Sprintf("  %d -- %d", edge.From(), edge.To()))
		}
		writeAttrs(g, edge.Attributes(), buffer)
		buffer.WriteString(";\n")
	}
}

func writeAttrs(g graph.Graph, attrs *map[interface{}]interface{}, buffer *bytes.Buffer) {
	if attrs != nil && len(*attrs) > 0 {
		buffer.WriteString(" [")
		for _, key := range sortAttrKeys(attrs) {
			buffer.WriteString(fmt.Sprintf(" %v=\"%v\"", key, (*attrs)[key]))
		}
		buffer.WriteString(" ]")
	}
}

func sortAttrKeys(attrs *map[interface{}]interface{}) []interface{} {
	var sortedKeys []interface{}
	for k, _ := range *attrs {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool { return sortedKeys[i].(string) < sortedKeys[j].(string) })
	return sortedKeys
}
