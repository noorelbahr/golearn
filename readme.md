# My First Go #
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. More info: https://golang.org

## Getting Started
It's easy to get started.

### Set Up GO Environment
Install Go and set `GOPATH` variable with simply follow the instruction from the video [here](https://www.youtube.com/watch?v=5qI8z_lB5Lw).

### Set Up Project

Make sure we have a database named `golearn`


Move the project to your workspace (i.e. `golang/github.com/noorelbahr`).

Then open terminal and move to the project directory
```
cd $GOPATH/github.com/{username}/golearn
```

Here we assume that our `Apache` and `MySQL` server are running.

---
### Run Our Project
Simply run :
```
make dev
```

Or
```
go run main.go
```

It will run initial migration for `users` table and add one default user, use the credential below to hit our login endpoint:
```
fullname: John Doe
username: johndoe
password: 123123
```

## Testing Our API
To test our API, click button bellow : 

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4ff03dae8d6678c1a248)

