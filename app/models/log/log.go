package model_log

type MiddlewareLogRequest struct {
	RequestID   string `json:"request_id"`
	Method      string `json:"method"`
	URL         string `json:"url"`
	UserAgent   string `json:"user_agent"`
	ContentType string `json:"content_type"`
	Time        string `json:"time"`
}
