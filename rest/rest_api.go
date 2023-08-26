package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/rest/apivo"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils/requests"
	log "github.com/sirupsen/logrus"
)

const (
	V1WsAuth = "/v1/ws/auth"

	V1Orders           = "/linear/v1/orders"
	V1AmendOrders      = "/linear/v1/amend_orders"
	V1CancelOrders     = "/linear/v1/cancel_orders"
	V1NewBatchOrders   = "/linear/v1/batchorders"
	V1AmendBatchOrders = "/linear/v1/amend_batchorders"

	V1OpenOrder  = "/linear/v1/open_orders"
	V1UserTrades = "/linear/v1/user/trades"
	V1Positions  = "/linear/v1/positions"

	V1AccountConfigs = "/linear/v1/account_configs"

	V1CondOrders = "/linear/v1/conditional_orders"

	V1LeverageRatio = "/linear/v1/leverage_ratio"

	V1AggregatedPositions = "/linear/v1/aggregated/positions"
	V1AggregatedTrades    = "/linear/v1/aggregated/trades"

	V1NewBlockTrade      = "/linear/v1/blocktrades"
	V1GetBlockTrade      = "/linear/v1/blocktrades"
	V1PlatformBlockTrade = "/linear/v1/platform_blocktrades"
	V1BtUserInfo         = "/linear/v1/user/info"
)

type SdkRespPageInfo struct {
	Code     int              `json:"code"`
	Message  string           `json:"message"`
	Data     *json.RawMessage `json:"data"`
	PageInfo struct {
		HasMore bool `json:"has_more"`
	} `json:"page_info"`
}

type BitcomRestClient struct {
	baseUrl    string
	ApiKey     string
	privateKey string
	httpClient *http.Client
	VerboseLog bool
}

func NewBitcomRestClient(baseUrl, apiKey, privateKey string) (*BitcomRestClient, error) {
	return &BitcomRestClient{
		baseUrl:    baseUrl,
		ApiKey:     apiKey,
		privateKey: privateKey,
		httpClient: &http.Client{Timeout: time.Second * 5},
	}, nil
}

func MergeQueryStr(url string, paramMap map[string]any) string {
	if q := utils.CombineQueryString(paramMap); q != "" {
		url += "?" + q
	}
	return url
}

func (client *BitcomRestClient) callApiAndParseResult(method, apiPath string, paramMap map[string]any,
	extraHeaders map[string]string, result any) (bool, error) {
	if paramMap == nil {
		paramMap = make(map[string]any)
	}

	if extraHeaders == nil {
		extraHeaders = make(map[string]string)
	}

	paramMap["timestamp"] = time.Now().UnixMilli()
	signature := utils.Sign(client.privateKey, apiPath, paramMap)
	paramMap["signature"] = signature

	extraHeaders["X-Bit-Access-Key"] = client.ApiKey
	extraHeaders["Accept"] = "application/json"

	if method != http.MethodGet {
		extraHeaders["Buffer-Type"] = "application/json"
	}

	url := client.baseUrl + apiPath
	var bodyStr string
	if method != http.MethodGet {
		bodyStr = utils.AnyToJsonStr(paramMap)
	} else {
		url = MergeQueryStr(url, paramMap)
	}

	if client.VerboseLog {
		log.Infof("SEND req: method=%v, url=%v, body=%v", method, url, bodyStr)
	}

	rawResp, err := requests.DoHttp(client.httpClient, method, url, extraHeaders, bodyStr)
	if err != nil {
		return false, err
	}

	if client.VerboseLog {
		log.Infof("GetSdkResponse: request=%v %v, bodyStr=%v, raw response=%v", method, url, bodyStr, string(rawResp.Buffer))
	}

	var ret SdkRespPageInfo
	err = json.Unmarshal(rawResp.Buffer, &ret)
	if err != nil {
		return false, err
	}

	if ret.Code != 0 {
		return false, fmt.Errorf(string(rawResp.Buffer))
	}

	err = json.Unmarshal(*ret.Data, result)

	return ret.PageInfo.HasMore, err
}

