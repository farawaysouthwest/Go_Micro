package controller

import (
	svc "go_mvc/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller build assets.

type Controller interface {

	// root route controller.
  GetRoot(c *gin.Context) 

  // user route controller.
  GetUser(c *gin.Context)
}

type controller struct {
	service  svc.Service
}

func NewController(service svc.Service) Controller {
	return controller{
		service: service,
	}
}

// Controller functions.

func (r controller) GetRoot(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "test",
		})
}

func (r controller) GetUser(c *gin.Context) {

	idString, ok := c.Params.Get("userid"); if ok {

	userId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		println(err)
	}
	
	user := r.service.FetchUser(userId)

	//return
	c.JSON(200, gin.H{
			"user": user,
		})
	}


}