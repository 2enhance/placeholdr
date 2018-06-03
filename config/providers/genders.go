package providers

var genders = []string{"male", "female"}

func Gender(seed int) string {
	return genders[seed%2]
}
