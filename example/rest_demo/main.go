package main

import (
	"os"

	"github.com/ericyeungcode/bitsdk/rest"
	"github.com/ericyeungcode/bitsdk/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	restClient, err := rest.NewBitcomRestClient(os.Getenv("BITCOM_REST_HOST"), os.Getenv("BITCOM_AK"), os.Getenv("BITCOM_SK"))
	if err != nil {
		log.Panic(err)
	}
	posList, err := restClient.GetPositions(nil)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	log.Infof("positions = %v", utils.AnyToJsonStr(posList))
}
