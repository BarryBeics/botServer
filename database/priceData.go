package database

import (
	"context"

	"time"

	"github.com/barrybeics/botServer/graph/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) SaveHistoricPrices(input *model.NewHistoricPriceInput) ([]*model.HistoricPrices, error) {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a slice to store the inserted HistoricPrices
	var insertedHistoricPrices []*model.HistoricPrices

	// Iterate over pairs and insert each one into the collection
	for _, pairInput := range input.Pairs {
		// Create a new HistoricPrices object for each pair with the provided timestamp
		historicPrices := &model.HistoricPrices{
			Pair:      []*model.Pair{{Symbol: pairInput.Symbol, Price: pairInput.Price}},
			Timestamp: input.Timestamp,
		}

		// Insert the new HistoricPrices object into the collection
		_, err := collection.InsertOne(ctx, historicPrices)
		if err != nil {
			log.Error().Err(err).Msg("Error saving historic price:")
			// Handle the error, perhaps return an error or log it
			return nil, err
		}

		// Append the inserted HistoricPrices to the result slice
		insertedHistoricPrices = append(insertedHistoricPrices, historicPrices)
	}

	// Return the array of inserted HistoricPrices
	return insertedHistoricPrices, nil
}

// HistoricPricesBySymbol fetches historic prices based on the given symbol and limit.
func (db *DB) HistoricPricesBySymbol(symbol string, limit int) ([]model.HistoricPrices, error) {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"pair.symbol": symbol} // Assuming your data model has a nested "pair" field

	cursor, err := collection.Find(ctx, filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		log.Error().Err(err).Msg("Error fetching historic prices by symbol")
		return nil, err
	}
	defer cursor.Close(ctx)

	var historicPrices []model.HistoricPrices
	if err := cursor.All(ctx, &historicPrices); err != nil {
		log.Error().Err(err).Msg("Error decoding historic prices")
		return nil, err
	}

	return historicPrices, nil
}

func (db *DB) AllHistoricPrices(limit int) ([]model.HistoricPrices, error) {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Error().Err(err).Msg("Error querying historic prices:")
		return nil, err
	}
	defer cursor.Close(ctx)

	var historicPrices []model.HistoricPrices
	if err := cursor.All(ctx, &historicPrices); err != nil {
		log.Error().Err(err).Msg("Error decoding historic prices:")
		return nil, err
	}

	return historicPrices, nil
}

// HistoricPricesAtTimestamp fetches historic prices at a specific timestamp.
func (db *DB) HistoricPricesAtTimestamp(timestamp int) ([]model.HistoricPrices, error) {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter by timestamp
	filter := bson.M{"Timestamp": timestamp}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching historic prices at timestamp")
		return nil, err
	}
	defer cursor.Close(ctx)

	var historicPrices []model.HistoricPrices

	// Iterate over the results
	for cursor.Next(ctx) {
		var result model.HistoricPrices
		if err := cursor.Decode(&result); err != nil {
			log.Error().Err(err).Msg("Error decoding historic prices at timestamp")
			return nil, err
		}

		// Append the result to the list
		historicPrices = append(historicPrices, result)
	}

	return historicPrices, nil
}

// GetUniqueTimestampCount fetches the count of unique timestamps.
func (db *DB) GetUniqueTimestampCount(ctx context.Context) (int, error) {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")

	// Use aggregation to get unique timestamps
	pipeline := bson.A{
		bson.D{{"$group", bson.D{{"_id", "$timestamp"}}}},
		bson.D{{"$group", bson.D{{"_id", nil}, {"count", bson.D{{"$sum", 1}}}}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error().Err(err).Msg("Error counting unique timestamps")
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		log.Error().Err(err).Msg("Error decoding unique timestamps")
		return 0, err
	}

	// Extract count from the result
	count := 0
	if len(result) > 0 {
		// Add the switch here to handle both int and int32
		switch v := result[0]["count"].(type) {
		case int:
			count = v
		case int32:
			count = int(v)
		default:
			log.Error().Msgf("Unexpected type for count: %T", v)
		}
	}

	return count, nil
}

func (db *DB) DeleteHistoricPricesByTimestamp(ctx context.Context, timestamp int) error {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")

	// Define a filter to match documents with the specified timestamp
	filter := bson.D{{"timestamp", timestamp}}

	// Perform the delete operation
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msgf("Error deleting historic prices with timestamp %s", timestamp)
		return err
	}

	log.Info().Msgf("Deleted %d historic prices with timestamp %s", result.DeletedCount, timestamp)

	return nil
}
