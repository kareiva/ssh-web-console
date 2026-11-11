package models

import "github.com/genshen/ssh-web-console/src/utils"

const (
	SIGN_IN_FORM_TYPE_ERROR_VALID = iota
	SIGN_IN_FORM_TYPE_ERROR_PASSWORD
	SIGN_IN_FORM_TYPE_ERROR_TEST
	SIGN_IN_FORM_TYPE_ERROR_HOST_KEY_UNKNOWN
	SIGN_IN_FORM_TYPE_ERROR_HOST_KEY_CHANGED
)

type UserInfo struct {
	utils.JwtConnection
	Username string `json:"username"`
	Password string `json:"-"`
	HostKey  []byte `json:"-"` // host key accepted at sign-in, in ssh wire format
}

type JsonResponse struct {
	HasError bool        `json:"has_error"`
	Message  interface{} `json:"message"`
	Addition interface{} `json:"addition"`
}
