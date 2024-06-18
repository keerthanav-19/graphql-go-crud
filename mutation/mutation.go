package mutation

import (
	"graphql-go-crud/resolvers/mutation_resolver/studentlist_mutation"
	"graphql-go-crud/schemas/studentsListSchema"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"updatestudent": &graphql.Field{
				Type:        studentsListSchema.UpdateStudentList,
				Description: "Update student",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"key": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: studentlist_mutation.UpdateStudent,
			},
			"deletestudent": &graphql.Field{
				Type:        studentsListSchema.UpdateStudentList,
				Description: "delete student",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: studentlist_mutation.DeleteStudent,
			},
			"insertstudent": &graphql.Field{
				Type:        studentsListSchema.UpdateStudentList,
				Description: "insert student",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"key": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: studentlist_mutation.InsertStudent,
			},
		},
	})
