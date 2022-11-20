package handler

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	model "github.com/saint-rivers/saint-ark/models"
	"github.com/saint-rivers/saint-ark/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readGetQueries(ctx *fiber.Ctx) (string, time.Time) {
	desiredDateFormat := "2006-01-02"

	format := ctx.Query("format")
	dateString := ctx.Query("date")

	date, err := time.Parse(desiredDateFormat, dateString)
	if err != nil {
		return format, time.Time{}
	}
	return strings.ToLower(format), date
}

func setMongoFilters(format string, date time.Time) primitive.D {
	// dateFilter := bson.D{}
	filterFields := bson.A{}

	var dateStart time.Time
	var dateEnd time.Time

	if !date.IsZero() {
		dateStart = date
		dateEnd = date.Add(time.Hour*time.Duration(23) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59))
	} else {
		year, month, day := time.Now().Date()
		dateStart = time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

		dateEnd = dateStart.Add(time.Hour*time.Duration(23) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59))
	}

	if format != "" {
		filterFields = append(filterFields, bson.M{"format": format})
	}

	// set date filter
	dateFilter := bson.D{
		{Key: "$gte", Value: dateStart},
		{Key: "$lt", Value: dateEnd},
	}
	filterFields = append(filterFields, bson.M{"uploaddate": dateFilter})

	filter := bson.D{{Key: "$and", Value: filterFields}}

	fmt.Println("filter", filter)
	return filter
}

// HealthCheck godoc
// @Summary Get all images from the server.
// @Tags image-handler
// @Accept */*
// @Param format query string false "specify file format"
// @Param date query string false "date"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/images [get]
func GetListedImages(ctx *fiber.Ctx, client *mongo.Client) error {
	imageCollection := getImageCollection(client)

	// use this function to get all queries from the context
	format, date := readGetQueries(ctx)

	// setup filters based on queries
	filter := setMongoFilters(format, date)

	// set mongo fetch options
	findOptions := options.Find()
	findOptions.SetLimit(10)

	// fetch data
	current, err := imageCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// read dataset results
	var results []*model.Resource
	for current.Next(context.TODO()) {
		var resource model.Resource
		err := current.Decode(&resource)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &resource)
	}
	if err := current.Err(); err != nil {
		log.Fatal(err)
	}
	current.Close(context.TODO())

	// respond to user request
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return ctx.JSON(fiber.Map{
		"message": "fetched data",
		"payload": results,
	})
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags image-handler
// @Accept 	multipart/form-data
// @Param myFile formData file false "single file upload"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/images [post]
func InsertImage(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := getImageCollection(client)
	createdUrl, fileName, originalFilename, err := utils.SaveFile(ctx)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	// get file format
	ext := filepath.Ext(fileName)[1:]

	// create object
	newImage := model.Resource{Id: fileName, Name: originalFilename, Format: strings.ToLower(ext), UploadDate: time.Now()}

	// insert
	insertResult, err := collection.InsertOne(context.TODO(), newImage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted: ", insertResult.InsertedID)

	return ctx.JSON(fiber.Map{
		"message": "successfully inserted data",
		"payload": map[string]string{
			"url":       createdUrl,
			"fileName":  fileName,
			"timestamp": time.Now().String(),
		},
	})
}

func getImageCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("saint-ark").Collection("resources")
}
