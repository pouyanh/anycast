package repository

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/actor"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type InMemoryServantRepository []prosecution.Servant

func (repo InMemoryServantRepository) Add(servant prosecution.Servant) error {
	return fmt.Errorf("not implemented")
}

func (repo InMemoryServantRepository) Remove(sid int) error {
	return fmt.Errorf("not implemented")
}

func (repo InMemoryServantRepository) GetByID(sid int) (prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo InMemoryServantRepository) GetAll() ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo InMemoryServantRepository) FindByService(topic string) ([]prosecution.Servant, error) {
	result := make([]prosecution.Servant, 0)
	for _, servant := range repo {
		result = append(result, servant)
	}

	return result, nil
}

func (repo InMemoryServantRepository) FindByLocation(location prosecution.Point) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

type RelationalServantRepository struct {
	DB actor.RelationalDatabase
}

func (repo RelationalServantRepository) Add(servant prosecution.Servant) error {
	return fmt.Errorf("not implemented")
}

func (repo RelationalServantRepository) Remove(sid int) error {
	return fmt.Errorf("not implemented")
}

func (repo RelationalServantRepository) GetByID(sid int) (prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo RelationalServantRepository) GetAll() ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo RelationalServantRepository) FindByService(topic string) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo RelationalServantRepository) FindByLocation(location prosecution.Point) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

type OnlineServantRepository struct {
	Dictionary actor.Dictionary
}

func (repo OnlineServantRepository) Add(servant prosecution.Servant) error {
	return fmt.Errorf("not implemented")
}

func (repo OnlineServantRepository) Remove(sid int) error {
	return fmt.Errorf("not implemented")
}

func (repo OnlineServantRepository) GetByID(sid int) (prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo OnlineServantRepository) GetAll() ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo OnlineServantRepository) FindByService(topic string) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo OnlineServantRepository) FindByLocation(location prosecution.Point) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}
