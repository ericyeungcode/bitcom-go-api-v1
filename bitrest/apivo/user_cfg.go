package apivo

type UsdxUserPairConfigVoItem struct {
	BaseCcy          string `json:"base_ccy"`
	QuoteCcy         string `json:"quote_ccy"`
	BaseDisplayName  string `json:"base_display_name"`
	QuoteDisplayName string `json:"quote_display_name"`

	////////////
	CustomizeFutureFeeRates bool   `json:"customize_future_fee_rates"`
	PerpetualTakerFeeRate   string `json:"perpetual_taker_fee_rate"`
	PerpetualMakerFeeRate   string `json:"perpetual_maker_fee_rate"`

	CustomizeOptionFeeRates bool   `json:"customize_option_fee_rates"`
	OptionTakerFeeRate      string `json:"option_taker_fee_rate"`
	OptionMakerFeeRate      string `json:"option_maker_fee_rate"`

	//------------------- Blocktrades -------------------------------
	CustomizeBlocktrades          bool   `json:"customize_blocktrades"`
	BlocktradeFutureMinOrderPrice string `json:"blocktrade_future_min_order_price"`
	BlocktradeFutureMaxOrderPrice string `json:"blocktrade_future_max_order_price"`
	BlocktradeFutureMinOrderQty   string `json:"blocktrade_future_min_order_qty"`
	BlocktradeFutureMaxOrderQty   string `json:"blocktrade_future_max_order_qty"`
	BlocktradeFuturePriceStep     string `json:"blocktrade_future_price_step"`
	BlocktradeFutureSizeStep      string `json:"blocktrade_future_size_step"`

	BlocktradeOptionMinOrderPrice string `json:"blocktrade_option_min_order_price"`
	BlocktradeOptionMaxPutPrice   string `json:"blocktrade_option_max_put_price"`
	BlocktradeOptionMinOrderQty   string `json:"blocktrade_option_min_order_qty"`
	BlocktradeOptionMaxOrderQty   string `json:"blocktrade_option_max_order_qty"`
	// BlocktradeOptionPriceStep unit is QuoteCcy (blocktrade doesn't need priceStepBase)
	BlocktradeOptionPriceStep string `json:"blocktrade_option_price_step"`
	BlocktradeOptionSizeStep  string `json:"blocktrade_option_size_step"`

	CustomizePosLimit    bool   `json:"customize_pos_limit"`
	FuturePosLimitByPair string `json:"future_pos_limit_by_pair"`
	OptionPosLimitByPair string `json:"option_pos_limit_by_pair"`

	CustomTierLeverage bool `json:"custom_tier_leverage"`
	OptionBuyOnly      bool `json:"option_buy_only"` // whitelist info should not be leaked
}

func (v *UsdxUserPairConfigVoItem) Pair() string {
	return v.BaseCcy + "-" + v.QuoteCcy
}

type UsdxUserOpenOrderParamVoItem struct {
	Ccy                            string `json:"ccy"`
	CcyDisplayName                 string `json:"ccy_display_name"`
	IsPm                           bool   `json:"is_pm"`
	IsCust                         bool   `json:"is_cust"`
	MaxOpenCount                   int    `json:"max_open_count"`
	MaxOptionOpenCountByCcy        int    `json:"max_option_open_count_by_ccy"`
	MaxOptionOpenCountByInstrument int    `json:"max_option_open_count_by_instrument"`
	MaxFutureOpenCountByCcy        int    `json:"max_future_open_count_by_ccy"`
	MaxFutureOpenCountByInstrument int    `json:"max_future_open_count_by_instrument"`
	MaxOptionTotalUsdPosByCcy      string `json:"max_option_total_usd_pos_by_ccy"`
	MaxFutureTotalUsdPosByCcy      string `json:"max_future_total_usd_pos_by_ccy"`
	MaxTotalUsdPosByCcy            string `json:"max_total_usd_pos_by_ccy"`
	MaxStopOpenCount               int    `json:"max_stop_open_count"`
}

type UsdxFeeRatesWithSourceVo struct {
	Pair         string `json:"pair"`
	TakerFeeRate string `json:"taker_fee_rate"` // final taker
	MakerFeeRate string `json:"maker_fee_rate"` // final maker
	Source       string `json:"source"`
	// empty means value doesn't exist
	TakerBasic       string `json:"taker_basic"`
	MakerBasic       string `json:"maker_basic"`
	TakerUserDefined string `json:"taker_user_defined"`
	MakerUserDefined string `json:"maker_user_defined"`
	TakerVipLevel    string `json:"taker_vip_level"`
	MakerVipLevel    string `json:"maker_vip_level"`
	HasVipLevel      bool   `json:"has_vip_level"`
	VipLevel         int    `json:"vip_level"`
}

type UsdxUserFeerateClassVoItem struct {
	BaseCcy          string `json:"base_ccy"`
	QuoteCcy         string `json:"quote_ccy"`
	BaseDisplayName  string `json:"base_display_name"`
	QuoteDisplayName string `json:"quote_display_name"`

	PerpetualFeeRateInfo *UsdxFeeRatesWithSourceVo `json:"perpetual_fee_rate_info"`
	OptionFeeRateInfo    *UsdxFeeRatesWithSourceVo `json:"option_fee_rate_info"`
}

func (v *UsdxUserFeerateClassVoItem) Pair() string {
	return v.BaseCcy + "-" + v.QuoteCcy
}

type UsdxUserFeerateCustInfo struct {
	Ccy                       string `json:"ccy"`
	HasCustomizedFeeRates     bool   `json:"has_customized_fee_rates"`
	HasSameUsdxFutureFeeRates bool   `json:"has_same_usdx_future_fee_rates"`
	HasSameUsdxOptionFeeRates bool   `json:"has_same_usdx_option_fee_rates"`
}

type UsdxUserConfigVo struct {
	CcyOpenParams       []*UsdxUserOpenOrderParamVoItem `json:"ccy_open_params,omitempty"` // show in public API, hide in IMS
	Pairs               []*UsdxUserPairConfigVoItem     `json:"pairs"`
	FeerateClassList    []*UsdxUserFeerateClassVoItem   `json:"feerate_class_list"`
	FeerateCustInfoList []*UsdxUserFeerateCustInfo      `json:"feerate_cust_info_list"`

	FinalFeeType  string `json:"final_fee_type"`
	FinalVipLevel string `json:"final_vip_level"`
}
