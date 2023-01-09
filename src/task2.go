package src

import (
	"fmt"
	"math/rand"
	// "strconv"
	"strings"
	"time"
)


var number = []rune("0123456789")
var special = []rune("!#$%^&@*")


func GenPass(password string, level string) string{
	
	rand.Seed(time.Now().UnixNano())
	result := ""

	upp1 := strings.ToUpper(password[0:1])
	upp2 := strings.ToUpper(password[4:5])

	low := upp1+password+upp2

	spcChar := make([]rune, 2)
		for i := range spcChar {
			spcChar[i] = special[rand.Intn(len(special))]
		}

	numChar := make([]rune, 3)
		for i := range numChar {
		numChar[i] = number[rand.Intn(len(number))]
		}
	
	if len(password) < 6 { 
		fmt.Println("password must be at least 6 characters !")
	} else if level == "low" {
		result += low +" Level : " + level
		} else if level == "mid" {
			result += low + string(numChar) + " Level : " + level
		} else if level == "strong" {
			result += low + string(spcChar) + string(numChar) + " Level : " + level
		} else {
			result += "error, something wrong !"
		}
		
	return result
	

}