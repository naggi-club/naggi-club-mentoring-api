package proposals

import (
	"github.com/gin-gonic/gin"
)

func ProposalsRegister(router *gin.RouterGroup) {
	router.GET("/proposals/", ProposalList)
	router.POST("/proposals/", CreateProposal)
}
