import m from "mithril";

export var NavBarEntryComponent = {
    view(vnode) {
        return (
            <div className="navbarentrycomponent noselect" >
                <i className={vnode.children[2]}></i>
                <a href={vnode.children[0]}>{vnode.children[1]}</a>
            </div>
        )
    }
};
