package common

type successResponse struct {
	Data   any `json:"data"`
	Paging any `json:"paging,omitempty"`
	Filter any `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter any) *successResponse {
	return &successResponse{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data any) *successResponse {
	return &successResponse{data, nil, nil}
}
