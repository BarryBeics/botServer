package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/barrybeics/botServer/graph/model"
)

// MutationResolver implementation
func (r *mutationResolver) CreateActivityReport(ctx context.Context, input *model.NewActivityReport) (*model.ActivityReport, error) {
	return db.SaveActivityReport(input), nil
}

// CreateTradeOutcomeReport is the resolver for the createTradeOutcomeReport field.
func (r *mutationResolver) CreateTradeOutcomeReport(ctx context.Context, input *model.NewTradeOutcomeReport) (*model.TradeOutcomeReport, error) {
	return db.SaveTradeOutcomeReport(input), nil
}

// ActivityReport is the resolver for the ActivityReport field.
func (r *queryResolver) ActivityReport(ctx context.Context, id string) (*model.ActivityReport, error) {
	return db.FindActivityReportByID(id), nil
}

// ActivityReports is the resolver for the ActivityReports field.
func (r *queryResolver) ActivityReports(ctx context.Context) ([]*model.ActivityReport, error) {
	return db.AllActivityReports(), nil
}

// TradeOutcomeReport is the resolver for the TradeOutcomeReport field.
func (r *queryResolver) TradeOutcomeReport(ctx context.Context, id string) (*model.TradeOutcomeReport, error) {
	return db.FindTradeOutcomeReportByID(id), nil
}

// TradeOutcomes retrieves trade outcome reports based on the BotName.
func (r *queryResolver) TradeOutcomes(ctx context.Context, BotName string) ([]*model.TradeOutcomeReport, error) {
	return db.TradeOutcomeReportsByBotName(ctx, BotName)
}

// TradeOutcomeReports is the resolver for the TradeOutcomeReports field.
func (r *queryResolver) TradeOutcomeReports(ctx context.Context) ([]*model.TradeOutcomeReport, error) {
	return db.AllTradeOutcomeReports(), nil
}
