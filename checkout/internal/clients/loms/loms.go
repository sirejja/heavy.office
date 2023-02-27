package loms

const (
	StocksPath      = "/stocks"
	CreateOrderPath = "/createOrder"
)

type Client struct {
	url            string
	urlStocks      string
	urlCreateOrder string
}

func New(url string) *Client {
	return &Client{
		url:            url,
		urlStocks:      url + StocksPath,
		urlCreateOrder: url + CreateOrderPath,
	}
}
