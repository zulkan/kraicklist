package utils

import "strings"

func TrimLowerString(str string) string {
	str = strings.ToLower(str)
	str = strings.Trim(str, " ")

	return str
}


//remove string which irrelevant (make sure already trimmed, and lowercased)
func ExtractTitle(text string) string {
	if strings.Contains(text, "phone Number removed") {
		return ""
	}

	// if contains for, exclude string after for
	lowerTextArr := strings.Split(text," ")

	result := ""
	for i, lowerTextSplitted := range lowerTextArr {
		if lowerTextSplitted == "for" {
			return result
		}
		if i > 0 {
			result += " "
		}
		result += lowerTextSplitted
	}

	return result
}