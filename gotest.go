package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	zhihu := "http://www.zhihu.com/collection/19697157?page="

	i := 1
	for true {
		resp, err := http.Get(zhihu + strconv.Itoa(i))
		i++
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		re, _ := regexp.Compile("/question/\\d*/answer/\\d*")
		link := re.FindAllString(string(body), -1)
		if link == nil {
			break
		}
		for _, l := range link {
			fmt.Println("http://www.zhihu.com"+l)
//			mail("http://www.zhihu.com"+l)
		}
	}
}

func mail(msg string) {
	auth := smtp.PlainAuth("", "mail@mail", "pwd", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "From", []string{"add@getpocket.com"},
		[]byte("Subject:Pocket\nTo:Pocket add@getpocket.com\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+"<a href=\""+msg+"\">"))
	if err != nil {
		fmt.Println("send mail failed", err)
	}
}
