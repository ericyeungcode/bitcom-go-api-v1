package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/bitrest"
	"github.com/ericyeungcode/bitcom-go-api-v1/bitws"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils"
	"github.com/recws-org/recws"
)

var wsHost = "wss://betaws.bitexch.dev"

func ReadWebsocket(ctx context.Context, ws *recws.RecConn) {
	for {
		select {
		case <-ctx.Done():
			go ws.Close()
			log.Printf("Websocket closed %s", ws.GetURL())
			return
		default:
			if !ws.IsConnected() {
				log.Printf("Websocket disconnected %s", ws.GetURL())
				continue
			}

			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Error: ReadMessage %s", ws.GetURL())
				return
			}

			log.Printf("RECV: %s", message)
		}
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	_ = cancelFunc

	ws := recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}

	go ReadWebsocket(ctx, &ws)

	log.Printf("Connecting: %v\n", wsHost)
	ws.Dial(wsHost, nil)

	restClient, err := bitrest.NewBitcomRestClient(os.Getenv("BITCOM_REST_HOST"), os.Getenv("BITCOM_AK"), os.Getenv("BITCOM_SK"))
	if err != nil {
		log.Panic(err)
	}
	wsToken, err := restClient.GetWsAuthToken()
	if err != nil {
		log.Printf("GetWsAuthToken err: %v", err)
		return
	}

	var subscription = &bitws.PrivateSubSeq{
		SubReq: &bitws.SubReq{
			Type:     bitws.SubTypeSubscribe,
			Channels: []string{"um_account"},
			Interval: "100ms",
		},
		Token: wsToken,
	}

	req := utils.AnyToJsonStr(subscription)
	log.Printf("Sending subscription: %v\n", req)
	if err := ws.WriteMessage(1, []byte(req)); err != nil {
		log.Printf("Error: WriteMessage %s", ws.GetURL())
		return
	}

	select {}
}
