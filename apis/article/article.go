package apis_article

import (
	"context"
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type articleAddBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author,omitempty"`
	Content     string `json:"content"`
	Status      int    `json:"status"`
}

func ArticleAdd(c *gin.Context) {
	var body articleAddBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	coll, err := lib.MongoCollection("article")
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	article := model.Article{
		Title:       body.Title,
		Description: body.Description,
		Author:      body.Author,
		Content:     body.Content,
		Status:      body.Status,
	}

	result, err := coll.InsertOne(context.TODO(), article)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": result.InsertedID,
	})
}

func ArticleGet(c *gin.Context) {
	coll, err := lib.MongoCollection("article")
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	var results []model.MArticle
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list": results,
	})

}
