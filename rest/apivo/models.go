package apivo

type OrderActionVo struct {
	OrderId        string `json:"order_id" example:"1001"`
	CreatedAt      int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt      int64  `json:"updated_at" example:"1585296000000"`
	UserId         string `json:"user_id" example:"801"`
	InstrumentId   string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	OrderType      string `json:"order_type" example:"limit"`
	Side           string `json:"side" example:"buy"`
	Price          string `json:"price" example:"0.03"`
	Qty            string `json:"qty" example:"1"`
	TimeInForce    string `json:"time_in_force" example:"gtc"`
	AvgPrice       string `json:"avg_price" example:"0.029"`
	FilledQty      string `json:"filled_qty" example:"1"`
	Status         string `json:"status" example:"filled"`
	IsLiquidation  bool   `json:"is_liquidation" example:"false"`
	AutoPrice      string `json:"auto_price" example:"7000"`
	AutoPriceType  string `json:"auto_price_type" example:"usd"`
	TakerFeeRate   string `json:"taker_fee_rate" example:"0.00005"`
	MakerFeeRate   string `json:"maker_fee_rate" example:"0.00002"`
	Label          string `json:"label" example:"strategy-A"`
	StopPrice      string `json:"stop_price"`
	ReduceOnly     bool   `json:"reduce_only"`
	PostOnly       bool   `json:"post_only"`
	RejectPostOnly bool   `json:"reject_post_only"`
	Mmp            bool   `json:"mmp"`
	Source         string `json:"source"`
	Hidden         bool   `json:"hidden"`

	FeeDeductionEnabled bool   `json:"fee_deduction_enabled"`
	FeeDeductionCcy     string `json:"fee_deduction_ccy"`
	FeeDeductionRate    string `json:"fee_deduction_rate"`
}

type OrderVo struct {
	OrderId       string `json:"order_id" example:"1001"`
	CreatedAt     int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt     int64  `json:"updated_at" example:"1585296000000"`
	UserId        string `json:"user_id" example:"801"`
	InstrumentId  string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	OrderType     string `json:"order_type" example:"limit"`
	Side          string `json:"side" example:"buy"`
	Price         string `json:"price" example:"0.03"`
	Qty           string `json:"qty" example:"1"`
	TimeInForce   string `json:"time_in_force" example:"gtc"`
	AvgPrice      string `json:"avg_price" example:"0.029"`
	FilledQty     string `json:"filled_qty" example:"1"`
	Status        string `json:"status" example:"filled"`
	IsLiquidation bool   `json:"is_liquidation" example:"false"`
	AutoPrice     string `json:"auto_price" example:"7000"`
	AutoPriceType string `json:"auto_price_type" example:"usd"`
	PNL           string `json:"pnl" example:"0.031"`
	CashFlow      string `json:"cash_flow" example:"0.027"`

	InitialMargin   string `json:"initial_margin" example:"0.04"`
	TakerFeeRate    string `json:"taker_fee_rate" example:"0.00005"`
	MakerFeeRate    string `json:"maker_fee_rate" example:"0.00002"`
	CancelReason    string `json:"cancel_reason" example:"ioc cancelled"`
	Label           string `json:"label" example:"strategy-A"`
	StopOrderId     string `json:"stop_order_id" example:"stop-x3gjsdhf3232"`
	StopPrice       string `json:"stop_price" example:"9800"`
	ReduceOnly      bool   `json:"reduce_only"`
	PostOnly        bool   `json:"post_only"`
	RejectPostOnly  bool   `json:"reject_post_only"`
	Mmp             bool   `json:"mmp"`
	ReorderIndex    int64  `json:"reorder_index"`
	Source          string `json:"source"`
	Hidden          bool   `json:"hidden"`
	IsUm            bool   `json:"is_um"`
	TpPrice         string `json:"tp_price"`
	SlPrice         string `json:"sl_price"`
	DisplayName     string `json:"display_name"`
	TpslTriggerType string `json:"tpsl_trigger_type"`
	TwmktId         int64  `json:"twmkt_id"`
	UserStgyId      uint64 `json:"user_stgy_id"`

	Fee    string `json:"fee" example:"0.00002"`
	FeeCcy string `json:"fee_ccy"`

	FeeDeductionEnabled bool   `json:"fee_deduction_enabled"`
	FeeInDeductionCcy   string `json:"fee_in_deduction_ccy" example:"0.00002"`
	FeeDeductionCcy     string `json:"fee_deduction_ccy"`
	FeeDeductionRate    string `json:"fee_deduction_rate"`

	FeeCcyDisplayName          string `json:"fee_ccy_display_name"`
	FeeDeductionCcyDisplayName string `json:"fee_deduction_ccy_display_name"`
}

