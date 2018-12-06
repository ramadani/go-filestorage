# go-filestorage

Golang package for file storage

## Requirements

- Go >= 1.10

## Installation

```bash
go get github.com/ramadani/go-filestorage
```

## Usage

```go
import storage "github.com/ramadani/go-filestorage"

config := &storage.Config{
	Root: "storage",
	URL:  "http://yourdomain.com/public",
}

localStorage := storage.NewStorage(config)
```

**Store the file to the storage**

```go
localStorage.PutFile("user", file)
```

It will be create a new `user` dir if not exists and then store the file into that.

**Store the file to the storage with specifying a file name**

```go
localStorage.PutFileAs("user", file, "user1.jpeg")
```

## **License**
The **go-filestorage** is an open-source software licensed under the [MIT License](LICENSE.md).