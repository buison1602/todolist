package helper

import "fmt"

var LoginError = fmt.Errorf("invalid user name or wrong password")

var EmailError = fmt.Errorf("invalid email")

var DataError = fmt.Errorf("data can not empty")
