package bitrest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/bitrest/apivo"
	"github.com/ericyeungcode/caliber"
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

func (m *BitcomMarketClient) LinearGetInstruments() ([]*apivo.UsdxInstrumentVo, error) {
	url := m.baseUrl + "/linear/v1/instruments"
	apiResp, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[[]*apivo.UsdxInstrumentVo]](m.httpClient, http.MethodGet, url, nil, "")
	if err != nil {
		return nil, err
	}
	return apiResp.Data, nil
}

func (m *BitcomMarketClient) LinearGetOrderbook(instrumentId string) (*apivo.OrderBookVo, error) {
	url := m.baseUrl + fmt.Sprintf("/linear/v1/orderbooks?instrument_id=%v", instrumentId)
	apiResp, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[apivo.OrderBookVo]](m.httpClient, http.MethodGet, url, nil, "")
	if err != nil {
		return nil, err
	}
	return &apiResp.Data, nil
}

func (m *BitcomMarketClient) LinearGetTicker(instrumentId string) (*apivo.TickerVo, error) {
	url := m.baseUrl + fmt.Sprintf("/linear/v1/tickers?instrument_id=%v", instrumentId)
	apiResp, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[apivo.TickerVo]](m.httpClient, http.MethodGet, url, nil, "")
	if err != nil {
		return nil, err
	}
	return &apiResp.Data, nil
}

func (m *BitcomMarketClient) LinearGetMarketTrades(instrumentId string) (*apivo.MarketTradeVo, error) {
	url := m.baseUrl + fmt.Sprintf("/linear/v1/market/trades?instrument_id=%v", instrumentId)
	apiResp, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[apivo.MarketTradeVo]](m.httpClient, http.MethodGet, url, nil, "")
	if err != nil {
		return nil, err
	}
	return &apiResp.Data, nil
}

func (m *BitcomMarketClient) LinearGetMarketSummary() (*apivo.MarketSummaryVo, error) {
	url := m.baseUrl + "/linear/v1/market/summary"
	apiResp, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[apivo.MarketSummaryVo]](m.httpClient, http.MethodGet, url, nil, "")
	if err != nil {
		return nil, err
	}
	return &apiResp.Data, nil
}
