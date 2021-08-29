package component

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IControlService interface {
	GetControlInfo(ctx *gin.Context, id int)
}

type ControlService struct {
//	repo IControlRepository
}

func NewControlService() *ControlService {
	return &ControlService{}
}

func (c *ControlService) GetControlInfo(ctx *gin.Context, id int) {
	fmt.Println("cont to Service")
	fmt.Println("id : ", id)
}