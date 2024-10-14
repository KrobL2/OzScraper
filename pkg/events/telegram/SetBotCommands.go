package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) SetBotCommands(commands []BotCommand) error {
	q := url.Values{}
	q.Add("commands", mustJsonMarshal(commands))

	_, err := c.doRequest("setMyCommands", q)
	if err != nil {
		return fmt.Errorf("failed to set bot commands: %w", err)
	}

	return nil
}

func mustJsonMarshal(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(data)
}
