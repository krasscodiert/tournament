import m from "mithril";
import vis from "vis";

export var GraphViewComponent = {
    container: undefined,
    nodes: undefined,
    edges: undefined,
    options: undefined,
    configure(){
        this.options = {
          manipulation: {
            addNode: this.addNode,
          }
        }
    },
    addNode(nodeData, callback){
        nodeData.label = 'hello world';
        callback(nodeData);
    },
    oninit(vnode){
        this.configure();
        this.nodes =  new vis.DataSet([
                       {id: 1, label: 'Node 1'},
                       {id: 2, label: 'Node 2'},
                       {id: 3, label: 'Node 3'},
                       {id: 4, label: 'Node 4'},
                       {id: 5, label: 'Node 5'}
                   ]);
        this.edges = new vis.DataSet([
                {from: 1, to: 3},
                {from: 1, to: 2},
                {from: 2, to: 4},
                {from: 2, to: 5}
            ]);
    },
    oncreate(vnode){
         this.container = document.getElementById("graphview");
         var network = new vis.Network(this.container, {nodes: this.nodes, edges: this.edges}, this.options);
    },
    view(vnode) {
        return (
            <div className="graphview" id="graphview" >
        
            </div>
        )
    }
};