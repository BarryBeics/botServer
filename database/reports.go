package database

import (
	"context"

	"time"

	"github.com/barrybeics/botServer/graph/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		ID:               res.InsertedID.(primitive.ObjectID).Hex(),
		Timestamp:        input.Timestamp,
		BotName:          input.BotName,
		PercentageChange: input.PercentageChange,
		Balance:          input.Balance,
		Symbol:           input.Symbol,
		Outcome:          input.Outcome,
	}
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

// TradeOutcomeReportsByBot retrieves trade outcome reports based on the BotName.
func (db *DB) TradeOutcomeReportsByBotName(ctx context.Context, botName string) ([]*model.TradeOutcomeReport, error) {
	collection := db.client.Database("go_trading_db").Collection("TradeOutcomeReports")

	filter := bson.D{{"botname", botName}}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msg("Error TradeOutcomeReportsByBot func:")
		return nil, err
	}

	var tradeOutcomeReports []*model.TradeOutcomeReport
	for cur.Next(ctx) {
		var tradeOutcomeReport *model.TradeOutcomeReport
		err := cur.Decode(&tradeOutcomeReport)
		if err != nil {
			log.Error().Err(err).Msg("Error Decode func:")
		}
		tradeOutcomeReports = append(tradeOutcomeReports, tradeOutcomeReport)
	}
	return tradeOutcomeReports, nil
}
