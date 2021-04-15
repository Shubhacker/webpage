package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/shubhacker/gqlgen-todos/graph/controller"
	"github.com/shubhacker/gqlgen-todos/graph/generated"
	"github.com/shubhacker/gqlgen-todos/graph/model"
)

func (r *queryResolver) MasterExcelFetch(ctx context.Context) (*model.MasterExcelResponce, error) {
	responce := controller.MaterExcelFetch(ctx)
	return responce, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
