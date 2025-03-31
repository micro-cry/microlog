package microlog

import (
	"fmt"
	"strings"
	"unicode"
)

// // // // // // // // // //

func NameValGo(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	text := strings.Join(s, " ")

	var builder strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		} else {
			builder.WriteRune(' ')
		}
	}
	text = builder.String()

	words := strings.Fields(text)

	if len(words) == 0 {
		return ""
	}

	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			if unicode.IsDigit(runes[0]) {
				words[i] = "Number" + string(runes)
			} else {
				words[i] = string(unicode.ToUpper(runes[0])) + strings.ToLower(string(runes[1:]))
			}
		}
	}

	return strings.Join(words, "")
}

// // // // //

func (b ColumType) Byte() byte {
	return byte(b)
}

func (b ColumType) String() string {
	return ColumMap[b]
}

// //

func (b KeyType) Byte() byte {
	return byte(b)
}

func (b KeyType) String() string {
	return KeyMap[b]
}

// //

func (emtp *EmbedTemplateObj) FileName() string {
	return emtp.Type + "-" + emtp.Name + ".tmpl"
}

func (emtp *EmbedTemplateObj) FullPath() string {
	return emtp.Path + "/" + emtp.FileName()
}

// //

func (column *InfoColumObj) TypeString() string {
	switch column.Type {

	case ColumBool, ColumByte, ColumString:
		return column.Type.String()

	case ColumBytes:
		if column.Length == 0 {
			return "[]byte"
		} else {
			return fmt.Sprintf("[%d]byte", column.Length)
		}

	case ColumDateTime:
		return "time.Time"
	}

	return "any"
}
