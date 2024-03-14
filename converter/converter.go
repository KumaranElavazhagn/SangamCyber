package converter

// The function `EnglishToTamil` converts English characters to their corresponding Tamil characters
// using a predefined map.
func EnglishToTamil(char byte) string {
	// Define a map to directly map English alphabets to their corresponding Tamil characters
	mappings := map[byte]string{
		'A': "அ", 'B': "ஆ", 'C': "இ", 'D': "ஈ", 'E': "உ",
		'F': "ஊ", 'G': "எ", 'H': "ஏ", 'I': "ஐ", 'J': "ஒ",
		'K': "ஓ", 'L': "ஔ", 'M': "ஃ", 'N': "க்", 'O': "ங்",
		'P': "ச்", 'Q': "ஞ்", 'R': "ட்", 'S': "ண்", 'T': "த்",
		'U': "ந்", 'V': "ன்", 'W': "ப்", 'X': "ம்", 'Y': "ய்",
		'Z': "ர்", 'a': "க", 'b': "ங", 'c': "ச", 'd': "ஞ",
		'e': "ட", 'f': "ண", 'g': "த", 'h': "ந", 'i': "ன",
		'j': "ப", 'k': "ம", 'l': "ய", 'm': "ர", 'n': "ற",
		'o': "ல", 'p': "ள", 'q': "ழ", 'r': "வ", 's': "ஷ",
		't': "ஸ", 'u': "ஹ", 'v': "கி", 'w': "ஙி", 'x': "சி",
		'y': "ஞி", 'z': "டி",
	}
	// Return the Tamil character if it exists in the map, otherwise return the input character itself
	if tamilChar, ok := mappings[char]; ok {
		return tamilChar
	}
	return string(char)
}

// The function `TamilToEnglish` converts Tamil characters to their corresponding English alphabets
// using a predefined mapping.
func TamilToEnglish(tamilChar string) string {
	// Define a map to directly map Tamil characters to their corresponding English alphabets

	mappings := map[string]string{
		"அ": "A", "ஆ": "B", "இ": "C", "ஈ": "D", "உ": "E",
		"ஊ": "F", "எ": "G", "ஏ": "H", "ஐ": "I", "ஒ": "J",
		"ஓ": "K", "ஔ": "L", "ஃ": "M", "க்": "N", "ங்": "O",
		"ச்": "P", "ஞ்": "Q", "ட்": "R", "ண்": "S", "த்": "T",
		"ந்": "U", "ன்": "V", "ப்": "W", "ம்": "X", "ய்": "Y",
		"ர்": "Z", "க": "a", "ங": "b", "ச": "c", "ஞ": "d",
		"ட": "e", "ண": "f", "த": "g", "ந": "h", "ன": "i",
		"ப": "j", "ம": "k", "ய": "l", "ர": "m", "ற": "n",
		"ல": "o", "ள": "p", "ழ": "q", "வ": "r", "ஷ": "s",
		"ஸ": "t", "ஹ": "u", "கி": "v", "ஙி": "w", "சி": "x",
		"ஞி": "y", "டி": "z",
	}
	// Return the English alphabet if the Tamil character exists in the map,
	// otherwise return the input Tamil character itself
	if englishChar, ok := mappings[tamilChar]; ok {
		return englishChar
	}
	// If the Tamil character doesn't exist in the map, return the input Tamil character itself
	return tamilChar
}
