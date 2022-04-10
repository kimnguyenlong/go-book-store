package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *queryResolver) Login(ctx context.Context, input *model.Login) (string, error) {
	var user model.User
	err := r.DB.Collection("users").FindOne(context.Background(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return "", fmt.Errorf("Email %v doesn't exist", input.Email)
	}
	if !user.CheckPassword(input.Password) {
		return "", fmt.Errorf("Incorrect password")
	}
	tokenString, err := user.CreateJWT()
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	cs, err := r.DB.Collection("authors").Find(context.Background(), bson.M{})
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

func (r *queryResolver) Topics(ctx context.Context) ([]*model.Topic, error) {
	cs, err := r.DB.Collection("topics").Find(context.Background(), bson.M{})
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

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	cs, err := r.DB.Collection("books").Find(context.Background(), bson.M{})
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

func (r *queryResolver) Cart(ctx context.Context) (*model.Cart, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var cart *model.Cart
	filter := bson.M{"userId": auth.UID}
	err = r.DB.Collection("carts").FindOne(context.Background(), filter).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *queryResolver) WishList(ctx context.Context) (*model.WishList, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var wishList *model.WishList
	filter := bson.M{"userId": auth.UID}
	err = r.DB.Collection("carts").FindOne(context.Background(), filter).Decode(&wishList)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return wishList, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
