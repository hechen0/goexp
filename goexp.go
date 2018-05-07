package main

import (
	"fmt"
	"net/url"
)



func main() {
	s := `{"source":"LINKAPP","platForm":"PC"}`

	u:=url.Values{}
	u.Add("param", s)

	t := u.Encode()

	fmt.Println(t)
}
