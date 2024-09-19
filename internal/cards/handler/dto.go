package handler

import (
	"cards-api/internal/cards"
	"time"
)

type (
	CreateCardCommand struct {
		Number         string `json:"number"`
		CVV            string `json:"cvv"`
		ExpirationDate string `json:"expiration_date"`
		HolderName     string `json:"holder_name"`
	}

	UpdateCardInformationCommand struct {
		Number         *string `json:"number,omitempty"`
		CVV            *string `json:"cvv,omitempty"`
		ExpirationDate *string `json:"expiration_date,omitempty"`
		HolderName     *string `json:"holder_name,omitempty"`
	}

	CreateCardResponse struct {
		ID string `json:"id"`
	}

	GetCardDetailsResponse struct {
		ID         string    `json:"id"`
		CustomerID string    `json:"customer_id"`
		Number     string    `json:"number"`
		HolderName string    `json:"holder_name"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func (cc *CreateCardCommand) toDomain() *cards.Card {
	return &cards.Card{
		Details: &cards.CardDetails{
			Number:         &cc.Number,
			CVV:            &cc.CVV,
			ExpirationDate: &cc.ExpirationDate,
			HolderName:     &cc.HolderName,
		},
	}
}

func (uc *UpdateCardInformationCommand) toDomain() *cards.CardDetails {
	return &cards.CardDetails{
		Number:         uc.Number,
		CVV:            uc.CVV,
		ExpirationDate: uc.ExpirationDate,
		HolderName:     uc.HolderName,
	}
}

func toCreateCardResponse(card *cards.Card) *CreateCardResponse {
	return &CreateCardResponse{
		ID: card.ID,
	}
}

func toGetCardDetailsResponse(card *cards.Card) *GetCardDetailsResponse {
	return &GetCardDetailsResponse{
		ID:         card.ID,
		CustomerID: card.CustomerID,
		Number:     *card.Details.Number,
		HolderName: *card.Details.HolderName,
		CreatedAt:  card.CreatedAt,
		UpdatedAt:  card.UpdatedAt,
	}
}
