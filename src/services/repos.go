package services

import (
	// "database/sql"
	"context"
	"github.com/pkg/errors"
	"github.com/ahsu1230/golangwebservertutorial/src/entities"
	"github.com/ahsu1230/golangwebservertutorial/src/logger"
)

func CreateHero1(ctx context.Context, rowId uint) (uint, error) {
	logger.Info("repo.CreateHeroSuccess", logger.Fields{ 
		"requestUuid": ctx.Value("requestUuid"),
		"rowId": rowId,
	})
	return rowId, nil
}

func CreateHero2(ctx context.Context, rowId uint) (uint, error) {
	logger.Info("repo.CreateHeroSuccess", logger.Fields{ 
		"requestUuid": ctx.Value("requestUuid"),
		"rowId": rowId,
	})
	return 0, errors.Wrap(entities.ErrSQL, "database entry already exists")
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