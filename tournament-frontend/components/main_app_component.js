var main_app_component = {
    view: function() {
        return m("main", [
            m("h1", {class: "title"}, "My first app"),
            m("button", "A button"),
        ])
    }
}

m.mount(document.body, main_app_component)