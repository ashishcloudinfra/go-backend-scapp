package helpers

// HandleRequest is a simple function to for testing
func HandleRequest(input string) string {
	if input == "test input" {
		return "expected output"
	}
	return "unexpected output"
}
