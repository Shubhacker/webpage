package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/shubhacker/gqlgen-todos/graph/controller"
	"github.com/shubhacker/gqlgen-todos/graph/generated"
	"github.com/shubhacker/gqlgen-todos/graph/model"
)

func (r *mutationResolver) UpsertToolData(ctx context.Context, input model.UpsertTool) (*model.UpsertToolResponce, error) {
	response := controller.UpsertToolData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpsertBookData(ctx context.Context, input model.UpsertBook) (*model.UpsertBookResponce, error) {
	response := controller.UpsertbookData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpdateToolData(ctx context.Context, input model.UpdateTools) (*model.UpsertToolResponce, error) {
	response := controller.UpdateToolsData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpdateBookData(ctx context.Context, input model.UpdateBook) (*model.UpsertBookResponce, error) {
	response := controller.UpdatebookData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpsertVideoData(ctx context.Context, input model.UpsertVideo) (*model.UpsertVideoResponce, error) {
	response := controller.UpsertVideoData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpdateVideoData(ctx context.Context, input model.UpdateVideo) (*model.UpsertVideoResponce, error) {
	response := controller.UpdateVideoData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpsertUserData(ctx context.Context, input model.UserUpsert) (*model.UserResponce, error) {
	response := controller.UpsertUserData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpdateUserData(ctx context.Context, input model.UpdateUser) (*model.UserResponce, error) {
	response := controller.UpdateUserData(ctx, input)
	return response, nil
}

func (r *mutationResolver) UpsertBlogData(ctx context.Context, input model.UpserBlogData) (*model.BlogResponce, error) {
	response := controller.UpsertBlogData(ctx, input)
	return response, nil
}

func (r *queryResolver) FetchTool(ctx context.Context, input *model.FetchToolsInput) (*model.ToolResponceData, error) {
	response := controller.FetchToolData(ctx, input)
	return response, nil
}

func (r *queryResolver) FetchData(ctx context.Context) (*model.Fetch, error) {
	response := controller.FetchTableData(ctx)
	return response, nil
}

func (r *queryResolver) FetchBook(ctx context.Context, input *model.FetchBookInput) (*model.BookResponce, error) {
	response := controller.FetchBookData(ctx, input)
	return response, nil
}

func (r *queryResolver) FetchVideo(ctx context.Context, input *model.FetchVideoInput) (*model.FetchVideoResponce, error) {
	response := controller.FetchVideoData(ctx, input)
	return response, nil
}

func (r *queryResolver) FetchBlog(ctx context.Context, input *model.FetchBlogInput) (*model.ResponceFetchBlog, error) {
	response := controller.FetchBlogData(ctx, input)
	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }