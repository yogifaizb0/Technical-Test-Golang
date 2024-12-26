package dto

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}