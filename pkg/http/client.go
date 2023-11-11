package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func NewClient(token string) *Client {

	return &Client{
		Client: resty.New().
			SetBaseURL("https://api.alldebrid.com/v4").
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetHeader("User-Agent", "alldebrid-cli").
			// // Doesn't seem to work despite being documented
			// SetAuthToken(token).
			// SetAuthScheme("Bearer").
			OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
				req.
					SetQueryParam("agent", "alldebrid-cli").
					// Fix for the above warning
					SetQueryParam("apikey", token)
				return nil
			}).
			OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
				var baseResponse BaseResponse[any]
				err := json.Unmarshal(resp.Body(), &baseResponse)
				if err != nil {
					return fmt.Errorf("error decoding error from alldebrid (HTTP %s). full JSON: %s", resp.Status(), resp.Body())
				}

				if baseResponse.Status != ResponseStatusSuccess {
					return fmt.Errorf("alldebrid error %s: %s", baseResponse.Error.Code, baseResponse.Error.Message)

				}

				return nil // if its success otherwise return error
			}),
	}
}

type Client struct {
	*resty.Client
}

type ResponseStatus string

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusError   ResponseStatus = "error"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse[T any] struct {
	Status ResponseStatus `json:"status"`
	Error  *Error         `json:"error"`
	Data   T              `json:"data"`
}
