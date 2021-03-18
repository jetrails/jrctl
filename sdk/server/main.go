package server

import (
	"fmt"
	"sort"
	"errors"
	"strings"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/utils"
)

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
			if !utils.Includes ( t, types ) {
				types = append ( types, t )
			}
		}
	}
	sort.Strings ( types )
	return types
}

func CollectServices () [] string {
	var services [] string
	for _, context := range LoadServers () {
		response := ListServices ( context )
		if response.Code == 200 {
			for _, service := range response.Payload {
				if !utils.Includes ( service, services ) {
					services = append ( services, service )
				}
			}
		}
	}
	sort.Strings ( services )
	return services
}

func FilterWithService ( selector, service string ) [] Context {
	var filtered [] Context
	for _, context := range LoadServers () {
		if utils.Includes ( selector, context.Types ) {
			response := ListServices ( context )
			if response.Code == 200 && utils.Includes ( service, response.Payload ) {
				filtered = append ( filtered, context )
			}
		}
	}
	return filtered
}

func IsValidType ( t string ) bool {
	types := CollectTypes ()
	return utils.Includes ( t, types )
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
			if utils.Includes ( filter, context.Types ) {
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
			Endpoint: "127.0.0.1:27482",
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
