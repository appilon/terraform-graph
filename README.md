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

## Thoughts

This is currently just a POC so here are some thoughts on where to take this tool. The goal of this tool is to provide some utility around visualization for terraform after it has reached a large scale (hundreds of resources). Terraform's resource data structure can be thought of as a directed graph, where interpolated values from one resource used in another creates an implicit dependency. Terraform currently supports exporting that graph with graphviz dot notation. The trouble with graphviz is it exports either a rastered image (png, jpg) or vector graphic (svg). Neither of those formats are interactive or can provide a dynamic user experience.

The goal of this tool was to provide some form of drop in replacement for graphviz. The technology of choice would be web tech, the frontrunner of visualization being d3.js. Before going any further in development I scoured the web for any similar approaches. There is one somewhat mature tool made in Python that does many great things:

https://github.com/28mm/blast-radius

Linked in that project's README is a really nice blog detailing development of that tool.

Some decisions for implementing this ourselves is 1) we always prefer work in Go, 2) although I'm not opposed to the above approach of hosting a web server to serve the web app, I would much prefer pipe the output of `terraform graph` into our new tool (much like how users typically pipe to `dot`). I also found a package called `browser` that allowed us a cross-platform method of opening a computer's default browser to our "single page app" (if you can call it that).

The first iteration involved ingesting the dot notation, parsing the graph with a Go based graphviz library, then coercing the data to a layout D3 could render. It was a cool first pass but I quickly realized illustrating graphs clearly is challenging, and forcing the data to some arbitrary new json structure felt brittle.

I then stumbled upon a popular frontend project called `dagre` which focuses on creating graphviz layouts for javascript. One of their supported renderers is d3:

https://github.com/dagrejs/dagre-d3

I think if we moved forward with the tool this should be the foundation of the frontend tech, it handles (almost with no issue) parsing dot notation and rendering out the graph with d3. We should have the ability to then create any kind of interaction story we want. That is, in JS I don't see there being much Go code at play here other than being a good choice for making cross platform binaries.

Some current aspects of the POC that I don't like:
 
Had to search/replace some shape annotation (see code) `terraform graph` outputted but `dagre` did not like. The good news is we control `terraform` so we can upgrade/improve what `terraform graph` outputs (currently the labels seem a bit strange with this `[root] ` prefix on nodes, I'm not sure either if the graph is actually accurate? I did not really try and make sense of that, just get the graph rendering).

Having extensive frontend experience I can say this is not how you would make a modern single page app. Currently I've inlined the html template into go (so everything is baked into the binary), I link the dependencies through old fashioned `<script src="..."></script>` (meaning currently this tool requires an internet connection at runtime). That can be solved by inlining the JS deps into the template but now we need to perform vendoring of the JS dependencies. As is the development story involves be working in an html file, when done I paste the whole thing into `template.go`, we will need some javascript packaging to make development tolerable. Although I love Go, I would argue given how little Go code there is here, we could do the entire tool in JS, but still make it a binary with one of these:

 * https://github.com/zeit/pkg
 * https://github.com/nexe/nexe
 * https://github.com/pmq20/node-packer

The versioning and inlining of javascript into HTML can be achieved with the plethora of bundling tech:

 * https://github.com/webpack/webpack
 * https://github.com/parcel-bundler/parcel
 * https://github.com/browserify/browserify (old but still good)

It is very possible there are Go tools that can address the issues mentioned as well! Regardless though the user experience will be driven with javascript and lots of javascript/frontend development will be required.
