package main

import (
	"os"

	"github.com/ericyeungcode/bitcom-go-api-v1/bitrest"
	"github.com/ericyeungcode/caliber/common"
	log "github.com/sirupsen/logrus"
)

func main() {
	marketClient := bitrest.NewBitcomMarketClient(os.Getenv("BITCOM_REST_HOST"))

	instList, err := marketClient.LinearGetInstruments()
	log.Infof("instList = %+v, err = %v", common.MustMarshalStr(instList), err)

	orderbook, err := marketClient.LinearGetOrderbook("BTC-USDT-PERPETUAL")
	log.Infof("orderbook = %+v, err = %v", orderbook, err)

	ticker, err := marketClient.LinearGetTicker("BTC-USDT-PERPETUAL")
	log.Infof("ticker = %+v, err = %v", ticker, err)
}
