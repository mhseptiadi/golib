### import
How to import the library
```
import (
	"github.com/mhseptiadi/golib/checkdocker"
	"github.com/mhseptiadi/golib/sendgmail"
)
```

### sendgmail
Golang library used for sending email using gmail provider

```
recipients := "target1@example.com,target2@example.com"
subject := "Test Email Subject"
body := "Test Email Body"
sender := "sender@yourgmaildomain.com"
password := "senderpassword"
sendgmail.Send(recipients, subject, body, sender, password)
```

### checkdocker
Golang library used for checking whether the list docker is running on the server

```
dockerNames := []string{"dockerimage1", "dockerimage2"}

err := checkdocker.Check(dockerNames)
if err != nil {
	fmt.Println(err)
}
```