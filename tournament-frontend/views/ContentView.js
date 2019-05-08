import "spectre.css/dist/spectre.min.css";
import "spectre.css/dist/spectre-exp.min.css";
import "spectre.css/dist/spectre-icons.min.css";

import m from "mithril";

import { NavBarView } from "./NavBarView";

export var ContentView = {
    view(vnode){
        return [
            <NavBarView/>,
            <div className="contentview">
                {vnode.children}
            </div>
        ];
    }
}