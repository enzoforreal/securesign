function uploadPDF() {
    let fileInput = document.getElementById('pdf-upload');
    let formData = new FormData();
    formData.append('pdf', fileInput.files[0]);

    fetch('/upload', {
        method: 'POST',
        body: formData
    })
    .then(response => response.text())
    .then(data => alert(data))
    .catch(error => alert("Erreur lors de l'upload: " + error));
}

function signPDF() {
    let filename = document.getElementById('pdf-name').value;

    fetch('/sign', {
        method: 'POST',
        body: JSON.stringify({ filename: filename }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.text())
    .then(data => alert(data))
    .catch(error => alert("Erreur lors de la signature: " + error));
}

function downloadPDF() {
    let filename = document.getElementById('signed-pdf-name').value;
    window.location.href = "/download/" + filename;
}
