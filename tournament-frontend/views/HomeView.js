import m from "mithril";

import { ContentView } from "./ContentView";

export var HomeView = {
    view(vnode) {
        return (
            <ContentView>
                Das ist die home seite
                <p></p>
                <button class="btn">yok</button>
            </ContentView>
        )
    }
};
