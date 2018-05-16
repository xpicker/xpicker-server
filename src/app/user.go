package app

import (
	"database/sql"
	"time"
	"config"
	"lib"
)

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
	lib.CheckErr(err)

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

func GetCheckUserTest() (user User) {
	DbConfig := "user=" + config.PostgresqlUser + " password=" + config.PostgresqlPassword + " dbname=" + config.PostgresqlName +
		" sslmode=disable" + " port=" + config.PostgresqlPort
	db, err := sql.Open("postgres", DbConfig)
	lib.CheckErr(err)

	rows, err := db.Query("select * from users where id = '1'")
	lib.CheckErr(err)

	for rows.Next() {
		user = ScanUser(rows)
		return
	}

	db.Close()
	return
}

func CheckUser(username, email, password string) (user User) {
	DbConfig := "user=" + config.PostgresqlUser + " password=" + config.PostgresqlPassword + " dbname=" +
		config.PostgresqlName + " sslmode=disable" + " port=" + config.PostgresqlPort
	db, err := sql.Open("postgres", DbConfig)
	lib.CheckErr(err)

	loginType, loginId := lib.GetLoginUserType(username, email)
	querySql := "select * from users where " + loginType + "=$1 and password=$2"
	rows, err := db.Query(querySql, loginId, password)
	lib.CheckErr(err)

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