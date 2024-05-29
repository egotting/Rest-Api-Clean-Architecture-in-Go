package controller

import (
	"fmt"

	"github.com/egotting/exceptions"
	"github.com/egotting/model/DTO/User"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest DTO.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := exceptions.NewBadRequestError(
			fmt.Sprintf("there are incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
}
