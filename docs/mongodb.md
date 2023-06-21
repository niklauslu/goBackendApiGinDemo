### 数据库使用(mongodb)

#### MongoDB的安装
本地开发可以使用[MongoDB Community Edition](https://www.mongodb.com/docs/manual/administration/install-community/#std-label-install-community)  
GUI客户端[Studio 3T](https://studio3t.com/)

#### golang使用MongoDB
[官方Driver](https://www.mongodb.com/docs/drivers/go/current/)

```bash
go get go.mongodb.org/mongo-driver/mongo
```

配置
```conf
file: .env
// 连接配置
MONGODB_URI=mongodb://test_admin:123456@localhost:27017/?maxPoolSize=20&w=majority
MONGODB_DATABASE=test
```

使用Bson映射模型(以collection:article为例)
```go
// file: model/article.go
type Article struct {
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Author      string             `bson:"author,omitempty"`
	Content     string             `bson:"content"`
	Status      int                `bson:"stat"`
}

// 获取数据用到，映射json
type MArticle struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Author      string             `bson:"author,omitempty" json:"author,omitempty"`
	Content     string             `bson:"content" json:"content" `
	Status      int                `bson:"status" json:"status" `
}
```

添加article示例
```go
type articleAddBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author,omitempty"`
	Content     string `json:"content"`
	Status      int    `json:"stat"`
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
```

查找artcles示例
```go
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
```