package convert

type path struct {
	value string
}

type Client struct {
	convert converCmd
	source  path
	digest  path
	Error   error
}

func NewClient() (*Client, error) {
	con, err := getConvertCmd()
	if err != nil {
		return nil, err
	}
	client := &Client{convert: con}
	return client, nil
}

func (client *Client) Version() string {
	return client.convert.Version()
}
