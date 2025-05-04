package person

type Entity struct {
	Name  string `json:"-" db:"name"`
	IIN   string `json:"-" db:"iin"`
	Phone string `json:"-" db:"phone"`
}

func ParseFromEntity(data Entity) (dest GetPersonResponse) {
	dest = GetPersonResponse{
		DTO: DTO{
			Name:  data.Name,
			IIN:   data.IIN,
			Phone: data.Phone,
		},
	}

	return
}

func ParseFromEntities(data []Entity) (dest []GetPersonResponse) {
	dest = make([]GetPersonResponse, 0)

	for _, item := range data {
		dest = append(dest, ParseFromEntity(item))
	}

	return
}
