package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get purchase with id
// @Tags purchases
// @Description get purchase with id
// @ID api
// @Param id  path string  true  "Uid of purchase"
// @Produce  json
// @Success 200 {object} models.PurchaseDTO
// @Failure 400,404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /purchase/get/{id} [get]
func (h *Handler) GetById(c *gin.Context) {
	uid := c.Param("id")
	logrus.Debugf("receive api GET with uid=%s", uid)
	purchase, err := h.repository.GetPurchaseByUidCache(uid)
	if err != nil {
		logrus.Errorf("error while getting purchase with uid=%s, err=%s", uid, err.Error())
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("this uid=%s is probably not presented on server", uid))
		return
	}
	c.JSON(http.StatusOK, purchase)

}
