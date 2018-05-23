# Terraform Graph

Hack job drop in replacement for graph viz, powered by Dagre-D3

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
Your browser should open with a visualization.
