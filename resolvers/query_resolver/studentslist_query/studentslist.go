package studentslist_query

import (
	"context"
	"graphql-go-crud/data/connector/mongoconnector"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Product     string `json:"product"`
	Image       string `json:"image"`
	Company     string `json:"company"`
	Model       string `json:"model"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type StudentResponse struct {
	Message    string    `json:"message"`
	Statuscode int       `json:"statuscode"`
	Data       []Student `json:"data"`
}


func FetchstudentList(params graphql.ResolveParams) (interface{}, error) {

	client, _ := mongoconnector.MongoConnector()

	collection := client.Database("vamshi-test").Collection("products")

	filter := bson.M{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	var students []Student
	for cur.Next(context.TODO()) {
		var student Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	output := StudentResponse{Message: "Successfully fetched students", Statuscode: 200, Data: students}
	return output, nil
}
