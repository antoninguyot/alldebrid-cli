package http

type Stream struct {
	Id       string      `json:"id"`
	Ext      string      `json:"ext"`
	Quality  interface{} `json:"quality"`
	Filesize float32     `json:"filesize"`
}

type LinkUnlockResponse BaseResponse[struct {
	Id         string   `json:"id"`
	Link       string   `json:"link"`
	Host       string   `json:"host"`
	HostDomain string   `json:"hostDomain"`
	Filename   string   `json:"filename"`
	Filesize   float32  `json:"filesize"`
	Streams    []Stream `json:"streams"`
}]

func (c Client) LinkUnlock(link string) (*LinkUnlockResponse, error) {
	response, err := c.R().
		SetResult(LinkUnlockResponse{}).
		SetQueryParam("link", link).
		Get("/link/unlock")
	if err != nil {
		return nil, err
	}
	return response.Result().(*LinkUnlockResponse), nil
}

type LinkStreamingResponse BaseResponse[struct {
	Filename string  `json:"filename"`
	Filesize float32 `json:"filesize"`
	Link     string  `json:"link,omitempty"`
	Delayed  int     `json:"delayed,omitempty"`
}]

func (c Client) LinkStreaming(id string, stream string) (*LinkStreamingResponse, error) {
	response, err := c.R().
		SetResult(LinkStreamingResponse{}).
		SetQueryParam("id", id).
		SetQueryParam("stream", stream).
		Get("/link/streaming")
	if err != nil {
		return nil, err
	}
	return response.Result().(*LinkStreamingResponse), nil
}

type DelayedStatus int

const (
	DelayedStatusProcessing DelayedStatus = 1
	DelayedStatusAvailable  DelayedStatus = 2
	DelayedStatusError      DelayedStatus = 3
)

type LinkDelayedResponse BaseResponse[struct {
	Status   DelayedStatus `json:"status"`
	TimeLeft int           `json:"time_left"`
	Link     string        `json:"link,omitempty"`
}]

func (c Client) LinkDelayed(id string) (*LinkDelayedResponse, error) {
	response, err := c.R().
		SetResult(LinkDelayedResponse{}).
		SetQueryParam("id", id).
		Get("/link/delayed")
	if err != nil {
		return nil, err
	}
	return response.Result().(*LinkDelayedResponse), nil
}
