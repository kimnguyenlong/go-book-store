package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *bookResolver) Topics(ctx context.Context, obj *model.Book) ([]*model.Topic, error) {
	if len(obj.TopicsID) == 0 {
		return nil, nil
	}
	var topicsId []primitive.ObjectID
	for _, id := range obj.TopicsID {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		topicsId = append(topicsId, objId)
	}
	filter := bson.M{"_id": bson.M{"$in": topicsId}}
	cs, err := r.DB.Collection("topics").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var topics []*model.Topic
	defer cs.Close(context.Background())
	err = cs.All(context.Background(), &topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
	if len(obj.AuthorsID) == 0 {
		return nil, nil
	}
	var authorsId []primitive.ObjectID
	for _, id := range obj.AuthorsID {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		authorsId = append(authorsId, objId)
	}
	filter := bson.M{"_id": bson.M{"$in": authorsId}}
	cs, err := r.DB.Collection("authors").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var authors []*model.Author
	defer cs.Close(context.Background())
	err = cs.All(context.Background(), &authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *bookResolver) Reviews(ctx context.Context, obj *model.Book) ([]*model.Review, error) {
	filter := bson.M{"bookId": obj.ID}
	cs, err := r.DB.Collection("reviews").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var reviews []*model.Review
	defer cs.Close(context.Background())
	err = cs.All(context.Background(), &reviews)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
