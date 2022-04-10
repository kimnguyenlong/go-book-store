package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book-store/graph/generated"
	"book-store/graph/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *cartItemResolver) Book(ctx context.Context, obj *model.CartItem) (*model.Book, error) {
	bookId, err := primitive.ObjectIDFromHex(obj.BookID)
	if err != nil {
		return nil, err
	}
	var book *model.Book
	err = r.DB.Collection("books").FindOne(context.Background(), bson.M{"_id": bookId}).Decode(&book)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return book, nil
}

// CartItem returns generated.CartItemResolver implementation.
func (r *Resolver) CartItem() generated.CartItemResolver { return &cartItemResolver{r} }

type cartItemResolver struct{ *Resolver }
