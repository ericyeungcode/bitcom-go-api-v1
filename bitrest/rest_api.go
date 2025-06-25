package bitrest

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ericyeungcode/bitcom-go-api-v1/bitrest/apivo"
	"github.com/ericyeungcode/bitcom-go-api-v1/utils"
	"github.com/ericyeungcode/caliber/request"
	log "github.com/sirupsen/logrus"
)

const (
	V1WsAuth = "/v1/ws/auth"

	// UM
	V1_UM_ACCOUNT_MODE     = "/um/v1/account_mode"
	V1_UM_ACCOUNTS         = "/um/v1/accounts"
	V1_UM_ACCOUNT_CONFIGS  = "/um/v1/um_account_configs"
	V1_UM_TRANSACTIONS     = "/um/v1/transactions"
	V1_UM_INTEREST_RECORDS = "/um/v1/interest_records"

	// SPOT
	V1_SPOT_ORDERS              = "/spot/v1/orders"
	V1_SPOT_CANCEL_ORDERS       = "/spot/v1/cancel_orders"
	V1_SPOT_OPENORDERS          = "/spot/v1/open_orders"
	V1_SPOT_USER_TRADES         = "/spot/v1/user/trades"
	V1_SPOT_AMEND_ORDERS        = "/spot/v1/amend_orders"
	V1_SPOT_WS_AUTH             = "/spot/v1/ws/auth"
	V1_SPOT_BATCH_ORDERS        = "/spot/v1/batchorders"
	V1_SPOT_AMEND_BATCH_ORDERS  = "/spot/v1/amend_batchorders"
	V1_SPOT_MMP_STATE           = "/spot/v1/mmp_state"
	V1_SPOT_MMP_UPDATE_CONFIG   = "/spot/v1/update_mmp_config"
	V1_SPOT_RESET_MMP           = "/spot/v1/reset_mmp"
	V1_SPOT_ACCOUNT_CONFIGS_COD = "/spot/v1/account_configs/cod"
	V1_SPOT_ACCOUNT_CONFIGS     = "/spot/v1/account_configs"
	V1_SPOT_AGG_TRADES          = "/spot/v1/aggregated/trades"

	// Linear
	V1_Linear_Orders           = "/linear/v1/orders"
	V1_Linear_AmendOrders      = "/linear/v1/amend_orders"
	V1_Linear_CancelOrders     = "/linear/v1/cancel_orders"
	V1_Linear_NewBatchOrders   = "/linear/v1/batchorders"
	V1_Linear_AmendBatchOrders = "/linear/v1/amend_batchorders"

	V1_Linear_OpenOrder  = "/linear/v1/open_orders"
	V1_Linear_UserTrades = "/linear/v1/user/trades"
	V1_Linear_Positions  = "/linear/v1/positions"

	V1_Linear_AccountConfigs = "/linear/v1/account_configs"

	// TODO: mmp

	V1_Linear_CondOrders = "/linear/v1/conditional_orders"

	V1_Linear_LeverageRatio = "/linear/v1/leverage_ratio"

	V1_Linear_AggregatedPositions = "/linear/v1/aggregated/positions"
	V1_Linear_AggregatedTrades    = "/linear/v1/aggregated/trades"

	V1_Linear_NewBlockTrade      = "/linear/v1/blocktrades"
	V1_Linear_GetBlockTrade      = "/linear/v1/blocktrades"
	V1_Linear_PlatformBlockTrade = "/linear/v1/platform_blocktrades"
	V1_Linear_BtUserInfo         = "/linear/v1/user/info"
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

	rawResp, err := request.HttpRequest(client.httpClient, method, url, extraHeaders, bodyStr)
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
		return false, errors.New(string(rawResp.Buffer))
	}

	err = json.Unmarshal(*ret.Data, result)

	return ret.PageInfo.HasMore, err
}

//////////////////
// Ws
//////////////////

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

func (client *BitcomRestClient) SpotGetWsAuthToken() (string, error) {
	var data struct {
		Token string `json:"token"`
	}

	_, err := client.callApiAndParseResult(http.MethodGet, V1_SPOT_WS_AUTH, nil, nil, &data)
	if err != nil {
		return "", err
	}
	return data.Token, nil
}

//////////////////
// Um account etc..
//////////////////

func (client *BitcomRestClient) QueryAccountMode(paramMap map[string]any) (*apivo.AccountModeVo, error) {
	var result apivo.AccountModeVo
	_, err := client.callApiAndParseResult(http.MethodGet, V1_UM_ACCOUNT_MODE, paramMap, nil, &result)
	return &result, err
}

