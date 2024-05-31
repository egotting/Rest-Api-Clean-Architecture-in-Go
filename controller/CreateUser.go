package controller

import (
	"github.com/egotting/exceptions/validation"
	"github.com/egotting/model/DTO/User"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest DTO.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
}
