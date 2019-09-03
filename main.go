package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/naggi-club/naggi-club-mentoring-api/common"
	"github.com/naggi-club/naggi-club-mentoring-api/proposals"
)

func Migrate(db *gorm.DB) {
	proposals.AutoMigrate()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")

	proposals.ProposalsRegister(api)
	return router
}

func main() {
	db := common.Init()
	Migrate(db)

	router := SetupRouter()
	router.Run(":3000")
}