func (client *BitcomRestClient) GetWsAuthToken() (string, error) {
	var data struct {
		Token string `json:"token"`
	}

	_, err := client.callApiAndParseResult(http.MethodGet, V1WsAuth, nil, nil, &data)
	if err != nil {
		return "", err
	}
	return data.Token, nil
}

func (client *BitcomRestClient) PlaceOrder(orderReq map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1Orders, orderReq,
		map[string]string{}, ordActVo)
	return
}

func (client *BitcomRestClient) NewBatchOrders(orderReq map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1NewBatchOrders, orderReq,
		map[string]string{}, batchVo)
	return
}

func (client *BitcomRestClient) AmendOrder(req map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1AmendOrders, req, nil, ordActVo)
	return
}

func (client *BitcomRestClient) AmendBatchOrders(req map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1AmendBatchOrders, req, nil, batchVo)
	return
}

func (client *BitcomRestClient) CancelOrders(req map[string]any) (cancelVo *apivo.OrderCancelVo, err error) {
	cancelVo = new(apivo.OrderCancelVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1CancelOrders, req, nil, cancelVo)
	return
}

func (client *BitcomRestClient) GetOrderHistory(req map[string]any) (orderVoList []*apivo.OrderVo, hasMore bool, err error) {
	hasMore, err = client.callApiAndParseResult(http.MethodGet, V1Orders, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) GetOpenOrders(req map[string]any) (orderVoList []*apivo.OrderVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1OpenOrder, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) GetUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1UserTrades, req, nil, &tradeVoList)
	return
}

func (client *BitcomRestClient) GetPositions(req map[string]any) (posVoList []*apivo.UsdxPositionVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1Positions, req, nil, &posVoList)
	return
}

func (client *BitcomRestClient) QueryAccountConfigs(req map[string]any) (data *apivo.UsdxUserConfigVo, err error) {
	data = new(apivo.UsdxUserConfigVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1AccountConfigs, req, nil, data)
	return
}

func (client *BitcomRestClient) QueryConditionalOrders(req map[string]any) (data []*apivo.UsdxCondOrderVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1CondOrders, req, nil, &data)
	return
}

func (client *BitcomRestClient) PlaceBlocktrades(orderReq map[string]any) (dataVo *apivo.BlockTradeResponse, err error) {
	dataVo = new(apivo.BlockTradeResponse)
	_, err = client.callApiAndParseResult(http.MethodPost, V1NewBlockTrade, orderReq,
		map[string]string{}, dataVo)
	return
}

func (client *BitcomRestClient) GetBlocktrades(req map[string]any) (data []*apivo.BlockTradeQueryVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1GetBlockTrade, req, nil, &data)
	return
}

func (client *BitcomRestClient) GetPlatformBlocktrades(req map[string]any) (data []*apivo.BlockTradeQueryVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1PlatformBlockTrade, req, nil, &data)
	return
}

func (client *BitcomRestClient) GetBlocktradeUserInfo(req map[string]any) (int64, error) {
	var userInfo struct {
		UserId string `json:"user_id"`
	}

	_, err := client.callApiAndParseResult(http.MethodGet, V1BtUserInfo, req, nil, &userInfo)
	if err != nil {
		return 0, err
	}
	return utils.MustAToInt64(userInfo.UserId), nil
}

func (client *BitcomRestClient) GetLeverageRatio(req map[string]any) (data *apivo.LeverageRatioVo, err error) {
	data = new(apivo.LeverageRatioVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1LeverageRatio, req, nil, data)
	return
}

func (client *BitcomRestClient) UpdateLeverageRatio(req map[string]any) (data *apivo.LeverageRatioVo, err error) {
	data = new(apivo.LeverageRatioVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1LeverageRatio, req,
		map[string]string{}, data)
	return
}

func (client *BitcomRestClient) GetAggregatedPositions(req map[string]any) (posVoList []*apivo.UsdxPositionVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1AggregatedPositions, req, nil, &posVoList)
	return
}

func (client *BitcomRestClient) GetAggregatedUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1AggregatedTrades, req, nil, &tradeVoList)
	return
}
