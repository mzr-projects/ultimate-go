package decoupling

import "fmt"

type notifierWithEmbedding interface {
	notifyWithEmbedding()
}

type embeddingUser struct {
	name  string
	email string
}

/*
Admin represents an admin user with privileges
Here admin is also implementing the notifyWithEmbedding interface thanks to the embedding feature in GO
*/
type admin struct {
	embeddingUser
	level string
}

func (v *embeddingUser) notifyWithEmbedding() {
	fmt.Printf("(embeddingUser) Sending email to: %s<%s>\n", v.name, v.email)
}

func (v *admin) notifyWithEmbedding() {
	fmt.Printf("(Admin) Sending email to: %s<%s>\n", v.name, v.email)
}

func EmbeddingDemo() {
	ad := admin{
		embeddingUser: embeddingUser{
			name:  "ad",
			email: "ad@example.com",
		},
		level: "super",
	}

	fmt.Printf("Hello, %s<%s>\n", ad.name, ad.email)

	/*
		We can access the inner type's method directly
	*/
	ad.embeddingUser.notifyWithEmbedding()

	/*
		The inner type's method is promoted Because the ad is using embedding embeddingUser it is promoting the embeddingUser
		notify() method so it can call it.
	*/
	ad.notifyWithEmbedding()

	/*
		Send the admin user a notification.
		Embedded inner type's implementation of the interface is "promoted" to the outer type
	*/
	SendNotification(&ad)
}

func SendNotification(n notifierWithEmbedding) {
	n.notifyWithEmbedding()
}
