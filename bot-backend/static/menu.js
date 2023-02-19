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

function fetchListFileRequests() {
    fetch('/api/dbx-list-file-requests', {
        method: 'GET',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
            'X-PSGNaviBot-Csrf-Token': window.__token,
        },
    })
    .then((res) => res.json())
    .then((data) => {
        console.log(data);
    })
    .catch(() => {
        window.Telegram.WebApp.showAlert('Unable to fetch list of file requests', () => {
            showSection('dropbox-menu');
        });
    });
}

function setupMenu() {
    document.getElementById("btn-dropbox").addEventListener("click", () => {
        showSection('dropbox-menu');
    });

    document.getElementById("btn-add-dropbox-file-request").addEventListener("click", () => {
        showSection('dropbox-add-file-request');
    });

    document.getElementById('btn-list-dropbox-file-requests').addEventListener("click", () => {
        showSection('dropbox-list-file-requests');
        fetchListFileRequests();
    });

    document.getElementById('btn-exit-dropbox').addEventListener('click', (evt) => {
        evt.preventDefault();
        resetForm('form-add-dropbox-file-request', true);
        showSection('menu');
    });

    document.getElementById('btn-cancel-dropbox-add-file-request').addEventListener('click', (evt) => {
        evt.preventDefault();
        resetForm('form-add-dropbox-file-request', true);
        showSection('dropbox-menu');
    });
}

function resetForm(id, clearInputs) {
    const aForm = document.getElementById(id);
    aForm.setAttribute('aria-busy', 'false');
    aForm.setAttribute('disabled', 'false');

    if (clearInputs) {
        switch (id) {
            case 'form-add-dropbox-file-request':
                document.getElementById('txt-filerequest-title').value = '';
                document.getElementById('txt-filerequest-desc').value = '';
                break;
        }
    }
}

function setupForms() {
    const formDbx = document.getElementById('form-add-dropbox-file-request')

    formDbx.addEventListener('submit', function (evt) {
        evt.preventDefault();
        
        formDbx.setAttribute('aria-busy', 'true');
        formDbx.setAttribute('disabled', 'true');

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
                resetForm('form-add-dropbox-file-request', true);
            });
        })
        .catch(() => {
            window.Telegram.WebApp.showAlert('Unable to add new file request', () => {
                resetForm('form-add-dropbox-file-request', false);
            });
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