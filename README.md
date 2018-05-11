# Terraform Graph

Hack job drop in replacement for graph viz, powered by D3

## Installation
```
go get -u github.com/appilon/terraform-graph
```
## Usage
From a terraform directory pipe terraform graph to terraform-graph (now realizing the binary name wasn't ideal)
```
terraform graph | terraform-graph
or
terraform graph > graph.out
terraform-graph graph.out
```
Your browser should open with a D3 visualization.

## Important Note
Part of this hack assumes terraform outputs a graph with root `"[root] root"` (did not want to have to walk the graph and determine which node is the root). This may not be true in all cases but for a working example check out this repo and run terraform from `tf-example`.