package main

const template = `
<!doctype html>

<meta charset="utf-8">
<title>Dagre Interactive Demo</title>

<script src="https://d3js.org/d3.v4.js"></script>
<script src="https://dagrejs.github.io/project/graphlib-dot/v0.6.3/graphlib-dot.js"></script>
<script src="https://dagrejs.github.io/project/dagre-d3/latest/dagre-d3.js"></script>

<style>
    svg {
        border: 1px solid #999;
        overflow: hidden;
    }

    .node {
        white-space: nowrap;
    }

    .node rect,
    .node circle,
    .node ellipse {
        stroke: #333;
        fill: #fff;
        stroke-width: 1.5px;
    }

    .cluster rect {
        stroke: #333;
        fill: #000;
        fill-opacity: 0.1;
        stroke-width: 1.5px;
    }

    .edgePath path.path {
        stroke: #333;
        stroke-width: 1.5px;
        fill: none;
    }
</style>

<style>
    h1,
    h2 {
        color: #333;
    }

    textarea {
        width: 800px;
    }

    label {
        margin-top: 1em;
        display: block;
    }

    .error {
        color: red;
    }
</style>

<body onLoad="tryDraw();">

    <h1>Dagre Interactive Demo</h1>

    <h2>Input</h2>

    <form>
        <label for="inputGraph">Graphviz Definition</label>
        <textarea id="inputGraph" rows="5" style="display: block" onKeyUp="tryDraw();">%s</textarea>

        <a id="graphLink">Link for this graph</a>
    </form>

    <h2>Graph Visualization</h2>

    <svg width=800 height=600>
        <g/>
    </svg>

    <script>
        // Input related code goes here

        function graphToURL() {
            var elems = [window.location.protocol, '//',
            window.location.host,
            window.location.pathname,
                '?'];

            var queryParams = [];
            if (debugAlignment) {
                queryParams.push('alignment=' + debugAlignment);
            }
            queryParams.push('graph=' + encodeURIComponent(inputGraph.value));
            elems.push(queryParams.join('&'));

            return elems.join('');
        }

        var inputGraph = document.querySelector("#inputGraph");

        var graphLink = d3.select("#graphLink");

        var oldInputGraphValue;

        var graphRE = /[?&]graph=([^&]+)/;
        var graphMatch = window.location.search.match(graphRE);
        if (graphMatch) {
            inputGraph.value = decodeURIComponent(graphMatch[1]);
        }
        var debugAlignmentRE = /[?&]alignment=([^&]+)/;
        var debugAlignmentMatch = window.location.search.match(debugAlignmentRE);
        var debugAlignment;
        if (debugAlignmentMatch) debugAlignment = debugAlignmentMatch[1];

        // Set up zoom support
        var svg = d3.select("svg"),
            inner = d3.select("svg g"),
            zoom = d3.zoom().on("zoom", function () {
                inner.attr("transform", d3.event.transform);
            });
        // force full width
        svg.attr("width", window.innerWidth);
        svg.call(zoom);

        // Create and configure the renderer
        var render = dagreD3.render();

        var g;
        function tryDraw() {
            console.log(inputGraph.value);
            if (oldInputGraphValue !== inputGraph.value) {
                inputGraph.setAttribute("class", "");
                oldInputGraphValue = inputGraph.value;
                try {
                    g = graphlibDot.read(inputGraph.value);
                } catch (e) {
                    inputGraph.setAttribute("class", "error");
                    throw e;
                }

                // Save link to new graph
                graphLink.attr("href", graphToURL());

                // Set margins, if not present
                if (!g.graph().hasOwnProperty("marginx") &&
                    !g.graph().hasOwnProperty("marginy")) {
                    g.graph().marginx = 20;
                    g.graph().marginy = 20;
                }

                g.graph().transition = function (selection) {
                    return selection.transition().duration(500);
                };

                // Render the graph into svg g
                d3.select("svg g").call(render, g);
            }
        }
    </script>
    `
