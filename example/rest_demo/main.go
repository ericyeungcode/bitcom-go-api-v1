package main

import (
	"os"
	"time"

	"github.com/ericyeungcode/bitsdk/rest"
	"github.com/ericyeungcode/bitsdk/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	restClient, err := rest.NewBitcomRestClient(os.Getenv("BITCOM_REST_HOST"), os.Getenv("BITCOM_AK"), os.Getenv("BITCOM_SK"))
	if err != nil {
		log.Panic(err)
	}
	wsToken, err := restClient.GetWsAuthToken()
	log.Infof("GetWsAuthToken = %v, err = %v", wsToken, err)

	time.Sleep(time.Second)

	posList, err := restClient.GetPositions(nil)
	log.Infof("GetPositions = %v, err = %v", utils.AnyToJsonStr(posList), err)

}
