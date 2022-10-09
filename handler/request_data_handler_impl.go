package handler

import (
	"basic-webhook-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestDataHandlerImpl struct {
	requestDataService service.RequestDataService
}

func NewRequestDataHandler(requestDataService service.RequestDataService) *RequestDataHandlerImpl {
	return &RequestDataHandlerImpl{
		requestDataService: requestDataService,
	}
}

func (handler *RequestDataHandlerImpl) Post(c *gin.Context) {
	request_data_response, err := handler.requestDataService.Create(c, "POST")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": request_data_response,
	})
}