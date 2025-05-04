package person

type (
	ValidateIINResponse struct {
		Correct     bool   `json:"correct"`
		Sex         string `json:"sex,omitempty"`
		DateOfBirth string `json:"date_of_birth,omitempty"`
	}

	DTO struct {
		Name  string `json:"name"`
		IIN   string `json:"iin"`
		Phone string `json:"phone"`
	}

	CreatePersonRequest struct {
		DTO
	}

	GetPersonResponse struct {
		DTO
	}
)

func ParseToEntity(req CreatePersonRequest) (dest Entity) {
	dest = Entity{
		Name:  req.Name,
		IIN:   req.IIN,
		Phone: req.Phone,
	}

	return
}
