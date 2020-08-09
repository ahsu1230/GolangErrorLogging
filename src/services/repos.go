package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ahsu1230/golangwebservertutorial/src/entities"
)

func CreateHeroSuccess(c context.Context, rowId uint) (uint, error) {
	return rowId, nil
}

func CreateHeroFailure(c context.Context) (uint, error) {
	return 0, errors.New("database exception: entry already exists")
}

func GetHeroSuccess(c context.Context, name, heroName string) (entities.Hero, error) {
	return entities.Hero{
		name,
		heroName,
	}, nil
}

func GetHeroFailure(c context.Context) (entities.Hero, error) {
	return entities.Hero{}, errors.New("Hero not found")
}