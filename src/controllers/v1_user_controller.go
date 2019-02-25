package controllers

import (
	"../constants"
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type V1UserController struct {
	userService services.V1UserService
	errorHelper helpers.ErrorHelper
}

func V1UserControllerHandler(router *gin.Engine) {

	handler := &V1UserController{
		userService: services.V1UserServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v1/users")
	{
		group.GET(":id", handler.GetById)
		group.POST(":id", handler.UpdateById)
	}

}

// @Summary Get user data
// @Description Get user data by user id
// @Accept  	json
// @Produce  	json
// @Param   	user_id     path    int     true        "User Id"
// @Success	200	{object} 	objects.V1UserObjectResponse
// @Router /v1/users/{user_id} [get]
func (handler *V1UserController) GetById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.userService.GetById(id)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}

func (handler *V1UserController) UpdateById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	requestObject := objects.V1UserObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.userService.UpdateById(id, requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
