package user

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	if !strings.Contains(address, "@") {
		return Email{}, errors.New("invalid email address")
	}

	return Email{address: address}, nil
}

func (e Email) String() string {
	return e.address
}

func (e *Email) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Email")
	}
	e.address = str
	return nil
}

func (e Email) Value() (driver.Value, error) {
	return e.address, nil
}

type Password struct {
	hash string
}

func (p *Password) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Password")
	}
	p.hash = str
	return nil
}

func (p Password) Value() (driver.Value, error) {
	return p.hash, nil
}

func NewPassword(hash string) Password {
	return Password{hash: hash}
}

func (p Password) Hash() string {
	return p.hash
}

type Skills struct {
	items []string
}

func NewSkills(items []string) Skills {
	return Skills{items: items}
}

func (s *Skills) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Skills")
	}
	s.items = strings.Split(str, ", ")
	return nil
}

func (s Skills) Value() (driver.Value, error) {
	if len(s.items) == 0 {
		return "", nil
	}
	return strings.Join(s.items, ", "), nil
}

func (s Skills) Items() []string {
	return s.items
}

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (r *Role) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Role")
	}
	*r = Role(str)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
