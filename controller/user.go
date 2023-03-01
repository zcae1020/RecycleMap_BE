package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func userRouter(user *gin.RouterGroup) {
	user.GET("/comments", getAllMyLocation)
	user.PATCH("/comments/:commentId", updateComment)
	user.DELETE("/comments/:commentId", deleteComment)
}

func getAllMyLocation(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error get all my location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	commentDto, err := service.FindCommentsByUID(token)

	if err != nil {
		log.Printf("[controller:user] error get all my location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, commentDto)
}

func updateComment(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error update my comment : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func deleteComment(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error delete my comment : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}