type BatchOrderDetailVo struct {
	OrderId       string `json:"order_id" example:"1001"`
	CreatedAt     int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt     int64  `json:"updated_at" example:"1585296000000"`
	UserId        string `json:"user_id" example:"801"`
	InstrumentId  string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	OrderType     string `json:"order_type" example:"limit"`
	Side          string `json:"side" example:"buy"`
	Price         string `json:"price" example:"0.03"`
	Qty           string `json:"qty" example:"1"`
	TimeInForce   string `json:"time_in_force" example:"gtc"`
	AvgPrice      string `json:"avg_price" example:"0.029"`
	FilledQty     string `json:"filled_qty" example:"1"`
	Status        string `json:"status" example:"filled"`
	IsLiquidation bool   `json:"is_liquidation" example:"false"`
	AutoPrice     string `json:"auto_price" example:"7000"`
	AutoPriceType string `json:"auto_price_type" example:"usd"`
	TakerFeeRate  string `json:"taker_fee_rate" example:"0.00005"`
	MakerFeeRate  string `json:"maker_fee_rate" example:"0.00002"`
	Label         string `json:"label" example:"strategy-A"`
	//StopPrice     string `json:"stop_price"`
	ReduceOnly     bool   `json:"reduce_only"`
	PostOnly       bool   `json:"post_only"`
	RejectPostOnly bool   `json:"reject_post_only"`
	Mmp            bool   `json:"mmp"`
	Source         string `json:"source"`
	Hidden         bool   `json:"hidden"`
	DisplayName    string `json:"display_name"`
	ErrorCode      int    `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`

	FeeDeductionEnabled bool   `json:"fee_deduction_enabled"`
	FeeDeductionCcy     string `json:"fee_deduction_ccy"`
	FeeDeductionRate    string `json:"fee_deduction_rate"`
}

type UsdxBatchVo struct {
	Orders []BatchOrderDetailVo `json:"orders"`
}

type OrderCancelVo struct {
	NumCancelled int64 `json:"num_cancelled"`
}

type TradeVo struct {
	TradeId         string `json:"trade_id" example:"3826"`
	OrderId         string `json:"order_id" example:"1001"`
	InstrumentId    string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Qty             string `json:"qty" example:"1"`
	Price           string `json:"price" example:"0.2275"`
	Sigma           string `json:"sigma" example:"0.0024"`
	UnderlyingPrice string `json:"underlying_price" example:"6750"`
	IndexPrice      string `json:"index_price" example:"6800"`
	UsdPrice        string `json:"usd_price" example:"1664"`
	Fee             string `json:"fee" example:"0.003"`
	FeeRate         string `json:"fee_rate" example:"0.0003"`
	Side            string `json:"side" example:"buy"`
	CreatedAt       int64  `json:"created_at" example:"1585296000000"`
	IsTaker         bool   `json:"is_taker" example:"true"`
	OrderType       string `json:"order_type" example:"limit"`
	IsBlockTrade    bool   `json:"is_block_trade" example:"false"`
	Label           string `json:"label" example:"hedge"`

	IsFeeDeducted        bool   `json:"is_fee_deducted"`
	FeeDeductionCcy      string `json:"fee_deduction_ccy"`
	FeeDeductionRate     string `json:"fee_deduction_rate"`
	FeeDeductionCcyIndex string `json:"fee_deduction_ccy_index"`
}

type UsdxPositionVo struct {
	UserId            int64  `json:"user_id"`
	InstrumentId      string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	ExpirationAt      int64  `json:"expiration_at" example:"1585296000000"`
	Qty               string `json:"qty" example:"5"`
	InitialMargin     string `json:"initial_margin" example:"0.023"`
	MaintenanceMargin string `json:"maintenance_margin" example:"0.0075"`
	AvgPrice          string `json:"avg_price" example:"0.016"`
	SessionAvgPrice   string `json:"session_avg_price" example:"0.012"`
	MarkPrice         string `json:"mark_price" example:"0.0401"`
	IndexPrice        string `json:"index_price" example:"7800"`
	LastPrice         string `json:"last_price"`
	QtyBase           string `json:"qty_base,omitempty" example:"0.001"`

	LiqPrice           string `json:"liq_price,omitempty" example:"6950"`
	SessionFunding     string `json:"session_funding,omitempty" example:"0.000023"`
	PositionPnl        string `json:"position_pnl" example:"0.0021"`
	PositionSessionUpl string `json:"position_session_upl" example:"0.0024"`
	PositionSessionRpl string `json:"position_session_rpl" example:"-0.0002"`

	// NEW
	Category string `json:"category" example:"option"`
	ROI      string `json:"roi" example:"0.0003"`

	OptionDelta         string `json:"option_delta,omitempty"`
	OptionGamma         string `json:"option_gamma,omitempty"`
	OptionVega          string `json:"option_vega,omitempty"`
	OptionTheta         string `json:"option_theta,omitempty"`
	OptionValue         string `json:"option_value,omitempty"`
	OptionPositionValue string `json:"option_position_value"`

	FutureValue         string `json:"future_value,omitempty"`
	FuturePositionValue string `json:"future_position_value"`

	Leverage    string `json:"leverage,omitempty" example:"30"`
	DisplayName string `json:"display_name"`

	AdlLevel string `json:"adl_level,omitempty"`
}

type UsdxCondOrderVo struct {
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Status       string `json:"status"`
	StopPrice    string `json:"stop_price"`
	TriggerType  string `json:"trigger_type"`
	RejectReason string `json:"reject_reason"`

	// fields of order
	CondOrderId  string `json:"cond_order_id"`
	InstrumentId string `json:"instrument_id"`
	UserId       int64  `json:"user_id"`
	Qty          string `json:"qty"`
	Price        string `json:"price"`
	Side         string `json:"side"`
	OrderType    string `json:"order_type"`
	TimeInForce  string `json:"time_in_force"`

	Source      string `json:"source"`
	Hidden      bool   `json:"hidden"`
	TpslMode    int    `json:"tpsl_mode"`
	DisplayName string `json:"display_name"`
	ReduceOnly  bool   `json:"reduce_only"`
	IsUserTwap  bool   `json:"is_user_twap"`
	UserStgyId  uint64 `json:"user_stgy_id"`
}

type LeverageRatioVo struct {
	Pair            string `json:"pair"`
	LeverageRatio   string `json:"leverage_ratio"`
	PairDisplayName string `json:"pair_display_name"`
}
