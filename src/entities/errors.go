package entities

import (
	"github.com/pkg/errors"
)

type AppError struct {
    Code     int    `json:"code"`
	Message  string `json:"message"`
	Error	 string  `json:"error"`
}

var ErrSQL = errors.New("SQL_ERROR")
var ErrRepo = errors.New("REPO_ERROR")
var ErrCtrl = errors.New("CTRL_ERROR")