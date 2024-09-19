package cards

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type (
	Repository interface {
		Create(c context.Context, card *Card) error
		Get(c context.Context, id, customerID string) (*Card, error)
		Update(c context.Context, card *Card) error
		Delete(c context.Context, id, customerID string) error
	}

	EncryptionService interface {
		Encrypt(data string) (string, error)
		Decrypt(encryptedData string) (string, error)
	}

	Service struct {
		repository        Repository
		encryptionService EncryptionService
	}
)

func NewService(repository Repository, encryptionService EncryptionService) *Service {
	return &Service{
		repository:        repository,
		encryptionService: encryptionService,
	}
}

func (svc *Service) AddCard(c context.Context, card *Card) error {
	if err := card.Validate(); err != nil {
		return err
	}

	card.ID = uuid.New().String()
	card.CustomerID = (c.Value("customer_id")).(string)
	cardNumber := *card.Details.Number
	card.Last4Digits = cardNumber[len(cardNumber)-4:]

	if err := svc.encryptSensitiveData(card); err != nil {
		return err
	}

	return svc.repository.Create(c, card)
}

func (svc *Service) encryptSensitiveData(card *Card) error {
	sensitiveData := fmt.Sprintf("%s|%s|%s", *card.Details.Number, *card.Details.CVV, *card.Details.ExpirationDate)
	encryptedData, err := svc.encryptionService.Encrypt(sensitiveData)
	if err != nil {
		return err
	}
	card.SensitiveData = encryptedData

	return nil
}

func (svc *Service) GetCardDetails(c context.Context, id string) (*Card, error) {
	customerID := (c.Value("customer_id")).(string)
	card, err := svc.repository.Get(c, id, customerID)
	if err != nil {
		return nil, err
	}
	card.ObfuscateCardNumber()

	return card, nil
}

func (svc *Service) UpdateCardInformation(c context.Context, id string, details *CardDetails) error {
	card, err := svc.GetCardDetails(c, id)
	if err != nil {
		return err
	}

	if err = svc.decryptSensitiveData(card); err != nil {
		return err
	}

	if details.Number != nil {
		card.Details.Number = details.Number
	}

	if details.CVV != nil {
		card.Details.CVV = details.CVV
	}

	if details.ExpirationDate != nil {
		card.Details.ExpirationDate = details.ExpirationDate
	}

	if details.HolderName != nil {
		card.Details.HolderName = details.HolderName
	}

	if err = card.Validate(); err != nil {
		return err
	}

	if err = svc.encryptSensitiveData(card); err != nil {
		return err
	}

	return svc.repository.Update(c, card)
}

func (svc *Service) decryptSensitiveData(card *Card) error {
	data, err := svc.encryptionService.Decrypt(card.SensitiveData)
	if err != nil {
		return err
	}
	sensitiveData := strings.Split(data, "|")

	card.Details.Number = &sensitiveData[0]
	card.Details.CVV = &sensitiveData[1]
	card.Details.ExpirationDate = &sensitiveData[2]

	return nil
}

func (svc *Service) DeleteCard(c context.Context, id string) error {
	customerID := (c.Value("customer_id")).(string)
	if _, err := svc.repository.Get(c, id, customerID); err != nil {
		return err
	}
	return svc.repository.Delete(c, id, customerID)
}
