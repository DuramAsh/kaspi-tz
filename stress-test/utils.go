package main

import (
	"math/rand"
	"strconv"
)

func generateRandomPerson() (dest Person) {
	iin := generateIIN(rand.Float32() <= 0.9) // 90% что валидный
	name := randomName()
	phone := "7777" + strconv.Itoa(rand.Intn(9000000)+1000000) // поидее можно и константу, но в базе прописал unique constaint :)

	dest = Person{
		Name:  name,
		IIN:   iin,
		Phone: phone,
	}

	return
}

func generateIIN(valid bool) (dest string) {
	digits := make([]int, 12)

	year := rand.Intn(30) + 1990
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1

	digits[0] = (year / 10 % 10)
	digits[1] = year % 10
	digits[2] = month / 10
	digits[3] = month % 10
	digits[4] = day / 10
	digits[5] = day % 10

	switch {
	case year >= 1800 && year < 1900:
		digits[6] = rand.Intn(2) + 1
	case year >= 1900 && year < 2000:
		digits[6] = rand.Intn(2) + 3
	default:
		digits[6] = rand.Intn(2) + 5
	}

	for i := 7; i <= 10; i++ {
		digits[i] = rand.Intn(10)
	}

	digits[11] = calculateChecksum(digits)

	if !valid {
		digits[11] = (digits[11] + rand.Intn(9) + 1) % 10 // делаем ошибку
	}

	for _, d := range digits {
		dest += strconv.Itoa(d)
	}

	return
}

func calculateChecksum(digits []int) (dest int) {
	sum := 0
	for i := range 11 {
		sum += digits[i] * (i + 1)
	}

	mod := sum % 11
	if mod < 10 {
		return mod
	}

	sum = 0
	weights := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
	for i := 0; i < 11; i++ {
		sum += digits[i] * weights[i]
	}
	
	mod = sum % 11
	if mod == 10 {
		return 0
	}

	return mod
}

func randomName() (dest string) {
	names := []string{"Макс 123", "Че жаксыбек", "Alikhan", "Ашим Жаксылык", "Alzhygan Shal kek", "Сет че там"}

	dest = names[rand.Intn(len(names))]

	return
}
