package server

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/spf13/viper"
)

func CollectTypes() []string {
	var types []string
	var contexts []Context
	viper.UnmarshalKey("servers", &contexts)
	for _, context := range contexts {
		for _, t := range context.Types {
			if !array.ContainsString(types, t) {
				types = append(types, t)
			}
		}
	}
	sort.Strings(types)
	return types
}

func CollectServices() []string {
	var services []string
	for _, context := range LoadServers() {
		response := ListServices(context)
		if response.Code == 200 {
			for service := range response.Payload {
				if !array.ContainsString(services, service) {
					services = append(services, service)
				}
			}
		}
	}
	sort.Strings(services)
	return services
}

func FilterWithService(selector, service string) []Context {
	return FiltersWithService([]string{selector}, service)
}

func FiltersWithService(selectors []string, service string) []Context {
	var filtered []Context
	for _, context := range LoadServers() {
		for _, selector := range selectors {
			if array.ContainsString(context.Types, selector) {
				response := ListServices(context)
				if response.Code == 200 {
					if _, found := response.Payload[service]; found {
						filtered = append(filtered, context)
					}
				}
			}
		}
	}
	return filtered
}

func IsValidType(t string) bool {
	types := CollectTypes()
	return array.ContainsString(types, t)
}

func IsValidTypeError(t string) error {
	if IsValidType(t) {
		return nil
	}
	list := strings.Join(CollectTypes(), ", ")
	return fmt.Errorf("%q must be one of: %v", "type", list)
}

func Filter(contexts []Context, filters []string) []Context {
	var filtered []Context
	for _, context := range contexts {
		found := 0
		for _, filter := range filters {
			if array.ContainsString(context.Types, filter) {
				found++
			}
		}
		if found == len(filters) {
			filtered = append(filtered, context)
		}
	}
	return filtered
}

func LoadServers() []Context {
	var contexts []Context
	debug := env.GetBool("debug", false)
	insecure := env.GetBool("insecure", true)
	viper.UnmarshalKey("servers", &contexts)
	if len(contexts) == 0 {
		context := Context{
			Debug:    debug,
			Insecure: insecure,
			Endpoint: "127.0.0.1:27482",
			Token:    "AUTH_TOKEN_IS_NOT_CONFIGURED",
			Types:    []string{"localhost"},
		}
		contexts = append(contexts, context)
	}
	for i := range contexts {
		contexts[i].Debug = debug
		contexts[i].Insecure = insecure
	}
	return contexts
}

func ForEach(Runner func(int, int, Context)) int {
	contexts := LoadServers()
	total := len(contexts)
	for index, context := range contexts {
		Runner(index, total, context)
	}
	return total
}

func FilterForEach(filters []string, Runner func(int, int, Context)) int {
	contexts := Filter(LoadServers(), filters)
	total := len(contexts)
	for index, context := range contexts {
		Runner(index, total, context)
	}
	return total
}

func FilterWithServiceForEach(selector, service string, Runner func(int, int, Context)) int {
	return FiltersWithServiceForEach([]string{selector}, service, Runner)
}

func FiltersWithServiceForEach(selectors []string, service string, Runner func(int, int, Context)) int {
	contexts := FiltersWithService(selectors, service)
	total := len(contexts)
	for index, context := range contexts {
		Runner(index, total, context)
	}
	return total
}
