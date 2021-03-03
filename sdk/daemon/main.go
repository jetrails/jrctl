package daemon

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

func LoadDaemonAuth () string {
	var c DaemonConfig
	if yamlfile, error := ioutil.ReadFile ("/etc/jetrailsd/config.yaml"); error == nil {
		if error = yaml.Unmarshal ( yamlfile, &c ); error == nil {
			return c.Auth
		}
	}
	return ""
}

func CollectServices () [] string {
	var services [] string
	var contexts [] Context
	viper.UnmarshalKey ( "daemons", &contexts )
	for _, context := range contexts {
		for _, service := range context.Services {
			if !includes ( service, services ) {
				services = append ( services, service )
			}
		}
	}
	return services
}

func IsValidService ( service string ) bool {
	services := CollectServices ()
	return includes ( service, services )
}

func IsValidServiceError ( service string ) error {
	if IsValidService ( service ) {
		return nil
	}
	list := strings.Join ( CollectServices (), ", " )
	return errors.New ( fmt.Sprintf ( "%q must be one of: %v", "service", list ) )
}

func Filter ( contexts [] Context, filters [] string ) [] Context {
	var filtered [] Context
	for _, context := range contexts {
		found := 0
		for _, filter := range filters {
			if includes ( filter, context.Services ) {
				found++
			}
		}
		if found == len ( filters ) {
			filtered = append ( filtered, context )
		}
	}
	return filtered
}

func LoadDaemons () [] Context {
	var contexts [] Context
	debug := viper.GetBool ("debug")
	viper.UnmarshalKey ( "daemons", &contexts )
	if len ( contexts ) == 0 {
		context := Context {
			Debug: debug,
			Endpoint: "localhost:27482",
			Token: LoadDaemonAuth (),
			Services: [] string { "apache", "nginx", "mysql", "varnish" },
		}
		contexts = append ( contexts, context )
	}
	for _, context := range contexts {
		context.Debug = debug
	}
	return contexts
}

func ForEach ( Runner func ( int, int, Context ) ) int {
	contexts := LoadDaemons ()
	total := len ( contexts )
	for index, context := range contexts {
		Runner ( index, total, context )
	}
	return total
}

func FilterForEach ( filters [] string, Runner func ( int, int, Context ) ) int {
	contexts := Filter ( LoadDaemons (), filters )
	total := len ( contexts )
	for index, context := range contexts {
		Runner ( index, total, context )
	}
	return total
}
