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

	"github.com/wealdtech/graph"
)

func Marshal(g graph.Graph) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("graph g {\n")

	graphDefaults(g, &buffer)
	nodeDefaults(g, &buffer)
	edgeDefaults(g, &buffer)

	nodeEntities(g, &buffer)
	buffer.WriteString("}")
	return buffer.Bytes()
}

func graphDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.GraphDefaults()) > 0 {
		buffer.WriteString("  graph [")
		for _, key := range sortAttrKeys(g.GraphDefaults()) {
			buffer.WriteString(fmt.Sprintf(" %v=%v", key, g.GraphDefault(key)))
		}
		buffer.WriteString(" ];\n")
	}
}

func nodeDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.NodeDefaults()) > 0 {
		buffer.WriteString("  node [")
		for _, key := range sortAttrKeys(g.NodeDefaults()) {
			buffer.WriteString(fmt.Sprintf(" %v=%v", key, g.NodeDefault(key)))
		}
		buffer.WriteString(" ];\n")
	}
}

func edgeDefaults(g graph.Graph, buffer *bytes.Buffer) {
	if len(*g.EdgeDefaults()) > 0 {
		buffer.WriteString("  edge [")
		for _, key := range sortAttrKeys(g.EdgeDefaults()) {
			buffer.WriteString(fmt.Sprintf(" %v=%v", key, g.EdgeDefault(key)))
		}
		buffer.WriteString(" ];\n")
	}
}

func nodeEntities(g graph.Graph, buffer *bytes.Buffer) {
	var sortedNodeKeys []int64
	for _, node := range g.Nodes() {
		sortedNodeKeys = append(sortedNodeKeys, node.Id())
	}
	sort.Slice(sortedNodeKeys, func(i, j int) bool { return sortedNodeKeys[i] < sortedNodeKeys[j] })
	for i := range sortedNodeKeys {
		node := g.Node(sortedNodeKeys[i])
		buffer.WriteString(fmt.Sprintf("  %d", node.Id()))
		if len(*node.Attributes()) > 0 {
			buffer.WriteString(" [")
			for _, key := range sortAttrKeys(node.Attributes()) {
				buffer.WriteString(fmt.Sprintf(" %v=%v", key, node.Attribute(key)))
			}
			buffer.WriteString(" ]")
		}
		buffer.WriteString(";\n")
		nodeEdges(g, node.Id(), buffer)
	}
}

func nodeEdges(g graph.Graph, nid int64, buffer *bytes.Buffer) {
	var sortedEdgeKeys []int64
	for _, edge := range g.Edges(nid) {
		if edge.From() == nid {
			sortedEdgeKeys = append(sortedEdgeKeys, edge.To())
		}
	}
	sort.Slice(sortedEdgeKeys, func(i, j int) bool { return sortedEdgeKeys[i] < sortedEdgeKeys[j] })
	for j := range sortedEdgeKeys {
		edge := g.Edge(nid, sortedEdgeKeys[j])
		buffer.WriteString(fmt.Sprintf("  %d -- %d", edge.From(), edge.To()))
		if len(*edge.Attributes()) > 0 {
			buffer.WriteString(" [")
			for _, key := range sortAttrKeys(edge.Attributes()) {
				buffer.WriteString(fmt.Sprintf(" %v=%v", key, edge.Attribute(key)))
			}
			buffer.WriteString(" ]")
		}
		buffer.WriteString(";\n")
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
