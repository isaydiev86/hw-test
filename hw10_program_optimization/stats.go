package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/buger/jsonparser"
)

var ErrInvalidEmail = errors.New("email does not contain @")

type User struct {
	ID       int64
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	res := make(DomainStat)
	for scanner.Scan() {
		user, err := getUser(scanner.Bytes())
		if err != nil {
			return nil, err
		}

		if !strings.Contains(user.Email, "@") {
			return nil, ErrInvalidEmail
		}

		if strings.HasSuffix(user.Email, domain) {
			tail := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			res[tail]++
		}
	}

	return res, nil
}

func getUser(line []byte) (*User, error) {
	var err error
	var user User

	user.Email, err = jsonparser.GetString(line, "Email")
	if err != nil {
		return nil, err
	}

	return &user, nil
}
