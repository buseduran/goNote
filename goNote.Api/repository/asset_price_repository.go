package repository

import (
	"context"
	"time"

	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type assetPriceRepository struct {
	collection *mongo.Collection
}

func NewAssetPriceRepository(collection *mongo.Collection) *assetPriceRepository {
	return &assetPriceRepository{collection: collection}
}

func (t *assetPriceRepository) CreateAssetPrice(assetPrice *domain.AssetPrice) (*mongo.InsertOneResult, error) {
	assetPrice.Timestamp = time.Now()
	return t.collection.InsertOne(context.Background(), assetPrice)
}

func (t *assetPriceRepository) DeleteAssetPrice(assetPriceID string, c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(assetPriceID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid assetPrice ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = t.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (t *assetPriceRepository) UpdateAssetPrice(assetPriceID string, assetPrice *models.UpdateAssetPrice) error {
	objectID, err := primitive.ObjectIDFromHex(assetPriceID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"price": assetPrice.Price,
	}}

	_, err = t.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (t *assetPriceRepository) GetAssetPriceHistory(assetID primitive.ObjectID, startDate time.Time, endDate time.Time, page int, pageSize int, c context.Context) (map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	skip := (page - 1) * pageSize

	// Build the filter
	filter := bson.M{"assetID": assetID}
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
		//return nil, c.Status(500).SendString("Error fetching prices history")
	}
	defer cursor.Close(c)

	var prices []domain.AssetPrice
	if err = cursor.All(c, &prices); err != nil {
		//return nil, c.Status(500).SendString("Error decoding prices history")
	}

	mappedPrices := make([]map[string]interface{}, len(prices))
	for i, price := range prices {
		mappedPrices[i] = price.ResponseMap()
	}

	// Include metadata if needed
	total, err := t.collection.CountDocuments(c, filter)

	response := map[string]interface{}{
		"data":     mappedPrices,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
	}

	return response, nil
}
