package utils

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"Status"`
	Code       string `json:"code"`
}
