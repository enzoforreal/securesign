package pdfsign

import (
	"errors"
	"os"

	"github.com/digitorus/pdfsign/verify"
)

func VerifyPDF(input string) (bool, *verify.Response, error) {
	inputFile, err := os.Open(input)
	if err != nil {
		return false, nil, err
	}
	defer inputFile.Close()

	apiResp, err := verify.File(inputFile)
	if err != nil {
		return false, nil, err
	}

	if apiResp.Error != "" {
		return false, apiResp, errors.New(apiResp.Error)
	}

	isValid := len(apiResp.Signers) > 0 && apiResp.Signers[0].ValidSignature
	return isValid, apiResp, nil
}
