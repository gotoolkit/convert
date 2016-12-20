package convert

import "fmt"

type path struct {
	value string
}

func (p *path) Ready() bool {
	return p.value != ""
}

func (p *path) Get() string {
	return p.value
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

func (c *Client) Src(srcPath string) *Client {
	c.source = path{srcPath}
	return c
}

func (c *Client) Digest(digestPath string) *Client {
	c.digest = path{digestPath}
	return c
}

func (c *Client) Out() (string, error) {
	err := c.ready()
	if err != nil {
		return "", err
	}

	return c.execute()

}
func (c *Client) execute() (string, error) {
	args := []string{
		c.source.Get(),
	}
	if c.digest.Ready() {
		args = append(args, c.digest.Get())
	}
	return c.convert.Execute(args)
}

func (c *Client) ready() error {
	if !c.source.Ready() {
		return fmt.Errorf("Source is not set")
	}
	return nil
}

func (c *Client) Version() string {
	return c.convert.Version()
}
