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
        if (data.status !== 'ok') {
            window.Telegram.WebApp.showAlert('Invalid session', () => {
                window.Telegram.WebApp.close();
            });
        }
    })
    .catch(() => {
        console.log('Unable to start session');
        window.Telegram.WebApp.showAlert('Unable to start session', () => {
            window.Telegram.WebApp.close();
        });
    });
}

function setupMenu() {
    document.getElementById("btn-dropbox").addEventListener("click", function () {
        showSection('dropbox-menu');
        /*
        const data = {};
        window.Telegram.WebApp.sendData(JSON.stringify(data));
        */
    });

    document.getElementById("btn-add-dropbox-file-request").addEventListener("click", function () {
        showSection('dropbox-add-file-request');
    });
}

function showSection(sectionId) {
    const sections = document.getElementsByTagName('section');
    console.log(sections);
    let idx = 0;
    while (idx < sections.length) {
        if (sections.item(idx).id != `section-${sectionId}`) {
            sections.item(idx).classList.add('hidden');
        } else {
            sections.item(idx).classList.remove('hidden');
        }
        idx += 1;
    }
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