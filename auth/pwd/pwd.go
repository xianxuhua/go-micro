package pwd

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

type MD5PasswordManager struct {
	SaltLen    int
	Iterations int
	KeyLen     int
}

func (m *MD5PasswordManager) Encode(rawPassword string) string {
	options := &password.Options{m.SaltLen, m.Iterations, m.KeyLen, sha512.New}
	salt, encodedPwd := password.Encode(rawPassword, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return newPassword
}

func (m *MD5PasswordManager) Verify(rawPassword, encodedPassword string) bool {
	options := &password.Options{m.SaltLen, m.Iterations, m.KeyLen, sha512.New}
	passwordInfo := strings.Split(encodedPassword, "$")
	return password.Verify(rawPassword, passwordInfo[2], passwordInfo[3], options)
}
