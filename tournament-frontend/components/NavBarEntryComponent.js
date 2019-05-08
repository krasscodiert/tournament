import m from "mithril";


export let NavBarEntryComponent = {
    view(vnode) {
        return (
            <div className="navbarentrycomponent" >
                <i className={vnode.children[2]}></i>
                <a href={vnode.children[0]}>{vnode.children[1]}</a>
            </div>
        )
    }
};
