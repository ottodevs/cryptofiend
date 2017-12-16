package binance

type ErrorInfo struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
}

type Balance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free,string"`
	Locked float64 `json:"locked,string"`
}

type AccountInfo struct {
	MakerCommission  int        `json:"makerCommission"`
	TakerCommission  int        `json:"takerCommission"`
	BuyerCommission  int        `json:"buyerCommission"`
	SellerCommission int        `json:"sellerCommission"`
	CanTrade         bool       `json:"canTrade"`
	CanWithdraw      bool       `json:"canWithdraw"`
	CanDeposit       bool       `json:"canDeposit"`
	Balances         []*Balance `json:"balances"`
}

type OrderType string

const (
	OrderTypeMarket OrderType = "MARKET"
	OrderTypeLimit  OrderType = "LIMIT"
)

type ExchangeInfo struct {
	Symbols []SymbolInfo
}

type SymbolStatus string

const (
	SymbolStatusTrading SymbolStatus = "TRADING"
)

type FilterType string

const (
	FilterTypePrice       FilterType = "PRICE_FILTER"
	FilterTypeLotSize     FilterType = "LOT_SIZE"
	FilterTypeMinNotional FilterType = "MIN_NOTIONAL"
)

type SymbolInfoFilter struct {
	Type FilterType `json:"filterType"`

	// PRICE_FILTER parameters
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`

	// LOT_SIZE parameters
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`

	// MIN_NOTIONAL parameters
	MinNotional string `json:"minNotional"`
}

type SymbolInfo struct {
	Symbol              string             `json:"symbol"`
	Status              SymbolStatus       `json:"status"`
	BaseAsset           string             `json:"baseAsset"`
	BaseAssetPrecision  int                `json:"baseAssetPrecision"`
	QuoteAsset          string             `json:"quoteAsset"`
	QuoteAssetPrecision int                `json:"quoteAssetPrecision"`
	OrderTypes          []OrderType        `json:"orderTypes"`
	Iceberg             bool               `json:"icebergAllowed"`
	Filters             []SymbolInfoFilter `json:"filters"`
}
