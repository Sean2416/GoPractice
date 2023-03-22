package controller_model

import (
	"github.com/go-playground/validator/v10"
	"net/url"
)

//留著做 urlVerify 範例
type UpdateConsultantReq struct {
	Fname        *string `json:"fname" binding:"required" example:"fname"`
	Introduction *string `json:"introduction" binding:"required"  example:"自我介紹"`
	UrlPicture *string `json:"url_picture" binding:"urlVerify" example:"https://s3-oss.tutorabc.com/stage/gogoldtalk/test/02.png"`
}

func (u *UpdateConsultantReq) GetFname() string {
	if u.Fname == nil {
		return ""
	}
	return *u.Fname
}

func (u *UpdateConsultantReq) GetIntroduction() string {
	if u.Introduction == nil {
		return ""
	}
	return *u.Introduction
}

var UrlVerify validator.Func = func(fl validator.FieldLevel) (r bool) {
	s, ok := fl.Field().Interface().(string)
	if ok && s != "" {
		_, err := url.ParseRequestURI(s)
		if err != nil {
			return false
		}
	}
	return true
}


type CreateReferralUrlReq struct {
	Url string `json:"url" binding:"required" example:"url"`
}

type GetConsultantResp struct {
	Basiccname  string `json:"basicCname"`
	Nationality string `json:"nationality"`
	Conname     string `json:"conName"`
	Basiclname  string `json:"basicLname"`
	Location    string `json:"location"`
	Consn       int    `json:"conSn"`
	Basicfname  string `json:"basicFname"`
	Workfor     string `json:"workFor"`
	Status      string `json:"status"`
}
