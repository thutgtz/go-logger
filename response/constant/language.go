package constant

import "strings"

type Langauge string

const (
	TH Langauge = "th"
	EN Langauge = "en"
)

func GetLangauge(lang string) Langauge {
	switch strings.ToLower(lang) {
	case string(TH):
		return TH
	case string(EN):
		return EN
	}
	return TH
}
