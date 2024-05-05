package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/web"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context)  {
    fmt.Println("Hello world")
    err := web.APIError{}
    fmt.Println(err)
     c.JSON(http.StatusOK, struct {
        Message string `json:"message"`
    }{
        Message: "Success",
    })
}
