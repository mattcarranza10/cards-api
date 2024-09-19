package handler

import (
	"cards-api/internal/cards"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *cards.Service
}

func NewHandler(service *cards.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddCard(c *gin.Context) {
	cmd := &CreateCardCommand{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	card := cmd.toDomain()
	if err := h.service.AddCard(c, card); err != nil {
		switch {
		case
			errors.Is(err, cards.ErrCardNumberRequired),
			errors.Is(err, cards.ErrCardCVVRequired),
			errors.Is(err, cards.ErrCardExpirationDateRequired),
			errors.Is(err, cards.ErrCardHolderNameRequired),
			errors.Is(err, cards.ErrInvalidCardNumber),
			errors.Is(err, cards.ErrInvalidCardExpirationDate),
			errors.Is(err, cards.ErrExpiredCard):
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, toCreateCardResponse(card))
}

func (h *Handler) GetCardDetails(c *gin.Context) {
	card, err := h.service.GetCardDetails(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, cards.ErrCardNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toGetCardDetailsResponse(card))
}

func (h *Handler) UpdateCardInformation(c *gin.Context) {
	cmd := &UpdateCardInformationCommand{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	cardDetails := cmd.toDomain()
	if err := h.service.UpdateCardInformation(c, c.Param("id"), cardDetails); err != nil {
		switch {
		case
			errors.Is(err, cards.ErrCardNumberRequired),
			errors.Is(err, cards.ErrCardCVVRequired),
			errors.Is(err, cards.ErrCardExpirationDateRequired),
			errors.Is(err, cards.ErrCardHolderNameRequired),
			errors.Is(err, cards.ErrInvalidCardNumber),
			errors.Is(err, cards.ErrInvalidCardExpirationDate),
			errors.Is(err, cards.ErrExpiredCard):
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.Is(err, cards.ErrCardNotFound):
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) DeleteCard(c *gin.Context) {
	if err := h.service.DeleteCard(c, c.Param("id")); err != nil {
		switch {
		case errors.Is(err, cards.ErrCardNotFound):
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) UpdateCardsConcurrently(c *gin.Context) {
	cmd := make(map[string]UpdateCardInformationCommand)
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	total := len(cmd)

	result := make(chan gin.H, 20)
	updateCardAsync := func(id string, cmd *UpdateCardInformationCommand) {
		err := h.service.UpdateCardInformation(c, id, cmd.toDomain())
		result <- gin.H{
			"id":     id,
			"status": getResultStatus(err),
		}
	}

	for k, v := range cmd {
		go updateCardAsync(k, &v)
	}

	response := make([]gin.H, 0, total)
	for i := 0; i < total; i++ {
		response = append(response, <-result)
	}

	close(result)

	c.JSON(http.StatusOK, response)
}

func getResultStatus(err error) string {
	if err != nil {
		return "failed"
	}
	return "success"
}
