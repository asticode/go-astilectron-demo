if (typeof asticode === "undefined") {
    var asticode = {};
}
asticode.loader = {
    hide: function() {
        document.getElementById("astiloader").style.display = "none";
    },
    init: function() {
        document.body.innerHTML = `
        <div class="astiloader" id="astiloader">
            <div class="astiloader-background"></div>
            <div class="astiloader-table"><div class="astiloader-content"><i class="fa fa-spinner fa-spin fa-3x fa-fw"></i></div></div>
        </div>
        ` + document.body.innerHTML
    },
    show: function() {
        document.getElementById("astiloader").style.display = "block";
    }
};