package dbrepos

import (
	"context"
	"errors"
	"github.com/timoruohomaki/open311togo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDbRepo) CreateService(p *models.Open311CreateUpdateService) (*models.Open311Service, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
   
	res, err := m.Client.GetServiceCollection().InsertOne(ctx, p)
	if err != nil {
	 return nil, errors.New("error in inserting service")
	}
	prodct, _ := m.GetService(res.InsertedID.(primitive.ObjectID))
	return prodct, nil
   }

func (m *mongoDbRepo) GetServices(limit, page int) ([]*models.Open311Service, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
   
	if limit == 0 || limit < 0 {
	 limit = 10
	}
	if page == 0 || page < 0 {
	 page = 1
	}
	skip := (page - 1) * limit
   
	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
   
	query := bson.M{}
   
	res, err := m.Client.GetServiceCollection().Find(ctx, query, &opt)
	if err != nil {
	 return nil, errors.New("error fetching all available Open311 services")
	}
	services := []*models.Open311Service{}
	for res.Next(ctx) {
	 service := &models.Open311Service{}
	 if err := res.Decode(&service); err != nil {
	  return nil, errors.New("error in scanning service")
	 }
	 services = append(services, service)
	}
	return services, nil
}

func (m *mongoDbRepo) GetService(id primitive.ObjectID) (*models.Open311Service, error) {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
   
	res := m.Client.GetServiceCollection().FindOne(ctx, bson.M{"_id": id})
	product := &models.Open311Service{}
	if err := res.Decode(&product); err != nil {
	 if err == mongo.ErrNoDocuments {
	  return nil, errors.New("service not found")
	 }
	 return nil, errors.New("error in fetching service")
	}
	return product, nil
}

func (m *mongoDbRepo) DeleteService(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
   
	res, err := m.Client.GetServiceCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
	 return errors.New("error in deleting service")
	}
	if res.DeletedCount == 0 {
	 return errors.New("no service deleted")
	}
	return nil
}

func (m *mongoDbRepo) UpdateService(id primitive.ObjectID, update *models.Open311CreateUpdateService) error {
	ctx, cancel := context.WithTimeout(m.ctx, timeout)
	defer cancel()
	updateQuery := bson.D{{"$set", update}}
	_, err := m.Client.GetServiceCollection().UpdateOne(ctx, bson.M{"_id": id}, updateQuery)
	if err != nil {
	 return err
	}
   
	return nil
   
}
