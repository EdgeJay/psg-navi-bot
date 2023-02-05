function init() {
    document.getElementById("btn-dropbox").addEventListener("click", function () {
        console.log("dropbox!");
    });
}

function onLoad() {
    console.log("window loaded");
    window.removeEventListener("load", onLoad);
    feather.replace();
    init();
}

window.addEventListener("load", onLoad);