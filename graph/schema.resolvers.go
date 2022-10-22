package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/walteraandrade/tnb-go-be/graph/generated"
	"github.com/walteraandrade/tnb-go-be/graph/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// CreateCommentary is the resolver for the createCommentary field.
func (r *mutationResolver) CreateCommentary(ctx context.Context, input model.NewCommentary) (*model.Comementary, error) {
	panic(fmt.Errorf("not implemented: CreateCommentary - createCommentary"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	dummyPost := model.Post{
		Title: "Dummy Post",
		Text:  "This is a dummy post",
	}

	posts = append()
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comementary, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
