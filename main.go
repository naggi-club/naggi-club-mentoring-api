package main

import (
  "net/http"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
  proposalModel struct {
    gorm.Model
    Name                 string `json:"name" binding:"required"`
    Email                string `json:"email" binding:"required`
    PrograminExperience  string `json:"programming_experience" binding:"required`
    Content              string `json:"content" binding:"required`
  }
)

var db *gorm.DB

func init() {
  var err error
  db, err = gorm.Open("mysql", "root:@/naggi_club_mentoring?charset=utf8&parseTime=True&loc=Local")

  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&proposalModel{})
}

func SetupRouter() *gin.Engine {
  // Set the router as the default one shipped with Gin
  router := gin.Default()

  // Serve frontend static files
  router.Use(static.Serve("/", static.LocalFile("./views", true)))

  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }

  api.GET("/proposals", ProposalHandler)
  api.POST("/proposals", CreateProposal)

  return router
}

func main() {
  router := SetupRouter()
  router.Run(":3000")
}

func ProposalHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "Proposal Handler not implemented yet",
  })
}

func CreateProposal(c *gin.Context) {
  proposal := proposalModel{
    Name: c.PostForm("name"),
    Email: c.PostForm("email"),
    PrograminExperience: c.PostForm("programming_experience"),
    Content: c.PostForm("content"),
  }

  db.Save(&proposal)

  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusCreated, gin.H {
    "status": http.StatusCreated,
    "message":"Proposal created successfully",
    "resourceId": proposal.ID,
  })
}
