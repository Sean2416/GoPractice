package consultant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"shop-payment/shop-payment/app"
	"shop-payment/shop-payment/controller_model"
	"net/http"
	"scm.tutorabc.com/tgo-framework/go-log"
	"strconv"
)

// @Tags Consultant
// @Summary 創建老師專屬URL
// @version 1.0
// @Accept  json
// @Produce  json
// @Param conSn path int true "顧問Sn" default(119691)
// @Param sendBody body controller_model.CreateReferralUrlReq true "body"  default("http://www.google.com")
// @Success 200 {string} string data
// @Router /v1/referral/{conSn} [POST]
func (u *consultantController) CreateReferralUrl(c *gin.Context) {
	////// 1. controller層不做太多邏輯, 做參數效驗
	////// 2. ctx 上下游鏈路調用, 必取往下傳
	////// 3. 展示如何使用 ShouldBindBodyWith 驗證參數, 詳情見gin 官方 github
	//////


	ctx := c.Request.Context()
	log.Info(ctx)
	appG := app.Gin{C: c}

	var reqBody controller_model.CreateReferralUrlReq
	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		appG.ResponseCasual(http.StatusOK, app.ParamBad, nil)
		return
	}
	conSn, err := strconv.ParseInt(c.Param("conSn"), 10, 64)
	if err != nil {
		appG.ResponseErrCasual(http.StatusOK, app.ParamBad, err)
		return
	}

	appG.ResponseCasual(http.StatusOK, app.SUCCESS, fmt.Sprintf("https://www.tutorabc.com?conSn=%d", conSn))
}
