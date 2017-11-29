let index = {
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    addFolder(name, path) {
        let div = document.createElement("div");
        div.className = "dir";
        div.onclick = function() { index.explore(path) };
        div.innerHTML = `<i class="fa fa-folder"></i><span>` + name + `</span>`;
        document.getElementById("dirs").appendChild(div)
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Explore default path
            index.explore();
        })
    },
    explore: function(path) {
        // Create message
        let message = {"name": "explore"};
        if (typeof path !== "undefined") {
            message.payload = path
        }

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            // Process path
            document.getElementById("path").innerHTML = message.payload.path;

            // Process dirs
            document.getElementById("dirs").innerHTML = ""
            index.addFolder(message.payload.path_dir.name, message.payload.path_dir.path);
            for (let i = 0; i < message.payload.dirs.length; i++) {
                index.addFolder(message.payload.dirs[i].name, message.payload.dirs[i].path);
            }

            // Process files
            document.getElementById("files").innerHTML = "";
            if (typeof message.payload.files !== "undefined") {
                let canvas = document.createElement("canvas");
                document.getElementById("files").append(canvas);
                new Chart(canvas, message.payload.files);
            } else {
                document.getElementById("files").innerHTML = "No files here...";
            }
        })
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    break;
            }
        });
    }
};