package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jajakim22/go-crud/initializers"
	"github.com/jajakim22/go-crud/models"
)


func PostsCreate(c *gin.Context){

	//GET DATA OFF REQ BODY
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

   //CREATE A POST 
   post := models.Post{Title: body.Title, Body: body.Body}

   result := initializers.DB.Create(&post)
  
   if result.Error != nil {
	c.Status(400)
	return
   }

   //RETURN DATA
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context) {
	//Get  ALL the posts
    var posts []models.Post
    initializers.DB.Find(&posts)

	//Respond with the posts
    c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	//Get The ID FROM URL
	id := c.Param("id")

	var post models.Post
	//Get One Post
	initializers.DB.First(&post, id )

		//Respond with the posts
		c.JSON(200, gin.H{
			"post": post,
		})
}


func PostsUpdate(c *gin.Context) {
	//GET ID FROM URL
    id := c.Param("id")

	//GET THE DATA OFF REQ BODY
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//FIND THE POST I AM UPDATING
	var post models.Post
	initializers.DB.First(&post, id )

	//UPDATE IT
	// Update attributes with `struct`, will only update non-zero fields
    initializers.DB.Model(&post).Updates(models.Post{
		 Title: body.Title,
		 Body: body.Body,
		})

	//RESPOND WITH IT
    c.JSON(200, gin.H{
		"post": post,
	})

}


func PostsDelete(c *gin.Context) {
	//GET ID FROM URL
	id := c.Param("id")

	//DELETE THE POST
	initializers.DB.Delete(&models.Post{}, id)

	//RESPOND
	c.Status(200)

 
}