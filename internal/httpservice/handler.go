package httpservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kingledion/ent-demo/internal"
	"github.com/kingledion/ent-demo/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler wraps the public facing endpoint handlers for the cohort service.
type Handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return Handler{service: service}
}

func (h *Handler) AddOrder(c *gin.Context) {
	// swagger:route POST /v1/order order addorder
	//
	// Add an order to the repository
	//
	// An order contains both information about the User and Merchant assocaited
	// with that order. Places this information into the repostiory and creates
	// the order relationship.
	//
	// Responses:
	//   200: ok
	//   400: stderr

	var request internal.Order
	err := c.Bind(&request)
	if err != nil {
		log.Printf("Could not process request: %v", err)
		c.Abort()
		c.JSON(http.StatusBadRequest, fmt.Errorf("Could not bind request: %v", err))
		return
	}

	err = h.service.AddOrder(c, request)
	if err != nil {
		log.Printf("Could not process request: %v", err)
		c.Abort()
		c.JSON(http.StatusBadRequest, fmt.Errorf("Could not process request: %v", err))
		return
	}

	c.Status(http.StatusOK)
	return
}

func (h *Handler) Recommend(c *gin.Context) {
	// swagger:route GET /v1/:userid/recommend order addorder
	//
	// Get all orders made by a user
	//
	// Retrieve information on the user identified by their id, and then return
	// all merchants at which that user placed an order.
	//
	// Responses:
	//   200: ok
	//   400: merchantlist

	uid := c.Param("userid")

	// check if UUID is valid

	merchants, err := h.service.Recommend(c, uid)
	if err != nil {
		log.Printf("Could not process request: %v", err)
		c.Abort()
		c.JSON(http.StatusBadRequest, fmt.Errorf("Could not process request: %v", err))
		return
	}

	c.JSON(http.StatusOK, merchants)
	return
}
