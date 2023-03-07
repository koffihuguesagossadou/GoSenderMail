package util

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func EmailSenderFuncs() {
	auth  := smtp.PlainAuth(
		"",
		"koffi.agossadou@gmail.com",
		"hcbbpgigkvmatyhq",
		"smtp.gmail.com",
	)
	
	to := "hugocomptespam@gmail.com"
	me := "koffi.agossadou@gmail.com"

	message := []byte("To: "+to+"\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail(
		"smtp.gmail.com:587", 
		auth, 
		me,
		[]string{to},
		message,
	)

	if err != nil {
		fmt.Println(err)
	}
}

/*

*/
func EmailSenderFuncsWithTemp(path string, receiver string) {

	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	
	mTemplate, err := template.ParseFiles(path)

	if err != nil{
		fmt.Println(err)
	}

	err = mTemplate.Execute(&body, struct{ Name string }{ Name: "Hugo" })

	if err != nil{
		fmt.Println(err)
	}

	auth  := smtp.PlainAuth(
		"",
		"koffi.agossadou@gmail.com",
		"hcbbpgigkvmatyhq",
		"smtp.gmail.com",
	)
	
	me := "koffi.agossadou@gmail.com"

	message := []byte("To: "+receiver+"\n" +
		"Subject: discount Gophers!" +
		"\n" +
		headers +
		"\n\n"+
		body.String()+"\r\n")

	err = smtp.SendMail(
		"smtp.gmail.com:587", 
		auth, 
		me,
		[]string{receiver},
		message,
	)

	if err != nil {
		fmt.Println(err)
	}
}