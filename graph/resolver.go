package graph

import "gitlab.com/pragmaticreviews/graphql-server/graph/model"

//go:generate go run github.com/99designs/gqlgen

// go:generateコマンドでgenerateフォルダ配下とschema.resolvers.goが再生成してくれる
type Resolver struct{
	videos []model.Video
}
