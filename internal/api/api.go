package api
//denotes HTTP reply structure

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Payload struct {
	Meta any `json:"meta,omitempty"` //metadata of Data
	Data any `json:"data,omitempty"` //Actual Data
}

//Response denotes the structure of the JSON response template
type Response struct {
	Payload   *Payload  `json:"payload,omitempty"`
	Messages  []string `json:"messages,omitempty"` //Error Msg
	ErrorCode int      `json:"errorCode,omitempty"` 
}

//Use this function to return data without needing metadata
func SuccessMsg(c *gin.Context, data any, msgs ...string) {
	c.JSON(http.StatusOK, 
		Response {Messages: msgs,
					Payload: &Payload {
						Data: data, 
		}, 
	})
}

func FailMsg(c *gin.Context, status, code int, msgs ...string) {
	c.JSON(status, Response {
		Messages: msgs, 
		ErrorCode: code,
	})
}
