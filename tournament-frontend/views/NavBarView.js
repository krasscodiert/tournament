import m from "mithril";

import {NavBarEntryComponent} from "../components/NavBarEntryComponent";

export var NavBarView = {
    entries : [],
    createEntry(d, h, i){
        return {
            desc: d,
            href: h,
            icon: i
        }
    },
    addEntry(d, h, i){
        if (i == undefined){
            i = "icon icon-arrow-right"
        }
        let entry = this.createEntry(d, h, i)

        this.entries.push(entry)
    },
    oninit(vnode){
        this.entries = []
        this.addEntry("Home", "/#!/home", "icon icon-location")
        this.addEntry("Create Tournament", "/#!/create-tournament")
        this.addEntry("Tournaments", "/#!/tournaments")
        this.addEntry("Join by Code!", "/#!/joinbycode")
    },
    view(vnode){
        return(
            <div className="navbar">
                <div className="navbarcontent">
                    {this.entries.map(e => {
                        return (
                            <NavBarEntryComponent>
                                {e.href}
                                {e.desc}
                                {e.icon}
                            </NavBarEntryComponent>
                        )
                    })
                    }
                </div>
            </div>
        )
    }
}
