var index = {
    init: function() {
        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();
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
    }
};