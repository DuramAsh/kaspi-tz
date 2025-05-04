package contragent

import (
	"strconv"
	"time"

	"kaspi-tz/internal/domain/person"
)

func (s *Service) ValidateIIN(iin string) (dest person.ValidateIINResponse, err error) {
	if len(iin) != 12 {
		return dest, ErrIINLength
	}

	digits := make([]int, 12)
	for i := range 12 {
		d, err := strconv.Atoi(string(iin[i]))
		if err != nil {
			err = ErrIINCharacters
			return dest, err
		}

		digits[i] = d
	}

	weights1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

	sum := 0
	for i := range 11 {
		sum += digits[i] * weights1[i]
	}

	checksum := sum % 11

	if checksum == 10 {
		sum = 0
		for i := range 11 {
			sum += digits[i] * weights2[i]
		}

		checksum = sum % 11
	}

	if checksum == 10 || digits[11] != checksum {
		err = ErrIINChecksum
		return
	}

	centuryCode := digits[6]
	centuryMap := map[int]int{
		1: 1800,
		2: 1800,
		3: 1900,
		4: 1900,
		5: 2000,
		6: 2000,
	}
	century, ok := centuryMap[centuryCode]
	if !ok {
		err = ErrIINCenturyNumber
		return
	}

	year := century + digits[0]*10 + digits[1]
	month := digits[2]*10 + digits[3]
	day := digits[4]*10 + digits[5]

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if date.Year() != year || date.Month() != time.Month(month) || date.Day() != day {
		err = ErrIINInvalidDate
		return
	}

	dest.DateOfBirth = date.Format("02.01.2006")

	dest.Sex = "male"

	if centuryCode%2 == 0 {
		dest.Sex = "female"
	}
	dest.Correct = true

	return
}
