package contragent

import "errors"

var (
	ErrIINLength        = errors.New("IIN length must be 12 characters")
	ErrIINCharacters    = errors.New("IIN must contain only digits")
	ErrIINChecksum      = errors.New("Invalid IIN checksum")
	ErrIINCenturyNumber = errors.New("Invalid century number")
	ErrIINInvalidDate   = errors.New("Invalid date in IIN")
)
