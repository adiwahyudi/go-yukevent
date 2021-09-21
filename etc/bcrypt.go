package etc

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
