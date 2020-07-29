package middleware

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"regexp"
)

func Markdown(content string) string {
	input := []byte(content)
	unsafe := blackfriday.MarkdownCommon(input)
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	html := p.SanitizeBytes(unsafe)
	return string(html)
}
