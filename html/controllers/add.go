package controllers

import (
	"add/interfaces"
	"add/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
type AddController struct{
	AddService  interfaces.Add
}
func InitController(addservice interfaces.Add) AddController {
    return AddController{addservice}
}
func (t *AddController)Add(ctx *gin.Context){
    var trans *models.Add  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans:= t.AddService.Add(trans)
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}