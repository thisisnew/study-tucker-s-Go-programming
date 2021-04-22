package main

import (
	"study-tuckers-go-programming/ex20/fedex"
	"study-tuckers-go-programming/ex20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaSender := &koreaPost.PostSender{}
	SendBook("어린왕자", koreaSender)
	SendBook("그리스인 조르바", koreaSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}
