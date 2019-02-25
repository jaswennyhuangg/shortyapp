package controllers

import (
	"../constants"
	"../helpers"
	"../services"
	"../objects"
	"github.com/gin-gonic/gin"
	
)

type V1ShortyController struct {
	shortyService services.V1ShortyService
	errorHelper helpers.ErrorHelper
}

func V1ShortyControllerHandler(router *gin.Engine) {

	handler := &V1ShortyController{
		shortyService: services.V1ShortyServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v1/shorty")
	{
		group.GET(":shortcode", handler.GetByShortCode)
		group.GET(":shortcode/stats", handler.GetByShortCodeStats)
		group.POST("shorten", handler.PostShortCode)
	}
}

// @Summary Get user data
// @Description Get user data by user id
// @Accept  	json
// @Produce  	json
// @Param   	user_id     path    int     true        "User Id"
// @Success	200	{object} 	objects.V1UserObjectResponse
// @Router /v1/users/{user_id} [get]
func (handler *V1ShortyController) GetByShortCode(context *gin.Context) {

	shortcode := context.Param("shortcode")
	result, err := handler.shortyService.GetByShortCode(shortcode)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	redirectTo := result.Url

	context.Redirect(302, redirectTo)


}

func (handler *V1ShortyController) GetByShortCodeStats(context *gin.Context) {

	shortcode := context.Param("shortcode")
	result, err := handler.shortyService.GetByShortCodeStats(shortcode)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}
	context.JSON(200, result)

}

func (handler *V1ShortyController) PostShortCode(context *gin.Context) {

	requestObject := &objects.V1ShortyObjectRequest{}
	err:= context.ShouldBind(requestObject)

	if err != nil {
		handler.errorHelper.HTTPResponseError(context,err,constants.RequestParameterInvalid)
	}

	result, errPost := handler.shortyService.PostShortCode(requestObject)
	if nil != errPost {
		handler.errorHelper.HTTPResponseError(context, errPost, constants.InternalServerError)
	}
	
	context.JSON(200, result)

}