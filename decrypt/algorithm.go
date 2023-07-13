package decrypt

func Nimbud(str string) string {
	decryptedStr := ""
	for _, value := range str {
		asciiCode := int(value)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
