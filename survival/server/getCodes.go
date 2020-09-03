package survivalserver

import (
	"github.com/ldsec/medco-connector/wrappers/i2b2"
	"github.com/pkg/errors"
)

//TODO use I2B2 ONT messaging

func GetCode(path string) (string, error) {
	res, err := i2b2.GetOntologyTermInfo(path)
	if err != nil {
		return "", err
	}
	if len(res) != 1 {
		return "", errors.Errorf("Result length of GetOntologyTermInfo is expected to be 1. Got: %d", len(res))
	}

	if res[0].Code == "" {
		return "", errors.New("Code is empty")
	}

	return res[0].Code, nil

}
