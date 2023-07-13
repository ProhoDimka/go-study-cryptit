package encrypt

func Nimbud(str string) string {
	encryptedStr := ""
	for _, value := range str {
		asciiCode := int(value)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
