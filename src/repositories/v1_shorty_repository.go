package repositories

import (
	"../database"
	"../models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type V1ShortyRepository struct {
	DB gorm.DB
}

func V1ShortyRepositoryHandler() (V1ShortyRepository) {
	repository := V1ShortyRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1ShortyRepository) GetByShortCode(shortCode string) (models.Shorty, error) {

	shortyResponse := models.Shorty{}

	query := repository.DB.Table("rl_shorty")
	query = query.Where("shortcode=?", shortCode)
	query = query.First(&shortyResponse)

	return shortyResponse, query.Error

}

func (repository *V1ShortyRepository) GetByShortCodeStats(shortCode string) (models.Shorty, error) {

	shortyResponse := models.Shorty{}

	query := repository.DB.Table("rl_shorty")
	query = query.Where("shortcode=?", shortCode)
	query = query.First(&shortyResponse)

	return shortyResponse, query.Error

}

func (repository *V1ShortyRepository) GetById(id int) (models.Shorty, error) {

	shortyResponse := models.Shorty{}
	query := repository.DB.Table("rl_shorty")
	query = query.Where("id=?", id)
	query = query.First(&shortyResponse)

	return shortyResponse, query.Error

}

func (repository *V1ShortyRepository) PostShortCode(Shorty *models.Shorty) (int,error) {
	query := repository.DB.Table("rl_shorty").Create(Shorty)
	return int(Shorty.ID),query.Error
}


func (repository *V1ShortyRepository) UpdateCount(shortyData models.Shorty) (models.Shorty, error) {

	fmt.Println(shortyData)

	query := repository.DB.Table("rl_shorty")
	query = query.Where("shortcode=?", shortyData.ShortCode)
	query = query.Updates(shortyData)

	return shortyData, query.Error
}

func (repository *V1ShortyRepository) LastSeen(shortyData models.Shorty) (models.Shorty, error) {

	fmt.Println(shortyData)

	query := repository.DB.Table("rl_shorty")
	query = query.Where("shortcode=?", shortyData.ShortCode)
	query = query.Updates(shortyData)

	return shortyData, query.Error
}
