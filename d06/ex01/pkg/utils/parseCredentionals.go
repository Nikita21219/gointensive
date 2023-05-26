package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Credentionals struct {
	PgUser   string
	PgDB     string
	Login    string
	Password string
}

func ParseFile(filename string) (*Credentionals, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	cred := &Credentionals{}
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "=")
		if len(splitLine) != 2 {
			return nil, fmt.Errorf("Wrong format access file")
		}

		switch splitLine[0] {
		case "POSTGRES_USER":
			cred.PgUser = splitLine[1]
		case "POSTGRES_DB":
			cred.PgDB = splitLine[1]
		case "LOGIN":
			cred.Login = splitLine[1]
		case "PASSWORD":
			cred.Password = splitLine[1]
		}
	}

	if cred.PgUser == "" || cred.PgDB == "" || cred.Login == "" || cred.Password == "" {
		return nil, fmt.Errorf("Wrong format access file")
	}

	return cred, nil
}
