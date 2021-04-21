package test_go_mod

import (
    "fmt"
    "github.com/jehiah/go-strftime"
    "time"
)

func Print2222() {
    fmt.Println("222222")
    strftime.Format("", time.Now())
}
