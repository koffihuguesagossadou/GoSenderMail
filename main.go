package main

import(
	"github.com/koffihugues/GoMailSender/util"
)

func main() {

	util.EmailSenderFuncsWithTemp("./templates/email.html", "hugocomptespam@gmail.com")
}