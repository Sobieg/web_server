package main

import "errors"

var (
	invalidSessionError = errors.New("200 session id not valid")
	fieldNotFoundError  = errors.New("201 field with name ")
)
