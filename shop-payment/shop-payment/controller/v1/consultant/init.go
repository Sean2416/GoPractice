package consultant

import "github.com/gin-gonic/gin"

type ConsultantController interface {
	CreateConsultant(*gin.Context)
	GetConsultant(*gin.Context)
	CreateReferralUrl(ctx *gin.Context)
}

type consultantController struct {
}

func NewConsultantController() ConsultantController {
	return &consultantController{}
}