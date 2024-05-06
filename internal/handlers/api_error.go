package handlers

import "time"

type APIError struct {
	ErrorCode    int
    ErrorMessage string
    CreatedAt    time.Time
}