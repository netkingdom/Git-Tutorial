package component

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ControlHandler struct {
	controlService IControlService
}



func NewControlHandler(service *ControlService) *ControlHandler {
	return &ControlHandler{
		controlService: service,
	}
}

func (c *ControlHandler) GetControlInfo(ctx *gin.Context)  {
	id, _ := strconv.Atoi( ctx.Param("id"))
	fmt.Println("come2")


	fmt.Println(ctx)
	c.controlService.GetControlInfo(ctx, id)
}