package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {

	mongoHost := GetEnvOrDefault("MONGO_HOST", "localhost")
	mongoPort := GetEnvOrDefault("MONGO_PORT", "27017")
	mongoUser := GetEnvOrDefault("MONGO_USER", "admin")
	mongoPass := GetEnvOrDefault("MONGO_PASS", "password")
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", mongoUser, mongoPass, mongoHost, mongoPort)

	redisHost := GetEnvOrDefault("REDIS_HOST", "localhost")
	redisPort := GetEnvOrDefault("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	mongoClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatalln(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})

	for {
		votes, err := client.LPop(context.Background(), "votes").Bytes()

		if err != nil {
			if err == redis.Nil {
				continue
			}
			log.Fatalln(err)
		}

		if err := _handleVote(mongoClient, votes); err != nil {
			log.Println(err)
			continue
		}

		time.Sleep(1 * time.Second)
	}
}

func _handleVote(db *mongo.Client, data []byte) error {
	vote, err := parseVote(data)
	if err != nil {
		return err
	}
	fmt.Println("Handling vote: ", vote.Vote, "VoterID: ", vote.VoterID)
	filter := bson.D{{"voterid", vote.VoterID}}

	var exist Vote
	err = db.Database("votes").Collection("votes").FindOne(context.Background(), filter).Decode(&exist)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = db.Database("votes").Collection("votes").InsertOne(context.Background(), vote)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	log.Println("updating...")
	//if exist update
	update := bson.D{{"$set", bson.D{{"vote", vote.Vote}}}}
	updateResult, err := db.Database("votes").Collection("votes").UpdateOne(context.Background(), filter, update)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	if err != nil {
		return err
	}

	return nil
}
