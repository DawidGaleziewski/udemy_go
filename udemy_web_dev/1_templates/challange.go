package tmpl

import "fmt"

// create a templating lang

func TAG(tag string, content string) string{
	return fmt.Sprintf("<%v>%v</%v>", tag, content, tag)
}

func A(content string) string{
	return TAG("a", content)
}

func test(){
	fmt.Println("howdy there")
}