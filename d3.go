package main

import (
	"encoding/json"
	"errors"

	"github.com/awalterschulze/gographviz"
)

type D3Tree struct {
	Name     string    `json:"name"`
	Children []*D3Tree `json:"children,omitempty"`
}

func expand(node string, graph *gographviz.Graph) *D3Tree {
	tree := &D3Tree{
		Name: node,
	}

	children, exists := graph.Edges.SrcToDsts[node]
	if exists {
		for child, _ := range children {
			tree.Children = append(tree.Children, expand(child, graph))
		}
	}

	return tree
}

func toD3Json(graph *gographviz.Graph) ([]byte, error) {
	// this is cheating, assume TF graph has this weird root
	root, exists := graph.Nodes.Lookup[`"[root] root"`]
	if !exists {
		return nil, errors.New("No root node")
	}

	tree := expand(root.Name, graph)

	return json.Marshal(tree)
}
