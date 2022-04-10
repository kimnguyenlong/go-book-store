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

func (r *wishListResolver) Books(ctx context.Context, obj *model.WishList) ([]*model.Book, error) {
	if len(obj.BooksID) == 0 {
		return nil, nil
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

// WishList returns generated.WishListResolver implementation.
func (r *Resolver) WishList() generated.WishListResolver { return &wishListResolver{r} }

type wishListResolver struct{ *Resolver }
