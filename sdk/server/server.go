package server

import (
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/spf13/viper"
)

func ContextsHaveSameToken(contexts []Context) bool {
	seen := ""
	for _, context := range contexts {
		if seen == "" {
			seen = context.Token
		} else {
			if context.Token != seen {
				return false
			}
		}
	}
	return true
}

func ContextsHaveSomeEndpoint(contexts []Context, endpoints []string) bool {
	for _, context := range contexts {
		if array.ContainsString(endpoints, context.Endpoint) {
			return true
		}
	}
	return false
}

func LoadContexts() []Context {
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

func GetContexts(filters []string) []Context {
	results := []Context{}
	contexts := LoadContexts()
	if len(filters) == 0 {
		return contexts
	}
	for _, context := range contexts {
		for _, filter := range filters {
			qualifies := true
			targets := strings.Split(filter, ",")
			for _, target := range targets {
				target = strings.TrimSpace(target)
				qualifies = qualifies && (target == "" || array.ContainsString(context.Types, target))
			}
			if qualifies {
				results = append(results, context)
				break
			}
		}
	}
	return results
}
