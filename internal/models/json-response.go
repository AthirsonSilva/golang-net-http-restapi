package models

// Type responsible for mapping this dummy JSON response
type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
