package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/inabatron/go_crud/inits"
	"github.com/inabatron/go_crud/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := inits.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get posts
	var posts []models.Post
	inits.DB.Find(&posts)
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")
	//get posts
	var post models.Post
	inits.DB.First(&post, id)
	//respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post to be updated
	var post models.Post
	inits.DB.First(&post, id)

	//update it
	inits.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//get the id off url
	id := c.Param("id")

	//delete post
	inits.DB.Delete(&models.Post{}, id)
	//response
	c.Status(200)
}