func (client *BitcomRestClient) QueryUmAccount(paramMap map[string]any) (*apivo.UmAccountVo, error) {
	var result apivo.UmAccountVo
	_, err := client.callApiAndParseResult(http.MethodGet, V1_UM_ACCOUNTS, paramMap, nil, &result)
	return &result, err
}

func (client *BitcomRestClient) QueryUmTxLogs(paramMap map[string]any) (bool, []*apivo.UmTxLogVo, error) {
	var result []*apivo.UmTxLogVo
	hasMore, err := client.callApiAndParseResult(http.MethodGet, V1_UM_TRANSACTIONS, paramMap, nil, &result)
	return hasMore, result, err
}

func (client *BitcomRestClient) QueryUmAccountConfigs(paramMap map[string]any) (*apivo.UmAccountConfigVo, error) {
	var result apivo.UmAccountConfigVo
	_, err := client.callApiAndParseResult(http.MethodGet, V1_UM_ACCOUNT_CONFIGS, paramMap, nil, &result)
	return &result, err
}

func (client *BitcomRestClient) QueryInterests(paramMap map[string]any) (bool, []*apivo.InterestVo, error) {
	var result []*apivo.InterestVo
	hasMore, err := client.callApiAndParseResult(http.MethodGet, V1_UM_INTEREST_RECORDS, paramMap, nil, &result)
	return hasMore, result, err
}

//////////////////
// Spot
//////////////////

func (client *BitcomRestClient) SpotPlaceOrder(orderReq map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_SPOT_ORDERS, orderReq,
		map[string]string{}, ordActVo)
	return
}

func (client *BitcomRestClient) SpotNewBatchOrders(orderReq map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_SPOT_BATCH_ORDERS, orderReq,
		map[string]string{}, batchVo)
	return
}

func (client *BitcomRestClient) SpotAmendOrder(req map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_SPOT_AMEND_ORDERS, req, nil, ordActVo)
	return
}

func (client *BitcomRestClient) SpotAmendBatchOrders(req map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_SPOT_AMEND_BATCH_ORDERS, req, nil, batchVo)
	return
}

func (client *BitcomRestClient) SpotCancelOrders(req map[string]any) (cancelVo *apivo.OrderCancelVo, err error) {
	cancelVo = new(apivo.OrderCancelVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_SPOT_CANCEL_ORDERS, req, nil, cancelVo)
	return
}

func (client *BitcomRestClient) SpotGetOrderHistory(req map[string]any) (orderVoList []*apivo.OrderVo, hasMore bool, err error) {
	hasMore, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_ORDERS, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) SpotGetOpenOrders(req map[string]any) (orderVoList []*apivo.OrderVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_OPENORDERS, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) SpotGetUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_USER_TRADES, req, nil, &tradeVoList)
	return
}

func (client *BitcomRestClient) SpotGetAggregatedUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_AGG_TRADES, req, nil, &tradeVoList)
	return
}

func (client *BitcomRestClient) SpotQueryAccountConfigs(req map[string]any) (data *apivo.SpotAccountConfigVo, err error) {
	data = new(apivo.SpotAccountConfigVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_ACCOUNT_CONFIGS, req, nil, data)
	return
}

func (client *BitcomRestClient) SpotQueryMmpState(req map[string]any) (data *apivo.MmpVo, err error) {
	data = new(apivo.MmpVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1_SPOT_MMP_STATE, req, nil, data)
	return
}

func (client *BitcomRestClient) SpotUpdateMmpConfig(req map[string]any) error {
	data := new(string)
	_, err := client.callApiAndParseResult(http.MethodPost, V1_SPOT_MMP_UPDATE_CONFIG, req, nil, data)
	return err
}

func (client *BitcomRestClient) SpotResetMmp(req map[string]any) error {
	data := new(string)
	_, err := client.callApiAndParseResult(http.MethodPost, V1_SPOT_RESET_MMP, req, nil, data)
	return err
}

func (client *BitcomRestClient) SpotQueryCodConfig() (bool, error) {
	data := &struct {
		Cod bool `json:"cod"`
	}{}
	_, err := client.callApiAndParseResult(http.MethodGet, V1_SPOT_ACCOUNT_CONFIGS_COD, nil, nil, data)
	return data.Cod, err
}

func (client *BitcomRestClient) SpotUpdateCodConfig(req map[string]any) error {
	data := new(string)
	_, err := client.callApiAndParseResult(http.MethodPost, V1_SPOT_ACCOUNT_CONFIGS_COD, req, nil, data)
	return err
}

