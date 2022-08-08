package dto

import "strings"

type ErrResp struct {
	Err              string      `json:"error"`
	ErrorDescription string      `json:"error_description"`
	Errors           interface{} `json:"errors"`
}

func (e ErrResp) Error() string {
	switch vv := e.Errors.(type) {
	case string:
		return vv
	case []string:
		return strings.Join(vv, ",")
	}
	return "unknown error"
}
