package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTData struct {
	Phone string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(phone string) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Phone": phone,
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {

		return "", nil
	}

	return s, nil

}

func (j *JWT) Pasre(token string) (bool, *JWTData) {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := t.Claims.(jwt.MapClaims)["phone"]
	return t.Valid, &JWTData{
		Phone: phone.(string),
	}
}
