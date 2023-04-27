package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"result": getField(),
		})
	})

	r.Run()
}

func getField() string {
	client, err := firestore.NewClient(context.Background(), "rhamilton-001")
	if err != nil {
		log.Fatal(err)
	}

	docRef := client.Collection("test_collection").Doc("test_doc")
	doc, err := docRef.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return doc.Data()["test_field01"].(string)
}
