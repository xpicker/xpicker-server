package app

import (
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	Id           string
	Username     string
	Password     string
	RegisterTime time.Time
	LastTime     time.Time
	Type         string
	Email        string
	Mobile       string
}
