package api

import "github.com/lishimeng/iot-link/internal/db/repo"

type SimpleResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Item    interface{} `json:"item,omitempty"`
	Page    *repo.Page  `json:"page,omitempty"`
}

func NewBean() *SimpleResult {
	return &SimpleResult{Code: 0, Message: ""}
}
