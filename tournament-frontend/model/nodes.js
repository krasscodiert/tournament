var Nodes = {
    createNode() {
        return {
            type: undefined,
            value: undefined,
            graphData: undefined,
            connections: [],
            validate: undefined
        }
    },
    NodeFactory : {
        CreatePlayerNode(playerName, r, g, b){
            let node = this.createNode();
            node.type = "player";
            node.value = {
                PlayerName: playerName,
                Color: []
            };
            return node;
        },
        CreateEndNode(){
            let node = this.createNode();
            node.type = "end";
            node.value = "?";
            return node;
        },
        CreateMatchNode(){
            let node = this.createNode();
            node.type = "match";
            node.value = undefined;
            return node;
        }
    }
};