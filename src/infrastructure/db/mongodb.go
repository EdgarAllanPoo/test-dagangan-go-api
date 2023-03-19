package db

import (
	"context"
	"fmt"
	"log"

	"github.com/EdgarAllanPoo/test-go-api/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection string = "products"

type DBHandler struct {
	MongoClient mongo.Client
	database    *mongo.Database
}

func NewDBHandler(connectString string, dbName string) (DBHandler, error) {
	dbHandler := DBHandler{}
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbName)
	return dbHandler, nil
}

func (dbHandler DBHandler) FindAllProducts(category string, limit, offset int) ([]*domain.Product, int64, error) {
	var results []*domain.Product
	collection := dbHandler.database.Collection(productCollection)
	filter := bson.D{}
	if category != "" {
		filter = bson.D{{Key: "category", Value: category}}
	}
	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	var totalRows int64

	totalRows, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return results, 0, err
	}

	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		return results, 0, err
	}
	for cur.Next(context.Background()) {
		var elem domain.Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	return results, totalRows, nil
}

func (dbHandler DBHandler) FindProductById(id int64) (*domain.Product, error) {
	var result domain.Product
	collection := dbHandler.database.Collection(productCollection)
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (dbHandler DBHandler) SaveProduct(product domain.Product) error {
	collection := dbHandler.database.Collection(productCollection)

	filter := bson.M{"id": product.Id}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("product with Id %d already exists", product.Id)
	}

	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) DeleteProduct(id int64) error {
	collection := dbHandler.database.Collection(productCollection)
	res, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("product with Id %d not found", id)
	}

	return nil
}

func (dbHandler DBHandler) UpdateProduct(id int64, product domain.Product) error {
	collection := dbHandler.database.Collection(productCollection)
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{
		"name":     product.Name,
		"price":    product.Price,
		"category": product.Category,
	}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("product with Id %d not found", id)
	}

	return nil
}
