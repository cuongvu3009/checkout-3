// https://github.com/01-edu/public/tree/master/subjects/capitalize

package piscine

func Capitalize(s string) string {

	//alphanumeric check function
	//islowercase check function
	//Capitalizer function

	slicy := []rune(s)

	for i := 0; i < len(slicy); i++ {

		if i == 0 || !alphanumeric(slicy[i-1]) && alphanumeric(slicy[i]) {
			slicy[i] = Capitalizer(slicy[i])
		}
	}
	return string(slicy)
}

func alphanumeric(each rune) bool {

	return ('a' <= each && each <= 'z' || 'A' <= each && each <= 'Z' || '0' <= each && each <= '9')

}

func islower(input rune) bool {
	return 'a' <= input && input <= 'z'
}

func Capitalizer(meow rune) rune {
	if islower(meow) {
		meow = meow - ('a' - 'A')
	}

	return meow
}
