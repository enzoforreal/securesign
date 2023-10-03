function uploadPDF() {
    const fileInput = document.getElementById('pdf-upload');
    const file = fileInput.files[0];

    if (!file) {
        alert('Veuillez sélectionner un fichier PDF à uploader.');
        return;
    }

    const formData = new FormData();
    formData.append('pdf', file);

    fetch('/upload', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
    })
    .catch(error => {
        console.error('Erreur:', error);
    });
}

function signPDF() {
    const pdfName = document.getElementById('pdf-name').value;

    if (!pdfName) {
        alert('Veuillez entrer le nom du fichier PDF à signer.');
        return;
    }

    fetch(`/sign?filename=${pdfName}`, {
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
    })
    .catch(error => {
        console.error('Erreur:', error);
    });
}

function downloadPDF() {
    const signedPdfName = document.getElementById('signed-pdf-name').value;

    if (!signedPdfName) {
        alert('Veuillez entrer le nom du fichier PDF signé à télécharger.');
        return;
    }

    window.open(`/download/${signedPdfName}`, '_blank');
}
