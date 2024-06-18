package studentlist_mutation

import (
	"context"
	"graphql-go-crud/data/connector/mongoconnector"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type InsertStudentResponse struct {
	Message    string
	Statuscode int
}

func InsertStudentEntry(id string, key string) error {
	client, err := mongoconnector.MongoConnector()

	collection := client.Database("vamshi-test").Collection("products")

	student := bson.M{
		"id":  id,
		"key": key,
	}

	_, err = collection.InsertOne(context.Background(), student)
	if err != nil {
		return err
	}

	return nil
}

func InsertStudent(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	key, _ := params.Args["key"].(string)
	log.Println(id, key)

	err := InsertStudentEntry(id, key)
	if err != nil {
		log.Println("Error inserting student:", err)
		output := InsertStudentResponse{Message: "error while inserting student", Statuscode: 200}
		return output, nil
	}

	output := InsertStudentResponse{Message: "successfully inserted student", Statuscode: 400}
	return output, nil
}
