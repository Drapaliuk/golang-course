package users

import (
	"errors"
	ds "lesson_7/documentstore"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll ds.Collection
}

func NewUser(id string, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}

func NewUserCollection() ds.Collection {
	return ds.NewCollection(ds.CollectionConfig{
		PrimaryKey: "ID",
		Name:       "users",
	})
}

func (s *Service) CreateUser() (*User, error) {
	user := map[string]ds.DocumentField{
		"ID":   {Type: ds.DocumentFieldTypeString, Value: "10"},
		"Name": {Type: ds.DocumentFieldTypeString, Value: "Jack"},
	}

	err := s.coll.Put(ds.NewDocument(user))

	if err != nil {
		return nil, nil
	}

	return nil, nil
}

func (s *Service) ListUsers() ([]User, error) {

}

func (s *Service) GetUser(userId string) (*User, error) {

}

func (s *Service) DeleteUser(userID string) error {

}
