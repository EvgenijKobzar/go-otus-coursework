package hash

import "golang.org/x/crypto/bcrypt"

func Make(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}

func Check(input, src string) bool {
	// CompareHashAndPassword сравнивает пароль с хешем
	err := bcrypt.CompareHashAndPassword([]byte(src), []byte(input))
	return err == nil
}
