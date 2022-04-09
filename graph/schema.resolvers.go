package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	filter := bson.M{"authorsId": bson.M{"$all": bson.A{obj.ID}}}
	cs, err := r.DB.Collection("books").Find(context.Background(), filter)
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	err = cs.All(context.Background(), &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookResolver) Topics(ctx context.Context, obj *model.Book) ([]*model.Topic, error) {
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
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var topics []*model.Topic
	err = cs.All(context.Background(), &topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
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
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var authors []*model.Author
	err = cs.All(context.Background(), &authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *bookResolver) Reviews(ctx context.Context, obj *model.Book) ([]*model.Review, error) {
	filter := bson.M{"bookId": obj.ID}
	cs, err := r.DB.Collection("reviews").Find(context.Background(), filter)
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var reviews []*model.Review
	err = cs.All(context.Background(), &reviews)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	now := time.Now().Unix()
	authorData := bson.M{
		"name":    input.Name,
		"created": now,
		"updated": now,
	}
	result, err := r.DB.Collection("authors").InsertOne(context.Background(), authorData)
	if err != nil {
		return nil, err
	}
	return &model.Author{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Name:    input.Name,
		Created: now,
		Updated: now,
	}, nil
}

func (r *mutationResolver) CreateTopic(ctx context.Context, input *model.NewTopic) (*model.Topic, error) {
	now := time.Now().Unix()
	topicData := bson.M{
		"name":    input.Name,
		"created": now,
		"updated": now,
	}
	result, err := r.DB.Collection("topics").InsertOne(context.Background(), topicData)
	if err != nil {
		return nil, err
	}
	return &model.Topic{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Name:    input.Name,
		Created: now,
		Updated: now,
	}, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input *model.NewBook) (*model.Book, error) {
	now := time.Now().Unix()
	bookData := bson.M{
		"name":      input.Name,
		"content":   input.Content,
		"created":   now,
		"updated":   now,
		"topicsId":  input.TopicsID,
		"authorsId": input.AuthorsID,
	}
	result, err := r.DB.Collection("books").InsertOne(context.Background(), bookData)
	if err != nil {
		return nil, err
	}
	return &model.Book{
		ID:        result.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		Content:   input.Content,
		TopicsID:  input.TopicsID,
		AuthorsID: input.AuthorsID,
		Created:   now,
		Updated:   now,
	}, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, input *model.NewReview) (*model.Review, error) {
	now := time.Now().Unix()
	reviewData := bson.M{
		"content": input.Content,
		"created": now,
		"updated": now,
		"bookId":  input.BookID,
	}
	result, err := r.DB.Collection("reviews").InsertOne(context.Background(), reviewData)
	if err != nil {
		return nil, err
	}
	return &model.Review{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Content: input.Content,
		Created: now,
		Updated: now,
		BookID:  input.BookID,
	}, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	cs, err := r.DB.Collection("authors").Find(context.Background(), bson.M{})
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var authors []*model.Author
	err = cs.All(context.Background(), &authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *queryResolver) Topics(ctx context.Context) ([]*model.Topic, error) {
	cs, err := r.DB.Collection("topics").Find(context.Background(), bson.M{})
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var topics []*model.Topic
	err = cs.All(context.Background(), &topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	cs, err := r.DB.Collection("books").Find(context.Background(), bson.M{})
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	err = cs.All(context.Background(), &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *topicResolver) Books(ctx context.Context, obj *model.Topic) ([]*model.Book, error) {
	filter := bson.M{"topicsId": bson.M{"$all": bson.A{obj.ID}}}
	cs, err := r.DB.Collection("books").Find(context.Background(), filter)
	defer cs.Close(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	err = cs.All(context.Background(), &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type topicResolver struct{ *Resolver }
