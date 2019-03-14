### sendgmail

```
recipients := "target1@example.com,target2@example.com"
subject := "Test Email Subject"
body := "Test Email Body"
sender := "sender@yourgmaildomain.com"
password := "senderpassword"
sendgmail.Send(recipients, subject, body, sender, password)
```

### checkdocker

```
dockerNames := []string{"dockerimage1", "dockerimage2"}

err := checkdocker.Check(dockerNames)
if err != nil {
	fmt.Println(err)
}
```