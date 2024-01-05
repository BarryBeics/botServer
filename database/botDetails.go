package database

import (
	"context"

	"github.com/barrybeics/botServer/graph/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateStrategy creates a new strategy in the database.
func (db *DB) CreateStrategy(ctx context.Context, input model.StrategyInput) (*model.Strategy, error) {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	// Convert StrategyInput to Strategy model
	strategy := &model.Strategy{
		BotInstanceName:      input.BotInstanceName,
		TradeDuration:        input.TradeDuration,
		IncrementsAtr:        input.IncrementsAtr,
		LongSMADuration:      input.LongSMADuration,
		ShortSMADuration:     input.ShortSMADuration,
		WINCounter:           input.WINCounter,
		LOSSCounter:          input.LOSSCounter,
		TIMEOUTGainCounter:   input.TIMEOUTGainCounter,
		TIMEOUTLossCounter:   input.TIMEOUTLossCounter,
		AccountBalance:       input.AccountBalance,
		MovingAveMomentum:    input.MovingAveMomentum,
		TakeProfitPercentage: &input.TakeProfitPercentage,
		StopLossPercentage:   &input.StopLossPercentage,
		Owner:                &input.Owner,
		CreatedOn:            input.CreatedOn,
	}

	_, err := collection.InsertOne(ctx, strategy)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting strategy into the database:")
		return nil, err
	}

	return strategy, nil
}

// UpdateStrategy updates an existing strategy in the database.
func (db *DB) UpdateStrategy(ctx context.Context, botInstanceName string, input model.StrategyInput) (*model.Strategy, error) {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	// Convert StrategyInput to Strategy model
	updatedStrategy := &model.Strategy{
		BotInstanceName:      input.BotInstanceName,
		TradeDuration:        input.TradeDuration,
		IncrementsAtr:        input.IncrementsAtr,
		LongSMADuration:      input.LongSMADuration,
		ShortSMADuration:     input.ShortSMADuration,
		WINCounter:           input.WINCounter,
		LOSSCounter:          input.LOSSCounter,
		TIMEOUTGainCounter:   input.TIMEOUTGainCounter,
		TIMEOUTLossCounter:   input.TIMEOUTLossCounter,
		AccountBalance:       input.AccountBalance,
		MovingAveMomentum:    input.MovingAveMomentum,
		TakeProfitPercentage: &input.TakeProfitPercentage,
		StopLossPercentage:   &input.StopLossPercentage,
		Owner:                &input.Owner,
		CreatedOn:            input.CreatedOn,
	}

	filter := bson.D{{"botinstancename", botInstanceName}}
	update := bson.D{{"$set", updatedStrategy}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("Error updating strategy in the database:")
		return nil, err
	}

	return updatedStrategy, nil
}

// DeleteStrategy deletes a strategy from the database.
func (db *DB) DeleteStrategy(ctx context.Context, botInstanceName string) (bool, error) {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	filter := bson.D{{"botinstancename", botInstanceName}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting strategy from the database:")
		return false, err
	}

	return result.DeletedCount > 0, nil
}

// GetStrategyByName retrieves a strategy from the database by its name.
func (db *DB) GetStrategyByName(ctx context.Context, botInstanceName string) (*model.Strategy, error) {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	filter := bson.D{{"botinstancename", botInstanceName}}

	var strategy model.Strategy
	err := collection.FindOne(ctx, filter).Decode(&strategy)
	if err != nil {
		log.Error().Err(err).Msg("Error getting strategy from the database:")
		return nil, err
	}

	return &strategy, nil
}

// GetAllStrategies retrieves all strategies from the database.
func (db *DB) GetAllStrategies(ctx context.Context) ([]*model.Strategy, error) {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error().Err(err).Msg("Error querying all strategies:")
		return nil, err
	}
	defer cursor.Close(ctx)

	var strategies []*model.Strategy
	if err := cursor.All(ctx, &strategies); err != nil {
		log.Error().Err(err).Msg("Error decoding all strategies:")
		return nil, err
	}

	return strategies, nil
}

// UpdateCountersAndBalance updates WIN, LOSS, TIMEOUT counters, and closingBalance in the database for a specific strategy.
func (db *DB) UpdateCountersAndBalance(ctx context.Context, botInstanceName string, incrementWIN, incrementLOSS, incrementTIMEOUTGain, incrementTIMEOUTLoss bool, accountBalance float64) error {
	collection := db.client.Database("go_trading_db").Collection("BotDetails")

	filter := bson.D{{"botinstancename", botInstanceName}}

	update := bson.D{
		{"$inc", bson.D{
			{"wincounter", func() int {
				if incrementWIN {
					return 1
				}
				return 0
			}()},
			{"losscounter", func() int {
				if incrementLOSS {
					return 1
				}
				return 0
			}()},
			{"timeoutgaincounter", func() int {
				if incrementTIMEOUTGain {
					return 1
				}
				return 0
			}()},
			{"timeoutlosscounter", func() int {
				if incrementTIMEOUTLoss {
					return 1
				}
				return 0
			}()},
		}},
		{"$set", bson.D{
			{"accountbalance", accountBalance},
		}},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Msg("Error updating counters and closing balance in the database:")
		return err
	}

	return nil
}
