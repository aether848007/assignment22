package repository

import (
	"context"
	"order-service/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepo struct {
	coll *mongo.Collection
}

func NewMongoOrderRepo(db *mongo.Database) *MongoOrderRepo {
	return &MongoOrderRepo{coll: db.Collection("orders")}
}

func (r *MongoOrderRepo) Create(ctx context.Context, o *entity.Order) error {
	o.ID = primitive.NewObjectID()
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	_, err := r.coll.InsertOne(ctx, o)
	return err
}

func (r *MongoOrderRepo) UpdateStatus(ctx context.Context, id string, status string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{"status": status, "updated_at": time.Now()},
	})
	return err
}

func (r *MongoOrderRepo) GetByID(ctx context.Context, id string) (*entity.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var order entity.Order
	err = r.coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	return &order, err
}

func (r *MongoOrderRepo) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *MongoOrderRepo) List(ctx context.Context) ([]*entity.Order, error) {
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*entity.Order
	for cursor.Next(ctx) {
		var o entity.Order
		if err := cursor.Decode(&o); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}
	return orders, nil
}
