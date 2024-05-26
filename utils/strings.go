package utils

import (
	"regexp"
	"strings"
)

type Strings struct {
}

func NewStrings() *Strings {
	return &Strings{}
}

func (s *Strings) StripXMLWhitespace(xml string) string {
	return strings.TrimSpace(regexp.MustCompile(`<\s+`).ReplaceAllString(regexp.MustCompile(`>\s+`).ReplaceAllString(xml, ">"), "<"))
}
