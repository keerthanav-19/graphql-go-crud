package queries

import (
	"graphql-go-crud/resolvers/query_resolver/studentslist_query"
	"graphql-go-crud/schemas/studentsListSchema"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			"getstudentList": &graphql.Field{
				Type:        studentsListSchema.StudentsList,
				Description: "Get students list from MongoDB",
				Args: graphql.FieldConfigArgument{
					"infra": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: studentslist_query.FetchstudentList,
			},
			"fetchpythonfile": &graphql.Field{
				Type: studentsListSchema.FileList,
				//Description: "Get students list from MongoDB",
				// Args: graphql.FieldConfigArgument{
				// 	"infra": &graphql.ArgumentConfig{
				// 		Type: graphql.NewNonNull(graphql.String),
				// 	},
				// },
				Resolve: studentslist_query.Fetchpythonfile,
			},
		},
	})
