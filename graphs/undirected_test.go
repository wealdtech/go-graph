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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wealdtech/graph/edges"
	"github.com/wealdtech/graph/nodes"
)

func TestAddRemoveNodes(t *testing.T) {
	g := NewUndirectedGraph()

	// Add a node
	node1 := nodes.NewSimpleNode(1)
	err := g.AddNode(node1)
	assert.NoError(t, err)
	storedNode1 := g.Node(1)
	assert.NotNil(t, storedNode1)
	assert.Equal(t, node1, storedNode1)
	allNodes := g.Nodes()
	assert.Len(t, allNodes, 1)

	// Add another node
	node2 := nodes.NewSimpleNode(2)
	err = g.AddNode(node2)
	assert.NoError(t, err)
	storedNode2 := g.Node(2)
	assert.NotNil(t, storedNode2)
	assert.Equal(t, node2, storedNode2)
	allNodes = g.Nodes()
	assert.Len(t, allNodes, 2)

	// Remove a node
	removedNode := g.RemoveNode(1)
	assert.NotNil(t, removedNode)
	assert.Equal(t, node1, removedNode)
	allNodes = g.Nodes()
	assert.Len(t, allNodes, 1)
}

func TestAddRemoveEdges(t *testing.T) {
	g := NewUndirectedGraph()

	// Add three nodes
	node1 := nodes.NewSimpleNode(1)
	err := g.AddNode(node1)
	assert.NoError(t, err)
	node2 := nodes.NewSimpleNode(2)
	err = g.AddNode(node2)
	assert.NoError(t, err)
	node3 := nodes.NewSimpleNode(3)
	err = g.AddNode(node3)
	assert.NoError(t, err)

	// Add an edge between two nodes
	edge12 := edges.NewUndirectedEdge(1, 2)
	err = g.AddEdge(edge12)
	assert.NoError(t, err)
	storedEdge12 := g.Edge(1, 2)
	assert.NotNil(t, storedEdge12)
	assert.Equal(t, edge12, storedEdge12)
	node1ConnectedNodes := g.ConnectedNodes(1, 1)
	assert.Len(t, node1ConnectedNodes, 1)
	// Should be able to pull the reverse
	storedEdge21 := g.Edge(2, 1)
	assert.NotNil(t, storedEdge21)
	assert.Equal(t, edge12, storedEdge21)

	// Try to add the same edge again; should fail
	badEdge := edges.NewUndirectedEdge(1, 2)
	err = g.AddEdge(badEdge)
	assert.Error(t, err)

	// Add another edge between two nodes
	edge13 := edges.NewUndirectedEdge(1, 3)
	err = g.AddEdge(edge13)
	assert.NoError(t, err)

	node1ConnectedNodes = g.ConnectedNodes(1, 1)
	assert.Len(t, node1ConnectedNodes, 2)
}
