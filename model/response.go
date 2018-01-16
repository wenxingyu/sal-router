package model

// message send back to user
type UserResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
