package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"github.com/garyburd/redigo/redis"
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

var RedisClient *redis.Pool

func init() {
	RedisClient = &redis.Pool{
		MaxIdle:     RedisMaxIdle,
		MaxActive:   RedisMaxActive,
		IdleTimeout: RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisAddr)
			CheckErr(err)
			c.Do("SELECT", RedisName)
			return c, nil
		},
	}
}

func ScanUser(rows *sql.Rows) (user User) {
	var (
		id           string
		username     string
		userType     string
		email        string
		mobile       string
		password     string
		registerTime time.Time
		lastTime     time.Time
	)

	err := rows.Scan(&id, &password, &username, &registerTime, &lastTime, &userType, &email, &mobile)
	CheckErr(err)

	user = User{
		Id:           id,
		Password:     password,
		Username:     username,
		RegisterTime: registerTime,
		LastTime:     lastTime,
		Type:         userType,
		Email:        email,
		Mobile:       mobile,
	}

	return user
}

func CheckUserTest() (user User) {
	DbConfig := "user=" + PostgresqlUser + " password=" + PostgresqlPassword + " dbname=" + PostgresqlName +
		" sslmode=disable" + " port=" + PostgresqlPort
	db, err := sql.Open("postgres", DbConfig)
	CheckErr(err)

	rows, err := db.Query("select * from users where id = '1'")
	CheckErr(err)

	for rows.Next() {
		user = ScanUser(rows)
		return user
	}

	db.Close()
	return user
}

func CheckUser(username, email, password string) (user User) {
	DbConfig := "user=" + PostgresqlUser + " password=" + PostgresqlPassword + " dbname=" + PostgresqlName +
		" sslmode=disable" + " port=" + PostgresqlPort
	db, err := sql.Open("postgres", DbConfig)
	CheckErr(err)
	var (
		id     string
		idType string
	)
	if username == "" {
		idType = "email"
		id = email
	} else {
		idType = "username"
		id = username
	}
	querySql := "select * from users where " + idType + "=$1 and password=$2"
	rows, err := db.Query(querySql, id, password)
	CheckErr(err)

	for rows.Next() {
		user = ScanUser(rows)
		break
	}

	var tmp User
	if user != tmp {
		db.Exec("update users set last_time=$1 where username=$2", time.Now(), user.Username)
	}
	db.Close()
	return user
}

func RedisSet(key, value, exKey, exValue string) {
	redisClient := RedisClient.Get()
	redisClient.Do("SET", key, value, exKey, exValue)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
