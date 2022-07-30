package main

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

func HtmlTmpl (content string) string{
	tmpl := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>` + content + `</body>
	</html>`
		
	return tmpl
}