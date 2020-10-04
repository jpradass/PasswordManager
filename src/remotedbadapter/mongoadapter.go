package remotedbadapter

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/PasswordManager/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection(conf *configuration.Configuration) (*mongo.Client, error, context.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// // defer cancel()
	ctx := context.Background()
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", conf.User, conf.DbPass, conf.Host, conf.DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err, nil
	}
	return client, nil, ctx
}

//SearchPassword ...
//Search for the password of indicated service
func SearchPassword(service string, conf *configuration.Configuration) (string, error) {
	client, err, ctx := getConnection(conf)
	if err != nil {
		return "", err
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		return "", err
	}

	collection := client.Database(conf.DB).Collection(conf.Collection)
	cur, err := collection.Find(ctx, bson.M{"service": service})
	if err != nil {
		return "", err
	}

	var passwordFetched []bson.M
	if err = cur.All(ctx, &passwordFetched); err != nil {
		return "", err
	}
	if len(passwordFetched) > 1 {
		return "", errors.New("Too many passwords fetched, be more specific")
	}
	pwd, ok := passwordFetched[0]["password"].(string)
	if !ok {
		return "", errors.New("Something went wrong trying to fetch the password")
	}
	return pwd, nil
}
