package infrastructure

import "github.com/luistm/go-bank-cli/entities"

// Storage ...
type Storage interface {
	entities.InfrastructureHandler
	Close() error
}
