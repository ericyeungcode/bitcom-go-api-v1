package main

import (
	"os"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/rest"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils"
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
	spotCfgs, err := restClient.SpotQueryAccountConfigs(nil)
	log.Infof("SpotQueryAccountConfigs = %v, err = %v", utils.AnyToJsonStr(spotCfgs), err)

	time.Sleep(time.Second)
	openOrders, err := restClient.LinearGetOpenOrders(nil)
	log.Infof("GetOpenOrders = %v, err = %v", utils.AnyToJsonStr(openOrders), err)

	time.Sleep(time.Second)
	posList, err := restClient.LinearGetPositions(nil)
	log.Infof("GetPositions = %v, err = %v", utils.AnyToJsonStr(posList), err)

}
