package http

type PinResponse BaseResponse[struct {
	Pin       string `json:"pin"`
	Check     string `json:"check"`
	ExpiresIn int    `json:"expires_in"`
	UserUrl   string `json:"user_url"`
	BaseUrl   string `json:"base_url"`
	CheckUrl  string `json:"check_url"`
}]

func (c Client) Pin() (*PinResponse, error) {
	response, err := c.R().
		SetResult(PinResponse{}).
		Get("/pin/get")
	if err != nil {
		return nil, err
	}
	return response.Result().(*PinResponse), nil
}

type PinCheckResponse BaseResponse[struct {
	Activated bool   `json:"activated"`
	ExpiresIn int    `json:"expires_in"`
	Apikey    string `json:"apikey"`
}]

func (c Client) PinCheck(check string, pin string) (*PinCheckResponse, error) {
	response, err := c.R().
		SetResult(PinCheckResponse{}).
		SetQueryParam("check", check).
		SetQueryParam("pin", pin).
		Get("/pin/check")

	if err != nil {
		return nil, err
	}
	return response.Result().(*PinCheckResponse), nil
}
