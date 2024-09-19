package cards

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type (
	Card struct {
		ID            string
		CustomerID    string
		Last4Digits   string
		SensitiveData string
		Details       *CardDetails
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	CardDetails struct {
		Number         *string
		CVV            *string
		ExpirationDate *string
		HolderName     *string
	}
)

func (c *Card) Validate() error {
	if err := c.validateRequiredFields(); err != nil {
		return err
	}

	if err := c.validateNumber(); err != nil {
		return err
	}

	if err := c.validateExpirationDate(); err != nil {
		return err
	}

	return nil
}

func (c *Card) validateRequiredFields() error {
	if c.Details.Number == nil || *c.Details.Number == "" {
		return ErrCardNumberRequired
	}

	if c.Details.CVV == nil || *c.Details.CVV == "" {
		return ErrCardCVVRequired
	}

	if c.Details.ExpirationDate == nil || *c.Details.ExpirationDate == "" {
		return ErrCardExpirationDateRequired
	}

	if c.Details.HolderName == nil || *c.Details.HolderName == "" {
		return ErrCardHolderNameRequired
	}

	return nil
}

func (c *Card) validateNumber() error {
	if len(*c.Details.Number) < 16 {
		return ErrInvalidCardNumber
	}
	return nil
}

func (c *Card) validateExpirationDate() error {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])/([0-9]{2})$`)
	if !re.MatchString(*c.Details.ExpirationDate) {
		return ErrInvalidCardExpirationDate
	}

	parts := re.FindStringSubmatch(*c.Details.ExpirationDate)
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	expirationTime := time.Date(year+2000, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	if expirationTime.Before(time.Now()) {
		return ErrExpiredCard
	}

	return nil
}

func (c *Card) ObfuscateCardNumber() {
	obfuscatedCardNumber := fmt.Sprintf("****-****-****-%s", c.Last4Digits)
	c.Details.Number = &obfuscatedCardNumber
}
