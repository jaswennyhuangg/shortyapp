package services

import (
	"../objects"
	"../repositories"
	"../models"
	"fmt"
	"time"
	"github.com/jinzhu/copier"
	"math/rand"
)

type V1ShortyService struct {
	request        objects.V1ShortyObjectResponse
	
	shortyRepository repositories.V1ShortyRepository
}

func V1ShortyServiceHandler() (V1ShortyService) {
	service := V1ShortyService{
		shortyRepository: repositories.V1ShortyRepositoryHandler(),
	}
	return service
}

func rightNow() time.Time {
	return time.Now()
}

func (service *V1ShortyService) GetByShortCode(shortCode string) (objects.V1ShortyObjectResponse, error) {
	shorty, err := service.shortyRepository.GetByShortCode(shortCode)
	if nil != err {
		return objects.V1ShortyObjectResponse{}, err
	}

	shorty.RedirectCount = shorty.RedirectCount+1 
	service.shortyRepository.UpdateCount(shorty)

	var testing = rightNow()

	shorty.LastSeen = &testing
	service.shortyRepository.LastSeen(shorty)

	result := objects.V1ShortyObjectResponse{}
	copier.Copy(&result, &shorty)
	return result, nil

}

func (service *V1ShortyService) GetByShortCodeStats(shortCode string) (objects.V1ShortyObjectStatusResponse, error) {
	shorty, err := service.shortyRepository.GetByShortCodeStats(shortCode)
	if nil != err {
		return objects.V1ShortyObjectStatusResponse{}, err
	}
	result := objects.V1ShortyObjectStatusResponse{}
	copier.Copy(&result, &shorty)
	return result, nil
}

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}

func (service *V1ShortyService) PostShortCode(shorty *objects.V1ShortyObjectRequest) (objects.V1ShortyObjectResponse, error) {
	
	shortyMasterData := &models.Shorty{
		Url: shorty.Url,
		ShortCode : shorty.ShortCode,
	}

	if shortyMasterData.ShortCode == "" {
		length := len(shortyMasterData.Url)
		shortyMasterData.ShortCode = randomString(length)
	}
fmt.Println(shortyMasterData)
	id, errinsert := service.shortyRepository.PostShortCode(shortyMasterData)
fmt.Println(errinsert)
	shortyResult, err := service.shortyRepository.GetById(id)
	fmt.Println(shortyResult) //bnyk atribut
	if nil != err {
		return objects.V1ShortyObjectResponse{}, err
	}
	
	result := objects.V1ShortyObjectResponse{}
	copier.Copy(&result, &shortyResult)

	return result, nil


}


