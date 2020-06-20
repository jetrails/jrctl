package utils

import "os"
import "fmt"
import "github.com/spf13/viper"
import "github.com/parnurzeal/gorequest"

func Request () * gorequest.SuperAgent {
	var request = gorequest.New ()
	var debug = viper.GetBool ("debug")
	request.SetDebug ( debug )
	return request
}

type ErrorResponse struct {
	Name string
	Message string
	Code int
	Type string
	Data string
}

func GetPublicEndpoint ( route string ) string {
	var postfix = viper.GetString ("endpoint_postfix")
	return fmt.Sprintf ( "https://api-public%s.jetrails.cloud/%s", postfix, route )
}

func InternalError ( e interface {} ) ErrorResponse {
	return ErrorResponse {
		Name: "InternalError",
		Message: fmt.Sprintf ( "%s", e ),
		Code: 1,
		Type: "INTERNAL_ERROR",
		Data: "",
	}
}

func HandleErrorResponse ( error ErrorResponse ) {
	if error.Code != 200 && error.Code != 0 {
		if error.Data == "" {
			fmt.Printf ( "%s: %s\n", error.Name, error.Message )
		} else {
			fmt.Printf ( "%s: %s\n", error.Name, error.Data )
		}
		os.Exit ( 1 )
	}
}
