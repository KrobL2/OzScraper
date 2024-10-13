package telegram

import (
	"encoding/json"
	"net/url"
	e "scrappy/pkg/errors"
	"strconv"
)

func (c *Client) SendMessageWithKeyboard(chatID int, text string, keyboard *InlineKeyboardMarkup) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	if keyboard != nil {
		keyboardBytes, err := json.Marshal(keyboard)
		if err != nil {
			return e.Wrap("can't marshal keyboard", err)
		}

		q.Add("reply_markup", string(keyboardBytes))
	}

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}
