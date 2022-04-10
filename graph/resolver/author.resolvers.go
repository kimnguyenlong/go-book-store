package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	filter := bson.M{"authorsId": bson.M{"$all": bson.A{obj.ID}}}
	cs, err := r.DB.Collection("books").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	defer cs.Close(context.Background())
	err = cs.All(context.Background(), &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

type authorResolver struct{ *Resolver }
