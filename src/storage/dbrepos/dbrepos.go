package dbrepos

import (
	"context"
	"time"

	"github.com/timoruohomaki/open311togo/storage"
)

type mongoDbRepo struct {
	Client	storage.DbInterface
	ctx 	context.Context
}

func NewMongoDbRepo(client storage.DbInterface, ctx context.Context) storage.DbMethod {
	return &mongoDbRepo{
		Client: client,
		ctx:	ctx,
	}
}

const timeout = 3 * time.Second