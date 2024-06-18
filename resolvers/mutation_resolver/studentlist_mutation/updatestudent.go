package studentlist_mutation

import (
	"context"
	"errors"
	"graphql-go-crud/data/connector/mongoconnector"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateStudentResponse struct {
	Message    string
	Statuscode int
}

func UpdateStudentEntry(id, key string) error {
	client, err := mongoconnector.MongoConnector()

	collection := client.Database("vamshi-test").Collection("products")

	filter := bson.M{"id": id}

	update := bson.M{"$set": bson.M{"company": key}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("No matching student found for update")
	}

	return nil
}

func UpdateStudent(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	key, _ := params.Args["key"].(string)
	log.Println(id, key)

	err := UpdateStudentEntry(id, key) 
	if err != nil {
		log.Println("Error updating student:", err)
		output := UpdateStudentResponse{Message: "error while updating students", Statuscode: 200}
		return output, nil
	}

	output := UpdateStudentResponse{Message: "sucessfully updated students", Statuscode: 400}
	return output, nil
}
