package lib

import (
	"log"
)



func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetLoginUserType(usernameStr, emailStr string) (loginType, loginId string) {
	if usernameStr == "" {
		loginType = "email"
		loginId = emailStr
	} else {
		loginType = "username"
		loginId = usernameStr
	}
	return
}