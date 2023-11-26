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
