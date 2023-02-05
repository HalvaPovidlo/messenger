package auth

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const credentialsFile = "credentials.json"

type data struct {
	Credentials map[string]credentials `json:"credentials"`
}

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) Add(login string, creds credentials) error {
	d, err := s.GetAll()
	if err != nil {
		return err
	}

	d.Credentials[login] = creds

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(credentialsFile, bytes, 0644)
}

func (s *storage) Get(login string) (credentials, error) {
	return credentials{}, nil
}

func (s *storage) GetAll() (data, error) {
	file, err := os.Open(credentialsFile)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return data{Credentials: make(map[string]credentials)}, nil
	case err != nil:
		return data{}, err
	}
	defer file.Close()

	var d data
	byteValue, _ := io.ReadAll(file)
	if err = json.Unmarshal(byteValue, &d); err != nil {
		return data{}, err
	}

	return d, nil

}
