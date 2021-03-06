package main

import "errors"

var (
	invalidSessionError  = errors.New("200 session id not valid")
	fieldNotFoundError   = errors.New("201 field with name ")
	developerSecretError = errors.New("202 wrong secret for developer")
	credError            = errors.New("203 session id and creds from different accounts")
)
