package services

import (
	"github.com/pkg/errors"
	"github.com/ahsu1230/golangwebservertutorial/src/entities"
)

func CreateHeroSuccess() (uint, error) {
	return 42, nil
}

func CreateHeroFailure() (uint, error) {
	return 0, errors.New("database exception: entry already exists")
}

func GetHeroSuccess() (entities.Hero, error) {
	return entities.Hero{
		"Tony",
		"Iron Man",
	}, nil
}

func GetHeroFailure() (entities.Hero, error) {
	return entities.Hero{}, errors.New("Hero not found")
}