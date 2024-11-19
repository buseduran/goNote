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

func (t *assetPriceRepository) GetAssetPriceHistory(c *fiber.Ctx) (map[string]interface{}, error) {
	assetID := c.Query("assetID")
	if assetID == "" {
		return nil, c.Status(400).SendString("AssetID is required")
	}

	// Convert assetID to ObjectID
	objAssetID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return nil, c.Status(400).SendString("Invalid AssetID")
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 50)

	if page < 1 {
		page = 1
	}
	skip := (page - 1) * pageSize

	// Parse startDate and endDate to time.Time
	var parsedStartDate, parsedEndDate time.Time
	if startDate != "" {
		parsedStartDate, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, c.Status(400).SendString("Invalid startDate format. Use YYYY-MM-DD")
		}
	}
	if endDate != "" {
		parsedEndDate, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, c.Status(400).SendString("Invalid endDate format. Use YYYY-MM-DD")
		}
	}

	// Build the filter
	filter := bson.M{"assetID": objAssetID}
	if !parsedStartDate.IsZero() && !parsedEndDate.IsZero() {
		filter["timestamp"] = bson.M{
			"$gte": parsedStartDate,
			"$lte": parsedEndDate,
		}
	} else if !parsedStartDate.IsZero() {
		filter["timestamp"] = bson.M{"$gte": parsedStartDate}
	} else if !parsedEndDate.IsZero() {
		filter["timestamp"] = bson.M{"$lte": parsedEndDate}
	}

	cursor, err := t.collection.Find(
		c.Context(),
		filter,
		options.Find().
			SetSkip(int64(skip)).
			SetLimit(int64(pageSize)).
			SetSort(bson.M{"timestamp": -1}),
	)
	if err != nil {
		return nil, c.Status(500).SendString("Error fetching prices history")
	}
	defer cursor.Close(c.Context())

	var prices []domain.AssetPrice
	if err = cursor.All(c.Context(), &prices); err != nil {
		return nil, c.Status(500).SendString("Error decoding prices history")
	}

	mappedPrices := make([]map[string]interface{}, len(prices))
	for i, price := range prices {
		mappedPrices[i] = price.ResponseMap()
	}

	// Include metadata if needed
	total, err := t.collection.CountDocuments(c.Context(), filter)
	if err != nil {
		return nil, c.Status(500).SendString("Error counting total documents")
	}

	response := map[string]interface{}{
		"data":     mappedPrices,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
	}

	return response, nil
}
