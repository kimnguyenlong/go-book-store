package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *topicResolver) Books(ctx context.Context, obj *model.Topic) ([]*model.Book, error) {
	filter := bson.M{"topicsId": bson.M{"$all": bson.A{obj.ID}}}
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

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

type topicResolver struct{ *Resolver }
