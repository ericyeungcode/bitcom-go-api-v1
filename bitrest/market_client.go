package bitrest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/bitrest/apivo"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils/requests"
)

type BitcomMarketClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewBitcomMarketClient(baseUrl string) *BitcomMarketClient {
	return &BitcomMarketClient{
		baseUrl:    baseUrl,
		httpClient: &http.Client{Timeout: time.Second * 5},
	}
}

func (m *BitcomMarketClient) LinearGetOrderbook(instrumentId string) (orderbook *apivo.OrderBookVo, err error) {
	url := m.baseUrl + fmt.Sprintf("/linear/v1/orderbooks?instrument_id=%v", instrumentId)
	orderbook = new(apivo.OrderBookVo)
	err = requests.DoHttpPayload(m.httpClient, http.MethodGet, url, nil, "", &orderbook)
	return
}