//////////////////
// Linear
//////////////////

func (client *BitcomRestClient) LinearPlaceOrder(orderReq map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_Orders, orderReq,
		map[string]string{}, ordActVo)
	return
}

func (client *BitcomRestClient) LinearNewBatchOrders(orderReq map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_NewBatchOrders, orderReq,
		map[string]string{}, batchVo)
	return
}

func (client *BitcomRestClient) LinearAmendOrder(req map[string]any) (ordActVo *apivo.OrderActionVo, err error) {
	ordActVo = new(apivo.OrderActionVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_AmendOrders, req, nil, ordActVo)
	return
}

func (client *BitcomRestClient) LinearAmendBatchOrders(req map[string]any) (batchVo *apivo.UsdxBatchVo, err error) {
	batchVo = new(apivo.UsdxBatchVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_AmendBatchOrders, req, nil, batchVo)
	return
}

func (client *BitcomRestClient) LinearCancelOrders(req map[string]any) (cancelVo *apivo.OrderCancelVo, err error) {
	cancelVo = new(apivo.OrderCancelVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_CancelOrders, req, nil, cancelVo)
	return
}

func (client *BitcomRestClient) LinearGetOrderHistory(req map[string]any) (orderVoList []*apivo.OrderVo, hasMore bool, err error) {
	hasMore, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_Orders, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) LinearGetOpenOrders(req map[string]any) (orderVoList []*apivo.OrderVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_OpenOrder, req, nil, &orderVoList)
	return
}

func (client *BitcomRestClient) LinearGetUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_UserTrades, req, nil, &tradeVoList)
	return
}

func (client *BitcomRestClient) LinearGetPositions(req map[string]any) (posVoList []*apivo.UsdxPositionVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_Positions, req, nil, &posVoList)
	return
}

func (client *BitcomRestClient) LinearQueryAccountConfigs(req map[string]any) (data *apivo.UsdxUserConfigVo, err error) {
	data = new(apivo.UsdxUserConfigVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_AccountConfigs, req, nil, data)
	return
}

func (client *BitcomRestClient) LinearQueryConditionalOrders(req map[string]any) (data []*apivo.UsdxCondOrderVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_CondOrders, req, nil, &data)
	return
}

func (client *BitcomRestClient) LinearPlaceBlocktrades(orderReq map[string]any) (dataVo *apivo.BlockTradeResponse, err error) {
	dataVo = new(apivo.BlockTradeResponse)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_NewBlockTrade, orderReq,
		map[string]string{}, dataVo)
	return
}

func (client *BitcomRestClient) LinearGetBlocktrades(req map[string]any) (data []*apivo.BlockTradeQueryVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_GetBlockTrade, req, nil, &data)
	return
}

func (client *BitcomRestClient) LinearGetPlatformBlocktrades(req map[string]any) (data []*apivo.BlockTradeQueryVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_PlatformBlockTrade, req, nil, &data)
	return
}

func (client *BitcomRestClient) LinearGetBlocktradeUserInfo(req map[string]any) (int64, error) {
	var userInfo struct {
		UserId string `json:"user_id"`
	}

	_, err := client.callApiAndParseResult(http.MethodGet, V1_Linear_BtUserInfo, req, nil, &userInfo)
	if err != nil {
		return 0, err
	}
	return utils.MustAToInt64(userInfo.UserId), nil
}

func (client *BitcomRestClient) LinearGetLeverageRatio(req map[string]any) (data *apivo.LeverageRatioVo, err error) {
	data = new(apivo.LeverageRatioVo)
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_LeverageRatio, req, nil, data)
	return
}

func (client *BitcomRestClient) LinearUpdateLeverageRatio(req map[string]any) (data *apivo.LeverageRatioVo, err error) {
	data = new(apivo.LeverageRatioVo)
	_, err = client.callApiAndParseResult(http.MethodPost, V1_Linear_LeverageRatio, req,
		map[string]string{}, data)
	return
}

func (client *BitcomRestClient) LinearGetAggregatedPositions(req map[string]any) (posVoList []*apivo.UsdxPositionVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_AggregatedPositions, req, nil, &posVoList)
	return
}

func (client *BitcomRestClient) LinearGetAggregatedUserTrades(req map[string]any) (tradeVoList []*apivo.TradeVo, err error) {
	_, err = client.callApiAndParseResult(http.MethodGet, V1_Linear_AggregatedTrades, req, nil, &tradeVoList)
	return
}
