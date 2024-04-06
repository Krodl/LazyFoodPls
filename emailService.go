package main

import (
	"fmt"

	"github.com/go-mail/mail"
)

func SendEmail(recipes []Recipe) {

	m := mail.NewMessage()

	m.SetHeader("From", "GibDaFood@gmail.com")

	m.SetHeader("To", "kordell.teenie@gmail.com")

	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", getBody(recipes))

	d := mail.NewDialer("smtp.gmail.com", 587, "GibDaFood@gmail.com", "hsbw tesy idgp ehla ")

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}
	fmt.Println("Email Sent Successfully!")
}

func getBody(recipes []Recipe) string {
	var b string
	for _, r := range recipes {
		htmlRecipe := fmt.Sprintf("<div><h3><a href=%s>%s</a></h3><div>%s</div></div>\n", r.url, r.title, r.ingredients)
		b = fmt.Sprintf("%s%s", b, htmlRecipe)
	}
	return b
}
