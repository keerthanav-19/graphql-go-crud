package studentsListSchema

import "github.com/graphql-go/graphql"

var StudentsList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StudentsList",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.NewList(StudentListInputType),
			},
			"statuscode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var StudentListInputType = graphql.NewObject(graphql.ObjectConfig{
	Name: "StudentListInputType",
	Fields: graphql.Fields{
		"ID":          &graphql.Field{Type: graphql.String},
		"Key":         &graphql.Field{Type: graphql.String},
		"Product":     &graphql.Field{Type: graphql.String},
		"Image":       &graphql.Field{Type: graphql.String},
		"Company":     &graphql.Field{Type: graphql.String},
		"Model":       &graphql.Field{Type: graphql.String},
		"Price":       &graphql.Field{Type: graphql.String},
		"Category":    &graphql.Field{Type: graphql.String},
		"Description": &graphql.Field{Type: graphql.String},
	},
})

var UpdateStudentList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UpdateStudentList",
		Fields: graphql.Fields{
			"statuscode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
