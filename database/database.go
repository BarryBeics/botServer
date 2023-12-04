package database

import (
	"context"

	"time"

	"github.com/barrybeics/botServer/graph/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	uri := "mongodb://fudgebot:cookiebot@mongo:27017/go_trading_db"
	log.Info().Str("mongodb_uri", uri).Msg("Connecting to MongoDB")

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.Auth = &options.Credential{
		Username:   "fudgebot",
		Password:   "cookiebot",
		AuthSource: "admin",
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Error().Err(err).Msg("Error client options func:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error ctx func:")
	}

	return &DB{
		client: client,
	}
}

func (db *DB) Close() {
	if db.client != nil {
		db.client.Disconnect(context.Background())
	}
}

func (db *DB) SaveActivityReport(input *model.NewActivityReport) *model.ActivityReport {
	collection := db.client.Database("go_trading_db").Collection("ActivityReports")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Error save func:")
	}
	return &model.ActivityReport{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Timestamp: input.Timestamp,
		Qty:       input.Qty,
		AvgGain:   input.AvgGain,
	}
}

func (db *DB) FindActivityReportByID(ID string) *model.ActivityReport {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error().Err(err).Msg("Error find by func:")
	}
	collection := db.client.Database("go_trading_db").Collection("ActivityReports")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	ActivityReport := model.ActivityReport{}
	res.Decode(&ActivityReport)
	return &ActivityReport
}

func (db *DB) AllActivityReports() []*model.ActivityReport {
	collection := db.client.Database("go_trading_db").Collection("ActivityReports")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error().Err(err).Msg("Error All func:")
	}
	var ActivityReports []*model.ActivityReport
	for cur.Next(ctx) {
		var ActivityReport *model.ActivityReport
		err := cur.Decode(&ActivityReport)
		if err != nil {
			log.Error().Err(err).Msg("Error Decode func:")
		}
		ActivityReports = append(ActivityReports, ActivityReport)
	}
	return ActivityReports
}

func (db *DB) SaveTradeOutcomeReport(input *model.NewTradeOutcomeReport) *model.TradeOutcomeReport {
	collection := db.client.Database("go_trading_db").Collection("TradeOutcomeReports")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Error save func:")
	}
	return &model.TradeOutcomeReport{
		ID:           res.InsertedID.(primitive.ObjectID).Hex(),
		Timestamp:    input.Timestamp,
		OpeningPrice: input.OpeningPrice,
		ClosePrice:   input.ClosePrice,
		Symbol:       input.Symbol,
		Outcome:      input.Outcome,
	}
}

func (db *DB) SaveHistoricPrices(input *model.NewHistoricPriceInput) []*model.HistoricPrices {
	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a slice to store the inserted HistoricPrices
	var insertedHistoricPrices []*model.HistoricPrices

	// Iterate over pairs and insert each one into the collection
	for _, pairInput := range input.Pairs {
		// Create a new HistoricPrices object for each pair
		historicPrices := &model.HistoricPrices{
			Pair: []*model.Pair{
				{
					Symbol: pairInput.Symbol,
					Price:  pairInput.Price,
				},
			},
		}

		// Insert the new HistoricPrices object into the collection
		_, err := collection.InsertOne(ctx, historicPrices)
		if err != nil {
			log.Error().Err(err).Msg("Error saving historic price:")
			// Handle the error, perhaps return an error or log it
			return nil
		}

		// Append the inserted HistoricPrices to the result slice
		insertedHistoricPrices = append(insertedHistoricPrices, historicPrices)
	}

	// Return the array of inserted HistoricPrices
	return insertedHistoricPrices
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

func (db *DB) FindTradeOutcomeReportByID(ID string) *model.TradeOutcomeReport {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error().Err(err).Msg("Error find by func:")
	}
	collection := db.client.Database("go_trading_db").Collection("TradeOutcomeReports")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	TradeOutcomeReport := model.TradeOutcomeReport{}
	res.Decode(&TradeOutcomeReport)
	return &TradeOutcomeReport
}

func (db *DB) AllTradeOutcomeReports() []*model.TradeOutcomeReport {
	collection := db.client.Database("go_trading_db").Collection("TradeOutcomeReports")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error().Err(err).Msg("Error All func:")
	}
	var TradeOutcomeReports []*model.TradeOutcomeReport
	for cur.Next(ctx) {
		var TradeOutcomeReport *model.TradeOutcomeReport
		err := cur.Decode(&TradeOutcomeReport)
		if err != nil {
			log.Error().Err(err).Msg("Error Decode func:")
		}
		TradeOutcomeReports = append(TradeOutcomeReports, TradeOutcomeReport)
	}
	return TradeOutcomeReports
}

// func (db *DB) SaveHistoricPrices(input *model.NewHistoricPriceInput) bool {
// 	collection := db.client.Database("go_trading_db").Collection("HistoricPrices")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := collection.InsertOne(ctx, input)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error saving historic price:")
// 		return false
// 	}

// 	return true
// }

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
