function startSession() {
    window.Telegram.WebApp.showAlert(window.__token);
    /*
    window.cookieStore.get("psg_navi_bot_session")
        .then((data) => {
            window.Telegram.WebApp.showAlert(JSON.stringify(data))
        })
        .catch((err) => {
            window.Telegram.WebApp.showAlert(err.message)
        });
    */
}

function setupMenu() {
    document.getElementById("btn-dropbox").addEventListener("click", function () {
        /*
        const data = {};
        window.Telegram.WebApp.sendData(JSON.stringify(data));
        */
    });
}

function init() {
    startSession();
    setupMenu();
}

function onLoad() {
    console.log("window loaded");
    window.removeEventListener("load", onLoad);
    feather.replace();
    init();
}

window.addEventListener("load", onLoad);