package db

import (
	"context"
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

func (dbHandler DBHandler) FindAllProducts() ([]*domain.Product, error) {
	var results []*domain.Product
	collection := dbHandler.database.Collection(productCollection)
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem domain.Product
		err2 := cur.Decode(&elem)
		if err2 != nil {
			log.Fatal(err2)
		}
		results = append(results, &elem)
	}
	return results, nil
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
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) DeleteProduct(id int64) error {
	collection := dbHandler.database.Collection(productCollection)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
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
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) FilterProductsByCategory(category string) ([]*domain.Product, error) {
	var results []*domain.Product
	collection := dbHandler.database.Collection(productCollection)
	cur, err := collection.Find(context.TODO(), bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem domain.Product
		err2 := cur.Decode(&elem)
		if err2 != nil {
			log.Fatal(err2)
		}
		results = append(results, &elem)
	}
	return results, nil
}
