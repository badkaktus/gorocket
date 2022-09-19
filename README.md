# Golang Rocket Chat REST API client

Use this simple client if you need to connect to Rocket Chat 
in Golang.

## How to use

Just import

```go
import (
	"github.com/badkaktus/gorocket"
)
```

Create client
```go
client := gorocket.NewClient("https://your-rocket-chat.com")

// login as the main admin user
login := gorocket.LoginPayload{
    User:     "admin-login",
    Password: "admin-password",
}

lg, err := client.Login(&login)

if err != nil {
    fmt.Printf("Error: %+v", err)
}
fmt.Printf("I'm %s", lg.Data.Me.Username)
```

## Manage user
```go
str := gorocket.NewUser{
    Email:                 "test@email.com",
    Name:                  "John Doe",
    Password:              "simplepassword",
    Username:              "johndoe",
    Active:                true,
}

me, err := client.UsersCreate(&str)
if err != nil {
    fmt.Printf("Error: %+v", err)
}
fmt.Printf("User was created %t", me.Success)
```

## Post a message
```go
// create a new channel
str := gorocket.CreateChannelRequest{
    Name:     "newchannel",
}

channel, err := client.CreateChannel(&str)
if err != nil {
    fmt.Printf("Error: %+v", err)
}
fmt.Printf("Channel was created %t", channel.Success)
// post a message
str := gorocket.Message{
    Channel:     "somechannel",
    Text:        "Hey! This is new message from Golang REST Client",
}

msg, err := client.PostMessage(&str)
if err != nil {
    fmt.Printf("Error: %+v", err)
}
fmt.Printf("Message was posted %t", msg.Success)
```
## PS
Feel free to create issue for add new endpoint to this client