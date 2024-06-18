package studentlist_mutation

import (
	"context"
	"errors"
	"graphql-go-crud/data/connector/mongoconnector"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteStudentResponse struct {
	Message    string
	Statuscode int
}

func DeleteStudentEntry(id string) error {
	client, err := mongoconnector.MongoConnector()

	collection := client.Database("vamshi-test").Collection("products")

	filter := bson.M{"id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("No matching student found for deletion")
	}

	return nil
}

func DeleteStudent(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	log.Println(id)

	err := DeleteStudentEntry(id)
	if err != nil {
		log.Println("Error deleting student:", err)
		output := DeleteStudentResponse{Message: "error while deleting student", Statuscode: 200}
		return output, nil
	}

	output := DeleteStudentResponse{Message: "successfully deleted student", Statuscode: 400}
	return output, nil
}
