package main

import "fmt"

// type message struct {
// 	msg      []byte
// 	id       string
// 	typ      string
// 	sender   string
// 	receiver string
// 	time     float32
// }
type message struct {
	Name, Text string
}

func (m message) log() (int, error) {
	fmt.Println(m)
	return 1, nil
}
