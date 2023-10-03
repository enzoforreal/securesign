package handlers

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
	"github.com/gin-gonic/gin"
)

const (
	signedDirectory = "./static/storage/signed"
	uploadDirectory = "./static/storage/uploads"
)

func SignHandler(c *gin.Context) {
	fmt.Println("Starting the signing process...")

	certPath, privKeyPath, err := getCertAndKeyPaths()
	if err != nil {
		handleInternalError(c, "Error getting cert and key paths", err)
		return
	}

	filename := c.Query("filename")
	fmt.Printf("Received filename: %s\n", filename)
	if !isValidFilename(filename) {
		c.String(http.StatusBadRequest, "Invalid file name.")
		return
	}

	inputPath := filepath.Join(uploadDirectory, filename)
	inputFile, info, err := openAndStatFile(inputPath)
	if err != nil {
		handleInternalError(c, "Error opening and stating file", err)
		return
	}
	defer inputFile.Close()

	if info.IsDir() {
		handleInternalError(c, fmt.Sprintf("The given path %s is a directory, not a file!", inputPath), err)
		return
	}

	if info.Size() == 0 {
		handleInternalError(c, fmt.Sprintf("The file %s is empty!", inputPath), err)
		return
	}

	headerBytes := make([]byte, 10)
	_, err = inputFile.Read(headerBytes)
	if err != nil {
		handleInternalError(c, "Error reading file header", err)
		return
	}
	fmt.Printf("File header: %x\n", headerBytes)

	_, err = inputFile.Seek(0, 0)
	if err != nil {
		handleInternalError(c, "Error seeking to the beginning of the file", err)
		return
	}

	rdr, err := pdf.NewReader(inputFile, info.Size())
	if err != nil {
		handleInternalError(c, "Error creating PDF reader", err)
		return
	}

	certificate, err := loadCertificate(certPath)
	if err != nil {
		handleInternalError(c, "Error loading certificate", err)
		return
	}

	privateKey, err := loadPrivateKey(privKeyPath)
	if err != nil {
		handleInternalError(c, "Error loading private key", err)
		return
	}

	outputPath := filepath.Join(signedDirectory, filename)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		handleInternalError(c, "Error creating signed file", err)
		return
	}
	defer func() {
		outputFile.Close()
		if err != nil {
			os.Remove(outputPath)
		}
	}()

	signData := createSignData(certificate, privateKey)
	err = sign.Sign(inputFile, outputFile, rdr, info.Size(), *signData)
	if err != nil {
		handleInternalError(c, "Error signing the file", err)
		return
	}

	fmt.Println("File signed successfully!")
	c.JSON(http.StatusOK, gin.H{
		"message": "File signed successfully!",
	})
}

func handleInternalError(c *gin.Context, msg string, err error) {
	fmt.Printf("%s: %v\n", msg, err)
	c.String(http.StatusInternalServerError, msg+": %v", err)
}

func isValidFilename(filename string) bool {
	return !strings.Contains(filename, "..") && strings.HasSuffix(filename, ".pdf")
}

func getCertAndKeyPaths() (string, string, error) {
	certPath := "./newcertificate.pem"
	privKeyPath := "./newprivatekey.pem"

	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return "", "", errors.New("certificate file does not exist")
	}

	if _, err := os.Stat(privKeyPath); os.IsNotExist(err) {
		return "", "", errors.New("private key file does not exist")
	}

	return certPath, privKeyPath, nil
}

func openAndStatFile(path string) (*os.File, os.FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening the file: %v", err)
	}

	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, nil, fmt.Errorf("error getting file info: %v", err)
	}
	return file, info, nil
}

func loadCertificate(path string) (*x509.Certificate, error) {
	certPEMBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading certificate: %v", err)
	}

	block, _ := pem.Decode(certPEMBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the certificate")
	}

	return x509.ParseCertificate(block.Bytes)
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyPEMBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %v", err)
	}

	block, _ := pem.Decode(keyPEMBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not of type *rsa.PrivateKey")
	}

	return rsaPrivateKey, nil
}

func createSignData(certificate *x509.Certificate, privateKey *rsa.PrivateKey) *sign.SignData {
	return &sign.SignData{
		ObjectId: 1,
		Signature: sign.SignDataSignature{
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.DoNotAllowAnyChangesPerms,
			Info: sign.SignDataSignatureInfo{
				Name:        "Dignat Renaud",
				Location:    "Pointe Noire",
				Reason:      "Je signe ce document",
				ContactInfo: "renaud@datamix.io",
				Date:        time.Now(),
			},
		},
		DigestAlgorithm: crypto.SHA256,
		Certificate:     certificate,
		Signer:          privateKey,
	}
}
