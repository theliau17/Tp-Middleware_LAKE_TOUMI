package models

import "fmt"

type CustomError struct {
	Message string `default:""`
	Code    int    `default:"200"`
}

type APIError struct {
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}
