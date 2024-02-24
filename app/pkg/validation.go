package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ContentCheck(fl validator.FieldLevel) bool {
	fmt.Println("@@@@@@チェックイン＠＠＠＠＠＠＠")
	parent := fl.Parent()

	title, ok := parent.FieldByName("Title").Interface().(string)
	fmt.Println(ok)
	content, ok := parent.FieldByName("Content").Interface().(string)
	fmt.Println(ok)
	fmt.Println(title != "" || content != "")
	return true
}
