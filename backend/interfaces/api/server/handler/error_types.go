package handler

import "fmt"

var (
	OK = fmt.Errorf("OK")
	ERR_BAD_REQUEST = fmt.Errorf("BAD_REQUEST")
	ERR_NOT_FOUND = fmt.Errorf("NOT_FOUND")
	ERR_INTERNAL = fmt.Errorf("INTERNAL ERROR")
)