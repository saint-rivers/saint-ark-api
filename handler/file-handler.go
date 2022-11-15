package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	model "github.com/saint-rivers/saint-ark/models"
	"github.com/saint-rivers/saint-ark/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetListedImages(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := getResourceCollection(client)

	findOptions := options.Find()
	findOptions.SetLimit(10)

	curr, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var results []*model.Resource

	for curr.Next(context.TODO()) {
		var resource model.Resource
		err := curr.Decode(&resource)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &resource)
	}

	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}
	curr.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return ctx.JSON(fiber.Map{
		"message": "fetched data",
		"payload": results,
	})
}

func InsertImage(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := getResourceCollection(client)
	_, fileName, originalFilename, err := utils.SaveFile(ctx)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	newImage := model.Resource{Id: fileName, Name: originalFilename, Format: "JGP", UploadDate: time.Now()}

	insertResult, err := collection.InsertOne(context.TODO(), newImage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted: ", insertResult.InsertedID)

	return ctx.JSON(fiber.Map{
		"message": "successfully inserted data",
		"payload": fileName,
	})
}

func getResourceCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("saint-ark").Collection("resources")
}
