package vo

type Email struct {
	Email string `json:"email"`
}

func (Email) Create(email string) (Email, error) {
	if email == "" {
		return Email{}, nil
	}
	if len(email) > 255 {
		return Email{}, nil
	}
	return Email{Email: email}, nil
}
