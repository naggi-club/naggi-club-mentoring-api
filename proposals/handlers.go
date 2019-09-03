package proposals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/naggi-club/naggi-club-mentoring-api/common"
)

func ProposalList(c *gin.Context) {
	db := common.GetDB()

	var proposals []Proposal
	db.Order("created_at desc").Find(&proposals)

	c.JSON(http.StatusOK, gin.H{
		"proposals": proposals,
	})
}

func CreateProposal(c *gin.Context) {
	db := common.GetDB()
	var proposal Proposal

	if err := c.ShouldBindJSON(&proposal); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err.Error(),
		})
		return
	}

	if errs := db.Save(&proposal).GetErrors(); len(errs) > 0 {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Error())
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"errors": errMessages,
		})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Proposal created successfully",
		"resourceId": proposal.ID,
	})
}
