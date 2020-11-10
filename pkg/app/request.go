package app

import (
	"fmt"
	log "github.com/MikelPan/go-learning/pkg/logging"
	"github.com/astaxie/beego/validation"

	"reflect"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	fmt.Println(errors)
	for _, err := range errors {
		fmt.Println(err.Message)
		fmt.Println(err.Key)
		fmt.Println(reflect.TypeOf(err.Message))
		//fmt.Println(reflect.TypeOf(err))

	}

	return
}