package apivo

type UsdxInstrumentVo struct {
	InstrumentId         string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	CreatedAt            int64  `json:"created_at" example:"1582790400000"`
	UpdatedAt            int64  `json:"updated_at" example:"1582790400000"`
	BaseCurrency         string `json:"base_currency" example:"BTC"`
	QuoteCurrency        string `json:"quote_currency" example:"USD"`
	StrikePrice          string `json:"strike_price" example:"9000"`
	ExpirationAt         int64  `json:"expiration_at" example:"1585296000000"`
	OptionType           string `json:"option_type" example:"call"`
	Category             string `json:"category"`
	MinPrice             string `json:"min_price" example:"0.0005"`
	MaxPrice             string `json:"max_price" example:"10"`
	PriceStep            string `json:"price_step" example:"0.0005"`
	MinSize              string `json:"min_size" example:"0.1"`
	SizeStep             string `json:"size_step" example:"0.1"`
	DeliveryFeeRate      string `json:"delivery_fee_rate" example:"0.0002"`
	ContractSize         string `json:"contract_size"`
	ContractSizeCurrency string `json:"contract_size_currency"`
	Active               bool   `json:"active"`
	Status               string `json:"status"`

	Groups     []int64  `json:"groups"`
	GroupSteps []string `json:"group_steps"`

	DisplayAt       int64  `json:"display_at" example:"1585296000000"`
	IsDisplay       bool   `json:"is_display"`
	IsDisplayIv     bool   `json:"is_display_iv"`
	OptionPrecision string `json:"option_precision"`

	DisplayName      string `json:"display_name"`
	BaseDisplayName  string `json:"base_display_name"`
	QuoteDisplayName string `json:"quote_display_name"`

	ImgUrl string `json:"img_url"`
}

type OrderBookVo struct {
	InstrumentId string      `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Timestamp    int64       `json:"timestamp" example:"1585299600000"`
	Bids         [][2]string `json:"bids"`
	Asks         [][2]string `json:"asks"`
	DisplayName  string      `json:"display_name"`
}

type UsdxPairVo struct {
	Pair  string `json:"pair" example:"BTC-USD"`
	Base  string `json:"base" example:"BTC"`
	Quote string `json:"quote" example:"USD"`

	PairDisplayName  string `json:"pair_display_name"`
	BaseDisplayName  string `json:"base_display_name"`
	QuoteDisplayName string `json:"quote_display_name"`
}

type TickerVo struct {
	Time         int64  `json:"time"`
	InstrumentId string `json:"instrument_id"`

	BestBid    string `json:"best_bid"`
	BestAsk    string `json:"best_ask"`
	BestBidQty string `json:"best_bid_qty"`
	BestAskQty string `json:"best_ask_qty"`
	AskSigma   string `json:"ask_sigma"`
	BidSigma   string `json:"bid_sigma"`

	LastPrice      string `json:"last_price"`
	LastQty        string `json:"last_qty"`
	Open24H        string `json:"open24h"`
	High24H        string `json:"high24h"`
	Low24H         string `json:"low24h"`
	PriceChange24H string `json:"price_change24h"`
	Volume24H      string `json:"volume24h"`
	VolumeUSD24H   string `json:"volume_usd24h,omitempty"`
	OpenInterest   string `json:"open_interest"`
	FundingRate    string `json:"funding_rate,omitempty"`
	FundingRate8H  string `json:"funding_rate8h,omitempty"`
	PremiumRate    string `json:"premium_rate,omitempty"`

	UnderlyingName  string `json:"underlying_name,omitempty"`
	UnderlyingPrice string `json:"underlying_price,omitempty"`
	MarkPrice       string `json:"mark_price"`
	IndexPrice      string `json:"index_price"`
	Sigma           string `json:"sigma,omitempty"`
	Delta           string `json:"delta,omitempty"`
	Vega            string `json:"vega,omitempty"`
	Theta           string `json:"theta,omitempty"`
	Gamma           string `json:"gamma,omitempty"`

	MinSellPrice string `json:"min_sell"`
	MaxBuyPrice  string `json:"max_buy"`

	DisplayName string `json:"display_name"`

	MarkHigh24H         string `json:"mark_high24h,omitempty"`
	MarkLow24H          string `json:"mark_low24h,omitempty"`
	MarkPriceChange24H  string `json:"mark_price_change24h,omitempty"`
	IndexHigh24H        string `json:"index_high24h,omitempty"`
	IndexLow24H         string `json:"index_low24h,omitempty"`
	IndexPriceChange24H string `json:"index_price_change24h,omitempty"`
}

type MarketTradeVo struct {
	CreatedAt       int64  `json:"created_at" example:"1585299600000"`
	TradeId         int64  `json:"trade_id" example:"3743"`
	InstrumentId    string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Price           string `json:"price" example:"0.034"`
	Qty             string `json:"qty" example:"1"`
	Side            string `json:"side" example:"buy"`
	Sigma           string `json:"sigma" example:"0.002"`
	IndexPrice      string `json:"index_price" example:"8000"`
	UnderlyingPrice string `json:"underlying_price" example:"7900"`
	MarkPrice       string `json:"mark_price"`
	IsBlockTrade    bool   `json:"is_block_trade" example:"false"`
	DisplayName     string `json:"display_name"`
}

type MarketSummaryVo struct {
	InstrumentId string `json:"instrument_id"`
	Timestamp    int64  `json:"timestamp"`

	BestBid    string `json:"best_bid"`
	BestAsk    string `json:"best_ask"`
	BestBidQty string `json:"best_bid_qty"`
	BestAskQty string `json:"best_ask_qty"`

	LastPrice    string `json:"last_price"`
	LastQty      string `json:"last_qty"`
	Open24H      string `json:"open24h"`
	High24H      string `json:"high24h"`
	Low24H       string `json:"low24h"`
	Volume24H    string `json:"volume24h"`
	OpenInterest string `json:"open_interest"`

	MarkPrice string `json:"mark_price"`

	MaxBuy  string `json:"max_buy"`
	MinSell string `json:"min_sell"`

	Delta string `json:"delta"`
	Gamma string `json:"gamma"`
	Vega  string `json:"vega"`
	Theta string `json:"theta"`
	//Sigma string `json:"sigma"`

	DisplayName string `json:"display_name"`
}
