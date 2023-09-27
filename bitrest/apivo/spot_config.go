package apivo

type SpotFeeRatesWithSourceVo struct {
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

type SpotPairConfigVo struct {
	Pair         string                    `json:"pair"`
	TakerFeeRate string                    `json:"taker_fee_rate"`
	MakerFeeRate string                    `json:"maker_fee_rate"`
	DisplayName  string                    `json:"display_name"`
	FeeRateClass *SpotFeeRatesWithSourceVo `json:"fee_rate_class"`
}

type SpotAccountConfigVo struct {
	UserID string `json:"user_id"`

	CustomizeOpenCounts   bool                `json:"customize_open_counts"`
	MaxOpenOrderCountAll  int64               `json:"max_open_order_count_all"`
	MaxOpenOrderCountPair int64               `json:"max_open_order_count_pair"`
	CustomizeFeeRates     bool                `json:"customize_fee_rates"`
	PairConfigList        []*SpotPairConfigVo `json:"pair_config_list"`
	ParentUserId          string              `json:"parent_user_id"`
	HasSameFeeRates       bool                `json:"has_same_fee_rates"`
}
