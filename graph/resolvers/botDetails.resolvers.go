package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/barrybeics/botServer/graph/model"
	"github.com/rs/zerolog/log"
)

// CreateStrategy is the resolver for the createStrategy field.
func (r *mutationResolver) CreateStrategy(ctx context.Context, input model.StrategyInput) (*model.Strategy, error) {
	// Assuming db is an instance of your DB type
	strategy, err := db.CreateStrategy(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Error creating strategy:")
		return nil, err
	}

	return strategy, nil
}

// UpdateStrategy is the resolver for the updateStrategy field.
func (r *mutationResolver) UpdateStrategy(ctx context.Context, botInstanceName string, input model.StrategyInput) (*model.Strategy, error) {
	// Assuming db is an instance of your DB type
	strategy, err := db.UpdateStrategy(ctx, botInstanceName, input)
	if err != nil {
		log.Error().Err(err).Msg("Error updating strategy:")
		return nil, err
	}

	return strategy, nil
}

// DeleteStrategy is the resolver for the deleteStrategy field.
func (r *mutationResolver) DeleteStrategy(ctx context.Context, botInstanceName string) (*bool, error) {
	// Assuming db is an instance of your DB type
	success, err := db.DeleteStrategy(ctx, botInstanceName)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting strategy:")
		return nil, err
	}

	return &success, nil
}

// UpdateCounters is the resolver for the updateCounters field.
func (r *mutationResolver) UpdateCounters(ctx context.Context, input model.UpdateCountersInput) (*bool, error) {
	// Use input values to call the underlying database operation
	err := db.UpdateCountersAndBalance(ctx, input.BotInstanceName, *input.WINCounter, *input.LOSSCounter, *input.TIMEOUTGainCounter, *input.TIMEOUTLossCounter, input.AccountBalance, *input.FeesTotal)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update counters.")
		return nil, err
	}

	success := true

	return &success, nil
}

// MarkAsTested is the resolver for the markAsTested field.
func (r *mutationResolver) MarkAsTested(ctx context.Context, input model.MarkAsTestedInput) (*bool, error) {
	err := db.UpdateTested(ctx, input.BotInstanceName, input.Tested)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update strategy is tested status.")
		return nil, err
	}

	success := true
	return &success, nil
}

// GetStrategyByName is the resolver for the getStrategyByName field.
func (r *queryResolver) GetStrategyByName(ctx context.Context, botInstanceName string) (*model.Strategy, error) {
	// Assuming db is an instance of your DB type
	strategy, err := db.GetStrategyByName(ctx, botInstanceName)
	if err != nil {
		log.Error().Err(err).Msg("Error getting strategy by name:")
		return nil, err
	}

	return strategy, nil
}

// GetAllStrategies is the resolver for the getAllStrategies field.
func (r *queryResolver) GetAllStrategies(ctx context.Context) ([]*model.Strategy, error) {
	// Assuming db is an instance of your DB type
	strategies, err := db.GetAllStrategies(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all strategies:")
		return nil, err
	}

	return strategies, nil
}
