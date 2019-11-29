package query

import (
  "../model"
	"github.com/graphql-go/graphql"
)

func SetUserQuery() graphql.Fields {
  // graphql.FieldはGraphQLのスキーマ定義の項目? https://github.com/graphql-go/graphql/blob/2e2b648ecbe42b217c0d7c2bf45fa607bf3f1201/definition.go#L604-L611
  // graphql.ResolveParams https://github.com/graphql-go/graphql/blob/199d20bbfed70dae8c7d4619d4e0d339ce738b43/definition.go#L571-L585
    query := graphql.Fields{
        "User": &graphql.Field{
            Type:    graphql.NewList(UserType()), // 配列で返却するときはgraphql.NewListを使う
            Args:    UserArgs(),
            Resolve: UserResolve,
        },
        "Product": &graphql.Field{
          Type: graphql.String,
          Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            return "ppppppp", nil
          },
        },
        "Address": &graphql.Field{
          Type: graphql.String,
          Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            return "aaaaaaaa", nil
          },
		    },
    }
    return query
}

// UserType 返り値として返せる値の定義
func UserType() *graphql.Object {
    return graphql.NewObject(graphql.ObjectConfig{
        Name: "User",
        Fields: graphql.Fields{
            "ID": &graphql.Field{
                Type: graphql.Int,
            },
            "Name": &graphql.Field{
                Type: graphql.String,
            },
            "Email": &graphql.Field{
                Type: graphql.String,
            },
        },
    })
}

func UserArgs() map[string]*graphql.ArgumentConfig {
    return graphql.FieldConfigArgument{
        "ID": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.Int),
        },
        "Name": &graphql.ArgumentConfig{
            Type: graphql.String,
        },
    }
}

func UserResolve(params graphql.ResolveParams) (interface{}, error) {
    var users []model.User
    for _, u := range model.UserAll() {
      if u.ID != params.Args["ID"] {
        continue
      }
      users = append(users, *u)
    }
    return users, nil
}
