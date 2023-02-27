package products

const (
	GetProductPath = "/get_product"
)

type Client struct {
	url            string
	token          string
	urlGetProducts string
}

func New(url, token string) *Client {
	return &Client{
		url:            url,
		token:          token,
		urlGetProducts: url + GetProductPath,
	}
}
