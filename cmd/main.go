package main

import (
	"fmt"
	"strings"

	"com.jd/easyopenid"

	uuid "github.com/satori/go.uuid"
)

func main() {

	aUuid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	fmt.Println(aUuid)

	openId, err := easyopenid.GetOpenId("7597ad358eb541ec8044ece9fc281743", "6256aebdf58e4b028ed31d51e1f8741b")
	if err == nil {
		fmt.Println(openId)
	}
	userId, err := easyopenid.GetUserId("7597ad358eb541ec8044ece9fc281743", "d7edkreinjjj8ceegeh7fpjatdhg8b5e")
	if err == nil {
		fmt.Println(userId)
	}
}
