package pdfsign

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
)

func LoadCertificateFromEnv() (*x509.Certificate, error) {
	certPath, exists := os.LookupEnv("CERTIFICATE")
	if !exists {
		return nil, errors.New("certificate not set in .env file")
	}

	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certPEM)
	if err != nil {
		return nil, err
	}

	return x509.ParseCertificate(block.Bytes)
}

func LoadPrivateKeyFromEnv() (*rsa.PrivateKey, error) {
	keyPath, exists := os.LookupEnv("PRIVATE_KEY")
	if !exists {
		return nil, errors.New("privateKey not set in .env file")
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyPEM)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func SignPDF(input string, output string, privateKeyPath string, certPath string, chainPath string) error {

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return err
	}

	certBytes, err := os.ReadFile(certPath)
	if err != nil {
		return err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return err
	}

	var chainCerts [][]*x509.Certificate
	if chainPath != "" {
		chainBytes, err := os.ReadFile(chainPath)
		if err != nil {
			return err
		}

		intermediateCerts, err := x509.ParseCertificates(chainBytes)
		if err != nil {
			return err
		}

		chainCerts = append(chainCerts, intermediateCerts)
	}

	inputFile, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return err
	}

	rdr, err := pdf.NewReader(inputFile, fileInfo.Size())
	if err != nil {
		return err
	}

	err = sign.Sign(inputFile, outputFile, rdr, fileInfo.Size(), sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        "DIGNAT RENAUD",
				Location:    "Location",
				Reason:      "Reason for signing",
				ContactInfo: "Contact Info",
				Date:        time.Now().Local(),
			},
		},
		Signer:            privateKey,
		DigestAlgorithm:   crypto.SHA256,
		Certificate:       cert,
		CertificateChains: chainCerts,
		TSA: sign.TSA{
			URL: "https://freetsa.org/tsr",
		},
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
