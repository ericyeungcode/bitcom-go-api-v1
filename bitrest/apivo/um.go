package apivo

// account_mode: classic/um
// risk_mode: regular/pm

type ClassicAccountVo struct {
	Currency        string `json:"currency"`
	ClassicRiskMode string `json:"classic_risk_mode"`
}

type AccountModeVo struct {
	UserId          int64              `json:"user_id"`
	AccountMode     string             `json:"account_mode"`
	AutoBorrow      bool               `json:"auto_borrow"`
	UmRiskMode      string             `json:"um_risk_mode"`
	ClassicAccounts []ClassicAccountVo `json:"classic_accounts"`
}

// //////////////
type UmAccountDetailVo struct {
	Currency               string `json:"currency"`
	CcyDisplayName         string `json:"ccy_display_name"`
	Equity                 string `json:"equity"`
	Liability              string `json:"liability"`
	IndexPrice             string `json:"index_price"`
	CashBalance            string `json:"cash_balance"`
	MarginBalance          string `json:"margin_balance"`
	AvailableBalance       string `json:"available_balance"`
	InitialMargin          string `json:"initial_margin"`
	SpotMargin             string `json:"spot_margin"`
	MaintenanceMargin      string `json:"maintenance_margin"`
	PotentialLiability     string `json:"potential_liability"`
	Interest               string `json:"interest"`
	InterestRate           string `json:"interest_rate"`
	Pnl                    string `json:"pnl"`
	TotalDelta             string `json:"total_delta"`
	SessionRpl             string `json:"session_rpl"`
	SessionUpl             string `json:"session_upl"`
	OptionPositionValue    string `json:"option_position_value"` // ignored for deri
	OptionValue            string `json:"option_value"`
	OptionPnl              string `json:"option_pnl"`
	OptionSessionRpl       string `json:"option_session_rpl"`
	OptionSessionUpl       string `json:"option_session_upl"`
	OptionDelta            string `json:"option_delta"`
	OptionGamma            string `json:"option_gamma"`
	OptionVega             string `json:"option_vega"`
	OptionTheta            string `json:"option_theta"`
	FuturePositionValue    string `json:"future_position_value"` // ignored for deri
	FutureValue            string `json:"future_value"`
	FuturePnl              string `json:"future_pnl"`
	FutureSessionRpl       string `json:"future_session_rpl"`
	FutureSessionUpl       string `json:"future_session_upl"`
	FutureSessionFunding   string `json:"future_session_funding"`
	FutureDelta            string `json:"future_delta"`
	FutureAvailableBalance string `json:"future_available_balance"`
	OptionAvailableBalance string `json:"option_available_balance"`
	UnsettledAmount        string `json:"unsettled_amount"`
	LiqPrice               string `json:"liq_price"`
}

type UmAccountVo struct {
	UserId                      int64                `json:"user_id"`
	CreatedAt                   int64                `json:"created_at"`
	TotalCollateral             string               `json:"total_collateral"`
	TotalMarginBalance          string               `json:"total_margin_balance"`
	TotalAvailable              string               `json:"total_available"`
	TotalInitialMargin          string               `json:"total_initial_margin"`
	TotalMaintenanceMargin      string               `json:"total_maintenance_margin"`
	TotalInitialMarginRatio     string               `json:"total_initial_margin_ratio"`
	TotalMaintenanceMarginRatio string               `json:"total_maintenance_margin_ratio"`
	TotalLiability              string               `json:"total_liability"`
	TotalUnsettledAmount        string               `json:"total_unsettled_amount"`
	TotalFutureValue            string               `json:"total_future_value"`
	TotalOptionValue            string               `json:"total_option_value"`
	SpotOrdersHcLoss            string               `json:"spot_orders_hc_loss"`
	TotalPositionPnl            string               `json:"total_position_pnl"`
	TotalPositionValue          string               `json:"total_position_value"`
	Leverage                    string               `json:"leverage"`
	Details                     []*UmAccountDetailVo `json:"details"`
}

type UmTxLogVo struct {
	TxTime       int64  `json:"tx_time"`
	TxType       string `json:"tx_type"`
	Ccy          string `json:"ccy"`
	InstrumentId string `json:"instrument_id"`
	Direction    string `json:"direction"` // support format : `open-buy`
	Qty          string `json:"qty"`
	Price        string `json:"price"`
	CashFlow     string `json:"cash_flow"`
	Position     string `json:"position"`
	FeePaid      string `json:"fee_paid"`
	FeeCcy       string `json:"fee_ccy"`
	FeeRate      string `json:"fee_rate"`
	Funding      string `json:"funding"`
	Change       string `json:"change"`
	Balance      string `json:"balance"`
	OrderID      string `json:"order_id"`
	TradeID      string `json:"trade_id"`
	Remark       string `json:"remark"`
	DisplayName  string `json:"display_name"`
}

type UmAccountCcyConfigVo struct {
	Currency                        string `json:"currency"`
	MaxLeverage                     string `json:"max_leverage"`
	MaxLiabilityOfAutoBorrowMode    string `json:"max_liability_of_auto_borrow_mode"`
	MaxLiabilityOfNonAutoBorrowMode string `json:"max_liability_of_non_auto_borrow_mode"`
	InterestFreeLiability           string `json:"interest_free_liability"`
	CcyDisplayName                  string `json:"ccy_display_name"`
}

type UmAccountConfigVo struct {
	CurrencyConfigs []UmAccountCcyConfigVo `json:"currency_configs"`
}

type InterestVo struct {
	Currency       string `json:"currency"`
	Time           int64  `json:"time"`
	LoanRate       string `json:"loan_rate"`
	Liability      string `json:"liability"`
	Interest       string `json:"interest"`
	DeductInterest bool   `json:"deduct_interest"`
	CcyDisplayName string `json:"ccy_display_name"`
}
