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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wealdtech/graph/edges"
	"github.com/wealdtech/graph/graphs"
	"github.com/wealdtech/graph/nodes"
)

func TestSimple(t *testing.T) {
	g := graphs.NewUndirectedGraph()
	node1 := nodes.NewSimpleNode(1)
	err := g.AddNode(node1)
	assert.NoError(t, err)

	output := Marshal(g)
	assert.Equal(t, `graph g {
  1;
}`, string(output))
}

func TestSimple2(t *testing.T) {
	g := graphs.NewUndirectedGraph()
	node1 := nodes.NewSimpleNode(1)
	node1.SetAttributes(map[interface{}]interface{}{"color": "red", "shape": "circle"})
	err := g.AddNode(node1)
	assert.NoError(t, err)
	node2 := nodes.NewSimpleNode(2)
	err = g.AddNode(node2)
	assert.NoError(t, err)
	edge12 := edges.NewUndirectedEdge(1, 2)
	err = g.AddEdge(edge12)
	assert.NoError(t, err)

	output := Marshal(g)
	assert.Equal(t, `graph g {
  1 [ color=red shape=circle ];
  1 -- 2;
  2;
}`, string(output))
}

func TestGraphLevelAttrs(t *testing.T) {
	g := graphs.NewUndirectedGraph()
	g.SetNodeDefaults(map[interface{}]interface{}{"color": "blue", "shape": "diamond"})
	node1 := nodes.NewSimpleNode(1)
	node1.SetAttributes(map[interface{}]interface{}{"color": "red", "shape": "circle"})
	err := g.AddNode(node1)
	assert.NoError(t, err)
	node2 := nodes.NewSimpleNode(2)
	err = g.AddNode(node2)
	assert.NoError(t, err)
	output := Marshal(g)
	assert.Equal(t, `graph g {
  node [ color=blue shape=diamond ];
  1 [ color=red shape=circle ];
  2;
}`, string(output))
}

func TestNodeOrdering(t *testing.T) {
	g := graphs.NewUndirectedGraph()
	for i := 1; i <= 10; i++ {
		g.AddNode(nodes.NewSimpleNode(int64(i)))
	}
	output := Marshal(g)
	assert.Equal(t, `graph g {
  1;
  2;
  3;
  4;
  5;
  6;
  7;
  8;
  9;
  10;
}`, string(output))
}
