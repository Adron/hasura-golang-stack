package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

func main() {
	graphqlClient := graphql.NewClient("https://graphql-go-workflow.hasura.app/v1/graphql")
	graphqlRequest := graphql.NewRequest(`
		{
		  tododeluge_Todo(where: {TodoId: {_is_null: true}}) {
			Id
			Details
			Todo
			Todos {
			  Id
			  Details
			  Todo
			}
		  }
		}
	`)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		fmt.Println(err)
	}

	fmt.Println(graphqlResponse)
}
