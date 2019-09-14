package bluesoul_test

// The the main function // want `duplicate articles: "[T]he the"`
func main() {
	// a a variable // want "duplicate articles: \"[a] a\""
	var x string
	// an an integer value // want "duplicate articles: \"[a]n an\""
	const i = 42

	// it can found duplicate article over a // want "duplicate articles: \"[a] a\""
	// a newline.
	println(x, i)
}
