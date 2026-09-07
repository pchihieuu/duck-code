package response

import "github.com/gin-gonic/gin"

// Envelope is the standard JSON shape every endpoint returns.
type Envelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// ErrorBody carries a machine-readable code plus a human message.
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// OK writes a 2xx success envelope.
func OK(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Envelope{Success: true, Data: data})
}

// OKWithMeta writes a 2xx success envelope with pagination/meta info.
func OKWithMeta(c *gin.Context, status int, data interface{}, meta interface{}) {
	c.JSON(status, Envelope{Success: true, Data: data, Meta: meta})
}

// Err writes an error envelope with the given HTTP status, code, and message.
func Err(c *gin.Context, status int, code, message string) {
	c.JSON(status, Envelope{Success: false, Error: &ErrorBody{Code: code, Message: message}})
}
