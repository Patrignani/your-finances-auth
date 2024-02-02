package interfaces

import "go.mongodb.org/mongo-driver/mongo/options"

type ISpecification interface {
	GetSpecification() (map[string]interface{}, *options.FindOptions)
}

type ISpecificationByOne interface {
	GetSpecification() (map[string]interface{}, *options.FindOneOptions)
}
