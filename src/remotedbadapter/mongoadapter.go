package remotedbadapter

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/PasswordManager/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type document struct {
	Date     time.Time `bson:"date"`
	Service  string    `bson:"service"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
}

func getConnection(conf *configuration.Configuration) (*mongo.Client, context.Context, error) {
	ctx := context.Background()
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", conf.User, conf.DbPass, conf.Host, conf.DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, err
	}
	return client, ctx, err
}

//SearchItem ...
//Search for the item of indicated service
func SearchItem(service string, itemdesc string, conf *configuration.Configuration) ([]byte, error) {
	collection, ctx, client, err := retrieveCollection(conf)
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)

	cur, err := collection.Find(ctx, bson.M{"service": service})
	if err != nil {
		return nil, err
	}

	var itemsFetched []bson.M
	if err = cur.All(ctx, &itemsFetched); err != nil {
		return nil, err
	}

	if len(itemsFetched) > 1 {
		return nil, fmt.Errorf("Too many %ss fetched, be more specific", itemdesc)
	}
	if len(itemsFetched) == 0 {
		return nil, fmt.Errorf("No %ss fetched. Do you have a typo in your service?", itemdesc)
	}

	item, ok := itemsFetched[0][itemdesc].(string)
	if !ok {
		return nil, fmt.Errorf("Something went wrong trying to fetch the %s", itemdesc)
	}

	byteitem, err := base64.StdEncoding.DecodeString(item)
	if err != nil {
		return nil, err
	}
	return byteitem, nil
}

//UpdateItem ...
//Update item based on what is given
func UpdateItem(service string, item []byte, itemdesc string, conf *configuration.Configuration) (string, error) {
	collection, ctx, client, err := retrieveCollection(conf)
	if err != nil {
		return "", err
	}

	defer client.Disconnect(ctx)

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"service": service},
		bson.D{
			{"$set", bson.D{{itemdesc, base64.StdEncoding.EncodeToString(item)}}},
		},
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s updated!", itemdesc), nil
}

//InsertService ...
//Insert a new service into the db
func InsertService(service string, user []byte, pwd []byte, conf *configuration.Configuration) (string, error) {
	collection, ctx, client, err := retrieveCollection(conf)
	if err != nil {
		return "", err
	}

	defer client.Disconnect(ctx)

	_, err = collection.InsertOne(ctx, document{
		Date:     time.Now(),
		Service:  service,
		Username: base64.StdEncoding.EncodeToString(user),
		Password: base64.StdEncoding.EncodeToString(pwd),
	})
	if err != nil {
		return "", err
	}
	return "Inserted new service", nil
}

//RemoveService ...
//Remove a given service from db
func RemoveService(service string, conf *configuration.Configuration) (string, error) {
	collection, ctx, client, err := retrieveCollection(conf)
	if err != nil {
		return "", err
	}

	defer client.Disconnect(ctx)

	_, err = collection.DeleteOne(
		ctx,
		bson.M{"service": service},
	)
	if err != nil {
		return "", err
	}
	return "Deleted service!", nil
}

func retrieveCollection(conf *configuration.Configuration) (*mongo.Collection, context.Context, *mongo.Client, error) {
	client, ctx, err := getConnection(conf)
	if err != nil {
		return nil, nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	return client.Database(conf.DB).Collection(conf.Collection), ctx, client, nil
}
