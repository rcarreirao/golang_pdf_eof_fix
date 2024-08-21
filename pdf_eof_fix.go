package pdf_eof_fix

import (
	"fmt"
	"os"
	"regexp"
)

func FixEofOnPdfFile(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	position := locateEofMark(content)
	newContent := retrieveNewContent(content, position)
	os.WriteFile(file, newContent, 0777)
}

func locateEofMark(content []byte) int {

	var regex string = `(?i)(%%EOF\n)+(\^@|\w.+?\r\n|\n)*`
	var size int = len(content)
	headerRegex, regexError := regexp.Compile(regex)
	if headerRegex != nil {
		matches := headerRegex.FindSubmatchIndex(content)
		fmt.Printf("[%d] \n", matches)
		if len(matches) > 1 {
			size = matches[1]
		}
	} else {
		fmt.Printf(fmt.Sprintf("\033[0;31m%s\033[0m", regexError))
	}
	return size
}

func retrieveNewContent(content []byte, position int) []byte {
	c := make([]byte, position)
	copy(c, content)
	return c
}
