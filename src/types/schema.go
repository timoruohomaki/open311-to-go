package types

type FeatureOfInterest struct {
	ID int
}

type ServiceDefinition struct {
	ID int
}

type ServiceRequest struct {
	ID int
}

type Defaults struct {
	MongoServiceCollection string
	MongoRequestCollection string
}
