package unlynx

import (
	"github.com/lca1/medco-connector/util"
	servicesmedco "github.com/lca1/medco-unlynx/services"
	libunlynx "github.com/lca1/unlynx/lib"
	"github.com/sirupsen/logrus"
	"go.dedis.ch/onet/v3"
	"go.dedis.ch/onet/v3/app"
	"os"
	"strconv"
)

// serializeCipherVector serializes a vector of cipher texts into a string-encoded slice
func serializeCipherVector(cipherVector libunlynx.CipherVector) (serializedVector []string) {
	for _, cipherText := range cipherVector {
		serializedVector = append(serializedVector, cipherText.Serialize())
	}
	return
}

// deserializeCipherVector deserializes string-encoded cipher texts into a vector
func deserializeCipherVector(cipherTexts []string) (cipherVector libunlynx.CipherVector, err error) {
	for _, cipherText := range cipherTexts {
		deserialized := libunlynx.CipherText{}
		err = deserialized.Deserialize(cipherText)
		if err != nil {
			logrus.Error("unlynx error deserializing cipher text: ", err)
			return
		}

		cipherVector = append(cipherVector, deserialized)
	}
	return
}

// newUnlynxClient creates a new client to communicate with unlynx
func newUnlynxClient() (unlynxClient *servicesmedco.API, cothorityRoster *onet.Roster) {

	// initialize medco client
	groupFile, err := os.Open(util.UnlynxGroupFilePath)
	if err != nil {
		logrus.Panic("unlynx error opening group file: ", err)
	}

	group, err := app.ReadGroupDescToml(groupFile)
	if err != nil || len(group.Roster.List) <= 0 {
		logrus.Panic("unlynx error parsing group file: ", err)
	}

	cothorityRoster = group.Roster
	unlynxClient = servicesmedco.NewMedCoClient(
		cothorityRoster.List[util.UnlynxGroupFileIdx],
		strconv.Itoa(util.UnlynxGroupFileIdx),
	)

	return
}
