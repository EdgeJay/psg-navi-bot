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
    });

    document.getElementById("btn-add-dropbox-file-request").addEventListener("click", function () {
        showSection('dropbox-add-file-request');
    });
}

function setupForms() {
    const formDbx = document.getElementById('form-add-dropbox-file-request')

    formDbx.addEventListener('submit', function (evt) {
        evt.preventDefault();
        
        formDbx.setAttribute('aria-busy', 'true');

        fetch('/api/dbx-add-file-request', {
            method: 'POST',
            cache: 'no-cache',
            credentials: 'same-origin',
            headers: {
                'Content-Type': 'application/json',
                'X-PSGNaviBot-Csrf-Token': window.__token,
            },
            body: JSON.stringify({
                title: document.getElementById('txt-filerequest-title').value,
                desc: document.getElementById('txt-filerequest-desc').value,
            })
        })
        .then((res) => res.json())
        .then((data) => {
            if (data.status !== 'ok') {
                throw new Error('Add new file request failed');
            }

            window.Telegram.WebApp.showAlert('File request created', () => {
                showSection('dropbox-menu');
            });
        })
        .catch(() => {
            formDbx.setAttribute('aria-busy', 'true');
            console.log('Unable to add new file request');
            window.Telegram.WebApp.showAlert('Unable to add new file request', () => {});
        });
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
    setupForms();
}

function onLoad() {
    console.log("window loaded");
    window.removeEventListener("load", onLoad);
    feather.replace();
    init();
}

window.addEventListener("load", onLoad);