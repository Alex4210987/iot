package http_param

import (
	"backend/model"
)

// Response 响应
type Response struct {
	Status  string     `json:"status"`
	Message string     `json:"message,omitempty"`
	Token   string     `json:"token,omitempty"`
	Data    model.Data `json:"data,omitempty"`
}
