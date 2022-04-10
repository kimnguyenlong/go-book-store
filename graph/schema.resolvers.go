package graph

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

func (r *cartItemResolver) Book(ctx context.Context, obj *model.CartItem) (*model.Book, error) {
	bookId, err := primitive.ObjectIDFromHex(obj.BookID)
	if err != nil {
		return nil, err
	}
	var book *model.Book
	err = r.DB.Collection("books").FindOne(context.Background(), bson.M{"_id": bookId}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

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
	if err != nil {
		return nil, err
	}
	return wishList, nil
}

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

func (r *wishListResolver) Books(ctx context.Context, obj *model.WishList) ([]*model.Book, error) {
	if len(obj.BooksID) == 0 {
		return []*model.Book{}, nil
	}
	var booksId []primitive.ObjectID
	for _, id := range obj.BooksID {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		booksId = append(booksId, objId)
	}
	filter := bson.M{"_id": bson.M{"$in": booksId}}
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

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// CartItem returns generated.CartItemResolver implementation.
func (r *Resolver) CartItem() generated.CartItemResolver { return &cartItemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

// WishList returns generated.WishListResolver implementation.
func (r *Resolver) WishList() generated.WishListResolver { return &wishListResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type cartItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type topicResolver struct{ *Resolver }
type wishListResolver struct{ *Resolver }
