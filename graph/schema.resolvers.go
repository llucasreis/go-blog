package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/llucasreis/go-blog/graph/generated"
	"github.com/llucasreis/go-blog/graph/model"
	"github.com/llucasreis/go-blog/internal/posts"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var post posts.Post

	post.Content = input.Content
	post.Title = input.Title

	postID := post.Save()

	return &model.Post{
		ID:      strconv.FormatInt(postID, 10),
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var resultPosts []*model.Post
	dbPosts := posts.GetAll()

	for _, post := range dbPosts {
		resultPosts = append(resultPosts, &model.Post{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return resultPosts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
