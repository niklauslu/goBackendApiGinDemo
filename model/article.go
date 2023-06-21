package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Author      string `bson:"author,omitempty"`
	Content     string `bson:"content"`
	Status      int    `bson:"status"`
}

type MArticle struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Author      string             `bson:"author,omitempty" json:"author,omitempty"`
	Content     string             `bson:"content" json:"content" `
	Status      int                `bson:"status" json:"status" `
}
