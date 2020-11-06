package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"

	"reflect"

	"github.com/MikelPan/go-learning/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	fmt.Println(reflect.TypeOf(errors))
	for _, err := range errors {
		fmt.Println(err.Message)
		fmt.Println(err.Key)
		fmt.Println(reflect.TypeOf(err.Message))
		fmt.Println(reflect.TypeOf(err))
		logging.Info( err.Key, err.Message)
	}

	return
}