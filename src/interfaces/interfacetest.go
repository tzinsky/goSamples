package main

import "fmt"

type sampleInterface interface {
    DoSomething(data string) (int, string)
}

func implenter (i sampleInterface) {
    return 7, "Go you crazy "+data
}

func main() {
    
}
