package server

import (
	"fmt"
	"errors"
	"strings"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/spf13/viper"
)

func includes ( a string, list [] string ) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func LoadServerAuth () string {
	var c ServerConfig
	if yamlfile, error := ioutil.ReadFile ("/etc/jetrailsd/config.yaml"); error == nil {
		if error = yaml.Unmarshal ( yamlfile, &c ); error == nil {
			return c.Auth
		}
	}
	return ""
}

func CollectTypes () [] string {
	var types [] string
	var contexts [] Context
	viper.UnmarshalKey ( "servers", &contexts )
	for _, context := range contexts {
		for _, t := range context.Types {
			if !includes ( t, types ) {
				types = append ( types, t )
			}
		}
	}
	return types
}

func FilterWithService ( selector, service string ) [] Context {
	var filtered [] Context
	for _, context := range LoadServers () {
		if includes ( selector, context.Types ) {
			response := ListServices ( context )
			if response.Code == 200 && includes ( service, response.Payload ) {
				filtered = append ( filtered, context )
			}
		}
	}
	return filtered
}

func IsValidType ( t string ) bool {
	types := CollectTypes ()
	return includes ( t, types )
}

func IsValidTypeError ( t string ) error {
	if IsValidType ( t ) {
		return nil
	}
	list := strings.Join ( CollectTypes (), ", " )
	return errors.New ( fmt.Sprintf ( "%q must be one of: %v", "type", list ) )
}

func Filter ( contexts [] Context, filters [] string ) [] Context {
	var filtered [] Context
	for _, context := range contexts {
		found := 0
		for _, filter := range filters {
			if includes ( filter, context.Types ) {
				found++
			}
		}
		if found == len ( filters ) {
			filtered = append ( filtered, context )
		}
	}
	return filtered
}

func LoadServers () [] Context {
	var contexts [] Context
	debug := viper.GetBool ("debug")
	viper.UnmarshalKey ( "servers", &contexts )
	if len ( contexts ) == 0 {
		context := Context {
			Debug: debug,
			Endpoint: "localhost:27482",
			Token: LoadServerAuth (),
			Types: [] string { "localhost" },
		}
		contexts = append ( contexts, context )
	}
	for i, _ := range contexts {
		contexts [ i ].Debug = debug
	}
	return contexts
}

func ForEach ( Runner func ( int, int, Context ) ) int {
	contexts := LoadServers ()
	total := len ( contexts )
	for index, context := range contexts {
		Runner ( index, total, context )
	}
	return total
}

func FilterForEach ( filters [] string, Runner func ( int, int, Context ) ) int {
	contexts := Filter ( LoadServers (), filters )
	total := len ( contexts )
	for index, context := range contexts {
		Runner ( index, total, context )
	}
	return total
}

func FilterWithServiceForEach ( selector, service string, Runner func ( int, int, Context ) ) int {
	contexts := FilterWithService ( selector, service )
	total := len ( contexts )
	for index, context := range contexts {
		Runner ( index, total, context )
	}
	return total
}
