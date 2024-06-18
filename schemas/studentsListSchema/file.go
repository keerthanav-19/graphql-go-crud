package studentsListSchema

import "github.com/graphql-go/graphql"

var FileList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FileList",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.String,
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

// var FileListInputType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "FileListInputType",
// 	Fields: graphql.Fields{
// 		"Input": &graphql.Field{Type: graphql.String},
// 	},
// })
