package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(ctx *gin.Context, param interface{}) {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	ctx.JSON(http.StatusOK, response)
}

func NewErrorResponse(ctx *gin.Context, status int, err error) {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Error occured"
	response.Meta.Messages = []string{err.Error()}

	ctx.JSON(status, response)
}