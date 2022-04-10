package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied!")
	}
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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if !input.Role.IsValid() {
		return nil, fmt.Errorf("Invalid Role")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	userData := bson.M{
		"name":     input.Name,
		"email":    input.Email,
		"password": string(hashedPassword),
		"role":     input.Role,
		"created":  now,
		"updated":  now,
	}
	result, err := r.DB.Collection("users").InsertOne(context.Background(), userData)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       result.InsertedID.(primitive.ObjectID).Hex(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
		Created:  now,
		Updated:  now,
	}, nil
}

func (r *mutationResolver) CreateTopic(ctx context.Context, input model.NewTopic) (*model.Topic, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied!")
	}
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

func (r *mutationResolver) RemoveTopic(ctx context.Context, id string) (*model.Topic, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied")
	}
	topicOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var topic *model.Topic
	filter := bson.M{"_id": topicOID}
	err = r.DB.Collection("topics").FindOneAndDelete(context.Background(), filter).Decode(&topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *mutationResolver) UpdateTopic(ctx context.Context, id string, name string) (*model.Topic, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied")
	}
	topicOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var topic *model.Topic
	filter := bson.M{"_id": topicOID}
	update := bson.M{"$set": bson.M{"name": name}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.DB.Collection("topics").FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied!")
	}
	now := time.Now().Unix()
	bookData := bson.M{
		"name":      input.Name,
		"price":     input.Price,
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
		Price:     input.Price,
		Content:   input.Content,
		TopicsID:  input.TopicsID,
		AuthorsID: input.AuthorsID,
		Created:   now,
		Updated:   now,
	}, nil
}

func (r *mutationResolver) RemoveBook(ctx context.Context, id string) (*model.Book, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied")
	}
	bookOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var book *model.Book
	filter := bson.M{"_id": bookOID}
	err = r.DB.Collection("books").FindOneAndDelete(context.Background(), filter).Decode(&book)
	if err != nil {
		return nil, err
	}
	// remove all reviews
	filter = bson.M{"bookId": id}
	_, err = r.DB.Collection("reviews").DeleteMany(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, update model.BookUpdate) (*model.Book, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth.Role != model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied")
	}
	bookOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var book *model.Book
	filter := bson.M{"_id": bookOID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updateData := bson.M{}
	if update.Name != nil {
		updateData["$set"] = bson.M{"name": *update.Name}
	}
	if update.Content != nil {
		updateData["$set"] = bson.M{"content": *update.Content}
	}
	if len(update.AddingTopicsID) > 0 {
		updateData["$addToSet"] = bson.M{"topicsId": bson.M{"$each": update.AddingTopicsID}}
	}
	if len(update.AddingAuthorsID) > 0 {
		updateData["$addToSet"] = bson.M{"authorsId": bson.M{"$each": update.AddingAuthorsID}}
	}
	err = r.DB.Collection("books").FindOneAndUpdate(context.Background(), filter, updateData, opts).Decode(&book)
	if err != nil {
		return nil, err
	}
	updateData = bson.M{}
	if len(update.RemovingTopicsID) > 0 {
		updateData["$pull"] = bson.M{"topicsId": bson.M{"$in": update.RemovingTopicsID}}
	}
	if len(update.RemovingAuthorsID) > 0 {
		updateData["$pull"] = bson.M{"authorsId": bson.M{"$in": update.RemovingAuthorsID}}
	}
	if _, ok := updateData["$pull"]; ok {
		err = r.DB.Collection("books").FindOneAndUpdate(context.Background(), filter, updateData, opts).Decode(&book)
		if err != nil {
			return nil, err
		}
	}
	return book, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReview) (*model.Review, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	reviewData := bson.M{
		"content": input.Content,
		"created": now,
		"updated": now,
		"bookId":  input.BookID,
		"userId":  auth.UID,
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
		UserID:  auth.UID,
	}, nil
}

func (r *mutationResolver) RemoveReview(ctx context.Context, bookID string, reviewID string) (*model.Review, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	reviewOID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, err
	}
	var review *model.Review
	filter := bson.M{"_id": reviewOID, "bookId": bookID, "userId": auth.UID}
	err = r.DB.Collection("reviews").FindOneAndDelete(context.Background(), filter).Decode(&review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *mutationResolver) UpdateReview(ctx context.Context, bookID string, reviewID string, content string) (*model.Review, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	reviewOID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, err
	}
	var review *model.Review
	filter := bson.M{"_id": reviewOID, "bookId": bookID, "userId": auth.UID}
	update := bson.M{"$set": bson.M{"content": content}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.DB.Collection("reviews").FindOneAndUpdate(context.Background(), filter, update, options).Decode(&review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *mutationResolver) SetCart(ctx context.Context, input model.CartData) (*model.Cart, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"userId": auth.UID}
	update := bson.M{"$set": bson.M{"items": input.Items}}
	opts := options.Update().SetUpsert(true)
	result, err := r.DB.Collection("carts").UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return nil, err
	}
	items := []*model.CartItem{}
	for _, item := range input.Items {
		items = append(items, &model.CartItem{
			BookID:   item.BookID,
			Quantity: item.Quantity,
		})
	}
	return &model.Cart{
		ID:     result.UpsertedID.(primitive.ObjectID).Hex(),
		UserID: auth.UID,
		Items:  items,
	}, nil
}

func (r *mutationResolver) UpdateWishList(ctx context.Context, input model.WishListUpdate) (*model.WishList, error) {
	auth, err := GetAuthFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var wishList *model.WishList
	filter := bson.M{"userId": auth.UID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	if len(input.Add) > 0 {
		addUpdate := bson.M{"$addToSet": bson.M{"booksId": bson.M{"$each": input.Add}}}
		err = r.DB.Collection("wish-lists").FindOneAndUpdate(context.Background(), filter, addUpdate, opts).Decode(&wishList)
		if err != nil {
			return nil, err
		}
	}
	if len(input.Remove) > 0 {
		removeUpdate := bson.M{"$pull": bson.M{"booksId": bson.M{"$in": input.Remove}}}
		err = r.DB.Collection("wish-lists").FindOneAndUpdate(context.Background(), filter, removeUpdate, opts).Decode(&wishList)
		if err != nil {
			return nil, err
		}
	}
	return wishList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
