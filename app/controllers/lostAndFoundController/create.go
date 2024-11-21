package lostAndFoundController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createLostAndFoundData struct {
	Type         bool   `json:"type"`                            // 1-失物 0-寻物
	Name         string `json:"name" binding:"required"`         // 物品名称
	Introduction string `json:"introduction" binding:"required"` // 物品介绍
	Campus       uint8  `json:"campus" binding:"required"`       // 校区 1-朝晖 2-屏峰 3-莫干山
	Kind         uint8  `json:"kind"`                            // 物品种类 1其他2证件3箱包4首饰5现金6电子产品7钥匙
	Place        string `json:"place" binding:"required"`        // 丢失或拾得地点
	Time         string `json:"time" binding:"required"`         // 丢失或拾得时间
	Imgs         string `json:"imgs" binding:"required"`         // 物品图片，多个图片以逗号分隔
	Contact      string `json:"contact" binding:"required"`      // 联系方式
}

// CreateLostAndFound 创建一条失物招领
func CreateLostAndFound(c *gin.Context) {
	var data createLostAndFoundData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = lostAndFoundService.SaveLostAndFound(models.LostAndFoundRecord{
		Type:         data.Type,
		Name:         data.Name,
		Introduction: data.Introduction,
		Campus:       data.Campus,
		Kind:         data.Kind,
		Place:        data.Place,
		Time:         data.Time,
		Imgs:         data.Imgs,
		Publisher:    utils.GetUser(c).StudentID,
		Contact:      data.Contact,
		IsProcessed:  2,
		IsApproved:   2,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}