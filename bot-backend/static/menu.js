function getSessionInfo() {
    window.cookieStore.get("session_info")
        .then((data) => {
            window.Telegram.WebApp.showPopup(JSON.stringify(data))
        })
        .catch((err) => {
            window.Telegram.WebApp.showPopup(err.message)
        });
}

function setupMenu() {
    document.getElementById("btn-dropbox").addEventListener("click", function () {
        const data = {};
        window.Telegram.WebApp.sendData(JSON.stringify(data));
    });
}

function init() {
    getSessionInfo();
    setupMenu();
}

function onLoad() {
    console.log("window loaded");
    window.removeEventListener("load", onLoad);
    feather.replace();
    init();
}

window.addEventListener("load", onLoad);