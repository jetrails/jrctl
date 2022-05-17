package utils

func CollectErrors(errs []error) []string {
	var messages []string
	for _, err := range errs {
		messages = append(messages, err.Error())
	}
	return messages
}
