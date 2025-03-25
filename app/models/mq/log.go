package mq

type LogMessage struct {
	RequestID string `json:"request_id"`
	Method    string `json:"method"`
	URL       string `json:"url"`
	UserAgent string `json:"user_agent"`
	Body      string `json:"body"`
}
