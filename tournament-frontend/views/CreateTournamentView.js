import m from "mithril";

import { ContentView } from "./ContentView";
import {GraphViewComponent} from "../components/GraphViewComponent"
import {util} from "../model/util";

export var CreateTournamentView = {
    players: [],
    addPlayer(){
        this.players.push(util.getRandomCSSColor());
    },
    removePlayer(){
        this.players.pop();
    },
    oninit(vnode){
        for(let i = 0; i < 5; i++){
            this.players[i] = util.getRandomCSSColor();
        }
    },
    oncreate(vnode){
        $("#playerlist").on('scroll', function(){
            this.scrolled=true;
        });
        //setInterval(updateScroll,1000);
    },
    scrolled : false,
    updateScroll(){
        if(!this.scrolled){
            var element = document.getElementById("playerlist");
            element.scrollTop = element.scrollHeight;
        }
    },
    view(vnode) {
        let x = 0;
        return (
            <ContentView>
                <div className="toolbar">
                    <div className="tournamentpanel">

                    </div>
                    <div className="playerspanel">
                        <div className="header noselect">
                            <h1>Playerlist</h1>
                        </div>
                        <div className="playerlist" id="playerlist">
                            {this.players.map(p => {
                                return(
                                    <span className="chip noselect" style={"background-color:" + p +";"}>Player {++x}</span>
                                )
                            })}
                        </div>
                        <div className="footer">
                            <div className="content">
                                <button className="btn btn-primary" onclick={this.addPlayer.bind(this)}>Add Player</button>
                                <button className="btn btn-primary" onclick={this.removePlayer.bind(this)}>Remove Player</button>
                            </div>
                        </div>
                    </div>
                </div>
                <GraphViewComponent/>
            </ContentView>
        )
    }
};
