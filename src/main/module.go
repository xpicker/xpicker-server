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
		id            string
		username      string
		_type         string
		email         string
		mobile        string
		password      string
		register_time time.Time
		last_time     time.Time
	)

	err := rows.Scan(&id, &password, &username, &register_time, &last_time, &_type, &email, &mobile)
	CheckErr(err)

	user = User{
		Id:           id,
		Password:     password,
		Username:     username,
		RegisterTime: register_time,
		LastTime:     last_time,
		Type:         _type,
		Email:        email,
		Mobile:       mobile,
	}

	return user
}

func CheckUserTest() (user User) {
	db_config := "user=" + PostgresqlUser + " password=" + PostgresqlPassword + " dbname=" + PostgresqlName +
		" sslmode=disable" + " port=" + PostgresqlPort
	db, err := sql.Open("postgres", db_config)
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
	db_config := "user=" + PostgresqlUser + " password=" + PostgresqlPassword + " dbname=" + PostgresqlName +
		" sslmode=disable" + " port=" + PostgresqlPort
	db, err := sql.Open("postgres", db_config)
	CheckErr(err)
	var (
		id      string
		id_type string
	)
	if username == "" {
		id_type = "email"
		id = email
	} else {
		id_type = "username"
		id = username
	}
	sql := "select * from users where " + id_type + "=$1 and password=$2"
	rows, err := db.Query(sql, id, password)
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

func RedisSet(key, value, exkey, exvalue string) {
	redis := RedisClient.Get()
	redis.Do("SET", key, value, exkey, exvalue)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
