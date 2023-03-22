package consultant

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/protobuf/types/known/timestamppb"
	"shop-payment/shop-payment/app"
	"shop-payment/shop-payment/controller_model"
	"shop-payment/shop-payment/service"
	"shop-payment/shop-payment/sutil"
	"net/http"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/api"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/constants"
	"scm.tutorabc.com/tgo-framework/go-log"
	"strconv"
	"time"
)

// @Tags Consultant
// @Summary 新增顧問
// @version 1.0
// @Accept  json
// @Produce  json
// @Param Token header string false "token" default(1234)
// @Param sendBody body api.AddUserRequest true  "body"
// @Success 200 {object} api.UserResponseData data
// @Router /v1/consultant [POST]
func (u *consultantController) CreateConsultant(c *gin.Context) {
	////// 1. controller層不做太多邏輯, 做參數效驗
	////// 2. ctx 上下游鏈路調用, 必取往下傳
	////// 3. 展示如何使用 ShouldBindBodyWith 驗證參數, 詳情見gin 官方 github
	//////

	ctx := c.Request.Context()
	appG := app.Gin{C: c}
	var reqBody api.AddUserRequest
	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		appG.ResponseErr(http.StatusOK, constants.ErrNo_Param, err)
		return
	}

	data := api.UserResponse{
		ApiMeta: app.GenResponseApiMeta(ctx, constants.ErrNo_OK),
		Data: &api.UserResponseData{
			LastUpdated: timestamppb.New(time.Now()),
		},
	}
	appG.Response(http.StatusOK, data)
}

// @Tags Consultant
// @Summary 取得顧問資料
// @version 1.0
// @Accept  json
// @Produce  json
// @Param conSN path string true "顧問Sn" default(2913)
// @Success 200 {object} controller_model.GetConsultantResp data
// @Router /v1/consultant/{conSn} [GET]
func (u *consultantController) GetConsultant(c *gin.Context) {
	////// 1. controller層不做太多邏輯, 做參數效驗
	////// 2. ctx 上下游鏈路調用, 必取往下傳
	////// 3. 展示了log 使用方式
	////// 4. 這個function使用 慵懶的 response, 不建使用

	ctx := c.Request.Context()
	lb := sutil.NewLogBuild(ctx, "GetConsultant")
	defer func() {
		if lb.IsError() {
			log.Error(lb)
			return
		}
		log.Info(lb)
	}()

	appG := app.Gin{C: c}
	var err error
	////驗證參數
	conSn, err := strconv.ParseInt(c.Param("conSn"), 10, 64)
	if err != nil {
		lb.SetError(err)
		appG.ResponseErr(http.StatusOK, constants.ErrNo_Param, err)
		return
	}
	lb.Set("conSN", conSn)
	conInfo, err := service.GetConsultant(ctx, conSn)
	if err != nil {
		err = errors.New("get member error" + err.Error())
		appG.ResponseErr(http.StatusOK, constants.ErrNo_Unknown, err)
		return
	}
	lb.Set("conInfos", conInfo)
	resp := controller_model.GetConsultantResp{
		Basiccname: conInfo.Basiccname,
		Conname:    conInfo.Conname,
		Location:   conInfo.Location,
	}
	appG.ResponseCasual(http.StatusOK, app.SUCCESS, resp)
}
