package decoupling

import "fmt"

/*
We used the interface to define the behavior of notification
*/
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

/*
notify implements the notifier interface with a pointer receiver
*/
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

/*
sendNotification accepts values that implement the notifier interface and sends notification
*/
func sendNotification(n notifier) {
	n.notify()
}

func NotificationDemo() {
	user := user{"Jordan", "jordan@gmail.com"}
	sendNotification(&user)
}
