function startSession() {
    // init session and get jwt token stored in cookies
    fetch('/api/init-menu-session', {
        method: 'POST',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
            'X-PSGNaviBot-Csrf-Token': window.__token,
        },
        body: JSON.stringify({
            init_data: window.Telegram.WebApp.initData,
        })
    })
    .then((res) => res.json())
    .then((data) => {
        console.log(data);
    });
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