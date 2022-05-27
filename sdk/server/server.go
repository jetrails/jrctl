package server

import (
	"fmt"
	"strings"

	"crypto/sha256"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/spf13/viper"
)

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

func ForEach(Runner func(int, int, Context)) int {
	contexts := LoadContexts()
	total := len(contexts)
	for index, context := range contexts {
		Runner(index, total, context)
	}
	return total
}

func FilterForEach(filters []string, Runner func(int, int, Context)) int {
	contexts := Filter(LoadContexts(), filters)
	total := len(contexts)
	for index, context := range contexts {
		Runner(index, total, context)
	}
	return total
}

func (c Context) Hash() string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", c)))
	return fmt.Sprintf("%x", hash.Sum(nil))[0:8]
}
