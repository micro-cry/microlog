package microlog

import (
	"fmt"
	"strings"
	"time"
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

func (emtp *EmbedTemplateObj) NewTemplate() *GlobalDocInfoObj {
	obj := new(GlobalDocInfoObj)
	obj.Params = make(map[string]string)

	obj.embeddedTemplate = emtp
	obj.TemplatePath = emtp.FullPath()
	obj.GenerationTime = time.Now().Format(time.RFC3339)

	obj.Params["ver"] = "'" + GlobalVersion + "'"
	obj.Params["name"] = "'" + GlobalName + "'"
	obj.Params["commit_hash"] = "'" + GlobalHash[32:] + "'"
	obj.Params["commit_date"] = "'" + GlobalDateUpdate + "'"

	return obj
}

//

func (obj *GlobalDocInfoObj) TemplateText() string {
	return obj.embeddedTemplate.Data
}

func (obj *GlobalDocInfoObj) NameGoFile() string {
	return obj.embeddedTemplate.Name + ".go"
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
