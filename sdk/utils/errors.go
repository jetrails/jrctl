package utils

func CollectErrors(errors []error) []string {
	var messages []string
	for _, error := range errors {
		messages = append(messages, error.Error())
	}
	return messages
}
