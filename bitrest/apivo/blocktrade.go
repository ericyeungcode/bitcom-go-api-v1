package apivo

type BlockTradeResponse struct {
	Label string `json:"label"`
	// BlockOrderStatus
	Status string `json:"status"`
}

type BlockTradeQueryVo struct {
	BlockOrderId string `json:"block_order_id" example:"1001"`
	Label        string `json:"label" example:"abc"`
	CreatedAt    int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt    int64  `json:"updated_at" example:"1585296000000"`
	UserId       int64  `json:"user_id" example:"801"`
	Counterparty int64  `json:"counterparty" example:"1002"`
	InstrumentId string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Side         string `json:"side" example:"buy"`
	Price        string `json:"price" example:"0.03"`
	Qty          string `json:"qty" example:"1"`
	Fee          string `json:"fee" example:"0.0001"`
	Status       string `json:"status" example:"filled"`
	Role         string `json:"role" example:"taker"`
	BtSource     string `json:"bt_source"`
	OrderId      string `json:"order_id"`

	// additional info for paradigm
	TradeId    string `json:"trade_id" example:"3826"`
	IndexPrice string `json:"index_price" example:"6800"`
	Sigma      string `json:"sigma" example:"0.0024"`
}
