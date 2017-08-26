var index = {
    init: function() {
        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Refresh list
            index.refreshList();
        })
    },
    listen: function() {
        astilectron.listen(function(message) {
            switch (message.name) {
                case "set.style":
                    index.listenSetStyle(message);
                    break;
            }
        });
    },
    listenSetStyle: function(message) {
        document.body.className = message.payload;
    },
    refreshList: function() {
        astilectron.send({"name": "get.list"}, function(message) {
            if (message.payload.length === 0) {
                return
            }
            let c = `<ul>`
            for (let i = 0; i < message.payload.length; i++) {
                c += `<li class="` + message.payload[i].type + `">` + message.payload[i].name + `</li>`
            }
            c += `</ul>`
            document.getElementById("list").innerHTML = c
        })
    }
};