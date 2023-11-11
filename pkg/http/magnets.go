package http

type Link struct {
	Link    string `json:"link"`
	Filenme string `json:"filename"`
	Size    int    `json:"size"`
}

type Magnet struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	Status   string `json:"status"`
	Links    []Link `json:"links"`
}

type ListMagnetsResponse BaseResponse[struct {
	Magnets []Magnet `json:"magnets"`
}]

func (c Client) ListMagnets() (*ListMagnetsResponse, error) {
	response, err := c.R().SetResult(ListMagnetsResponse{}).Get("/magnet/status")
	if err != nil {
		return nil, err
	}
	return response.Result().(*ListMagnetsResponse), nil
}

type ShowMagnetResponse BaseResponse[struct {
	Magnets Magnet `json:"magnets"`
}]

func (c Client) ShowMagnet(id string) (*ShowMagnetResponse, error) {
	response, err := c.R().
		SetResult(ShowMagnetResponse{}).
		SetQueryParam("id", id).
		Get("/magnet/status")
	if err != nil {
		return nil, err
	}

	return response.Result().(*ShowMagnetResponse), nil
}

type UploadedMagnet struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Magnet string `json:"magnet"`
	Ready  bool   `json:"ready"`
	Hash   string `json:"hash"`
	Size   int    `json:"size"`
	Error  *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type UploadMagnetResponse BaseResponse[struct {
	Magnets []UploadedMagnet `json:"magnets"`
}]

func (c Client) UploadMagnet(magnet string) (*UploadMagnetResponse, error) {
	response, err := c.R().
		SetResult(UploadMagnetResponse{}).
		SetQueryParam("magnets[]", magnet).
		Get("/magnet/upload")
	if err != nil {
		return nil, err
	}

	return response.Result().(*UploadMagnetResponse), nil
}

type UploadFileResponse BaseResponse[struct {
	Files []UploadedMagnet `json:"files"`
}]

func (c Client) UploadFile(path string) (*UploadFileResponse, error) {
	response, err := c.R().
		SetResult(UploadFileResponse{}).
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("files[]", path).
		Post("/magnet/upload/file")
	if err != nil {
		return nil, err
	}

	return response.Result().(*UploadFileResponse), nil
}

type DeleteMagnetResponse BaseResponse[struct {
	Message string `json:"message"`
}]

func (c Client) DeleteMagnet(id string) (*DeleteMagnetResponse, error) {
	response, err := c.R().
		SetResult(DeleteMagnetResponse{}).
		SetQueryParam("id", id).
		Get("/magnet/delete")
	if err != nil {
		return nil, err
	}

	return response.Result().(*DeleteMagnetResponse), nil
}

type RestartMagnetResponse BaseResponse[struct {
	Message string `json:"message"`
}]

func (c Client) RestartMagnet(id string) (*RestartMagnetResponse, error) {
	response, err := c.R().
		SetResult(RestartMagnetResponse{}).
		SetQueryParam("id", id).
		Get("/magnet/restart")
	if err != nil {
		return nil, err
	}

	return response.Result().(*RestartMagnetResponse), nil
}
