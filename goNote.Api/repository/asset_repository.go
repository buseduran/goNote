package repository

import (
	"context"
	"log"
	"time"

	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type assetRepository struct {
	collection *mongo.Collection
}

func NewAssetRepository(collection *mongo.Collection) *assetRepository {
	return &assetRepository{collection: collection}
}

func (t *assetRepository) CreateAsset(asset *domain.Asset) (*mongo.InsertOneResult, error) {
	asset.CreatedAt = time.Now()
	asset.UpdatedAt = time.Now()
	return t.collection.InsertOne(context.Background(), asset)
}
func (t *assetRepository) DeleteAsset(assetID string, c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid asset ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = t.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
func (t *assetRepository) UpdateAsset(assetID string, asset *domain.Asset) error {
	objectID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":         asset.Name,
		"base_unit":    asset.BaseUnit,
		"value_in_try": asset.ValueInTRY,
		"updated_at":   time.Now(),
	}}
	_, err = t.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (t *assetRepository) GetAll() (*[]domain.Asset, error) {
	var assets []domain.Asset
	cursor, err := t.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var asset domain.Asset
		cursor.Decode(&asset)
		assets = append(assets, asset)
	}
	return &assets, nil
}
func (t *assetRepository) CreateUserAsset(userAsset *domain.UserAsset) (*mongo.InsertOneResult, error) {
	userAsset.CreatedAt = time.Now()
	userAsset.UpdatedAt = time.Now()
	return t.collection.InsertOne(context.Background(), userAsset)
}

func (t *assetRepository) GetUserAssetPagination(userID primitive.ObjectID, startDate time.Time, endDate time.Time, page int, pageSize int, c context.Context) (map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	skip := (page - 1) * pageSize
	filter := bson.M{"userID": userID}
	if !startDate.IsZero() && !endDate.IsZero() {
		filter["timestamp"] = bson.M{
			"$gte": startDate,
			"$lte": endDate,
		}
	} else if !startDate.IsZero() {
		filter["timestamp"] = bson.M{"$gte": startDate}
	} else if !endDate.IsZero() {
		filter["timestamp"] = bson.M{"$lte": endDate}
	}
	cursor, err := t.collection.Find(
		c,
		filter,
		options.Find().
			SetSkip(int64(skip)).
			SetLimit(int64(pageSize)).
			SetSort(bson.M{"timestamp": -1}),
	)
	if err != nil {
		//return error
	}
	defer cursor.Close(c)

	var userAssets []domain.AssetPrice
	if err = cursor.All(c, &userAssets); err != nil {
		//return err
	}
	mappedUserAssets := make([]map[string]interface{}, len(userAssets))
	for i, userAsset := range userAssets {
		mappedUserAssets[i] = userAsset.ResponseMap()
	}
	total, err := t.collection.CountDocuments(c, filter)
	response := map[string]interface{}{
		"data":     mappedUserAssets,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
	}
	return response, nil
}
