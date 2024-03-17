package pkg

import (
	"github.com/go-playground/validator/v10"
)

func ContentsCheck(fl validator.FieldLevel) bool {
	parent := fl.Parent()

	title, ok := parent.FieldByName("Title").Interface().(string)
	if !ok {
		return false
	}
	content, ok := parent.FieldByName("Content").Interface().(string)
	if !ok {
		return false
	}

	// タイトルとコンテンツどちらかが入力されていないとエラー
	return title != "" && content != ""
}
