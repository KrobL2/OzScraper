package telegram

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"read-adviser-bot/lib/e"
)

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() {
		err = e.WrapIfErr("can't do request", err)
	}()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
