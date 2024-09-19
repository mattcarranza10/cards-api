package cards

import "errors"

var (
	ErrCardNumberRequired         = errors.New("card number is required")
	ErrCardCVVRequired            = errors.New("card cvv is required")
	ErrCardExpirationDateRequired = errors.New("card expiration date is required")
	ErrCardHolderNameRequired     = errors.New("card holder name is required")
	ErrInvalidCardNumber          = errors.New("invalid card number")
	ErrInvalidCardExpirationDate  = errors.New("invalid card expiration date")
	ErrExpiredCard                = errors.New("card is expired")
	ErrCardNotFound               = errors.New("card not found")
)
