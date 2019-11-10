package view

import (
	"github.com/gin-gonic/gin"
	"log"
)

//login req
type GetLoginStructReq struct {
	name     string `json:"name" form:"name"`
	password string `json:"password",form:"password"`
}
type GetLoginStructRsp struct {
	Rsp struct {
		Status string `json:"status"`
		Data   string `json:"data"`
		Desc   string `json:"desc"`
	}
}

func Login(c *gin.Context) {
	rsp := GetLoginStructRsp{}.Rsp
	req := &GetLoginStructReq{}
	rsp.Desc = "error"
	if err := c.Bind(req); err != nil {
		c.JSON(200, rsp)
		log.Println("未知的req")
	} else {
		rsp.Status = "ok"
		rsp.Desc = ""
		c.JSON(200, rsp)
	}
}
func Submit(c *gin.Context) {

}
