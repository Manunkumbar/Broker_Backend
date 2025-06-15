package models

type Holding struct {
	UserID       int     `db:"user_id" json:"user_id"`
	StockSymbol  string  `db:"stock_symbol" json:"stock_symbol"`
	Qty          float64 `db:"quantity" json:"quantity"`
	AveragePrice float64 `db:"average_price" json:"average_price"`
}

type Order struct {
	UserID      int     `db:"user_id" json:"user_id"`
	StockSymbol string  `db:"stock_symbol" json:"stock_symbol"`
	OrderType   string  `db:"order_type" json:"order_type"`
	Quantity    float64 `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
	Status      string  `db:"status" json:"status"`
}

type Position struct {
	UserID      int     `db:"user_id" json:"user_id"`
	StockSymbol string  `db:"stock_symbol" json:"stock_symbol"`
	Quantity    float64 `db:"quantity" json:"quantity"`
	Pnl         float64 `db:"pnl" json:"pnl"`
}
