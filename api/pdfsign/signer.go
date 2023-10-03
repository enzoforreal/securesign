package pdfsign

import (
	"log"
	"os"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
)

func SignPDF(input string, output string) error {
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

	// TODO: Replace the placeholder values with actual certificate, private key, etc.
	err = sign.Sign(inputFile, outputFile, rdr, fileInfo.Size(), sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        "John Doe",
				Location:    "Location",
				Reason:      "Reason for signing",
				ContactInfo: "Contact Info",
				Date:        time.Now().Local(),
			},
		},
		Signer:            nil, // Replace with actual private key
		DigestAlgorithm:   nil, // Replace with desired algorithm
		Certificate:       nil, // Replace with actual certificate
		CertificateChains: nil, // Replace with certificate chain
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
