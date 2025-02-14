package storage

import (
	"context"
	"time"

	// "github.com/timoruohomaki/open311togo/server"
	"github.com/timoruohomaki/open311togo/models"
	// "github.com/timoruohomaki/open311togo/telemetry"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// database repository

type DbMethod interface {
	CreateService(s *models.Open311CreateUpdateService) (*models.Open311Service, error)
	GetServices(limit, page int) ([]*models.Open311Service, error)
	GetService(id primitive.ObjectID) (*models.Open311Service, error)
	DeleteService(id primitive.ObjectID) error
	UpdateService(id primitive.ObjectID, update *models.Open311Service) error
}

// database interface

type DbInterface interface {
	GetServiceCollection() *mongo.Collection
}

type MDB struct {
	Dbcli *mongo.Client
}

func (db *MDB) GetServiceCollection() *mongo.Collection {
	return db.Dbcli.Database("open311db").Collection("services")

}

// create connection

func ConnectToNoSql(dsn string) (DbInterface, error) {
	client, err := NewDatabase(dsn)

	if err != nil {
		return nil, err
	}

	return &MDB{
		Dbcli: client,
	}, nil
}

// get mongodb client

func NewDatabase(dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	// TODO
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, err

}
