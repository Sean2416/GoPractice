package controller_model

type UpdateMemberInfoReq struct {
	PhoneCode       string `json:"phone_code"`
	Phone           string `json:"phone"`
	Gender          int64  `json:"gender"`
	NickName        string `json:"nick_name"`
	ChnName         string `json:"chn_name"`
	Email           string `json:"email"`
	UrlPicture      string `json:"url_picture"`
	UrlBanner       string `json:"url_banner"`
	AboutMe         string `json:"about_me"`
	MoreInformation string `json:"more_information"`
	Professional    string `json:"professional"`
}

type SendSMSLimitReq struct {
	PhoneCode string `json:"phone_code" binding:"required" example:"886"`
	Phone     string `json:"phone" binding:"required" example:"0930049641"`
	ClientIP  string `json:"client_ip" example:"127.0.0.1"`
}

type GetSmSLimitResp struct {
	TotalLimit    int
	DailyLimit    int
	SlidingWindow map[string]bool
	SingleLimit   map[string]int
	KeyDate       string
	Err           string
}

type SendSMSLimitResp struct {
	IsLimiting  bool   `json:"is_limiting"`
	ShowMessage string `json:"show_message"`
	UserVal     int    `json:"user_val"`
	SysVal      int    `json:"sys_val"`
	K           string `json:"k"`
}
