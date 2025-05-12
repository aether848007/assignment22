package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/internal/entity"
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) *MongoRepo {
	return &MongoRepo{collection: db.Collection("products")}
}

func (r *MongoRepo) Create(ctx context.Context, p *entity.Product) error {
	_, err := r.collection.InsertOne(ctx, p)
	return err
}

func (r *MongoRepo) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var result entity.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	return &result, err
}

func (r *MongoRepo) Update(ctx context.Context, id string, p *entity.Product) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": p})
	return err
}

func (r *MongoRepo) Delete(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *MongoRepo) List(ctx context.Context, filter bson.M, limit int64, skip int64) ([]entity.Product, error) {
	cursor, err := r.collection.Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	var products []entity.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
