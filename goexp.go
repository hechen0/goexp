package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	x := `
<?xml version="1.0" encoding="UTF-8"?>
<Person foo="bar" something="other" else="casd" oo="aa">
    <Id>123</Id>
    <Phone where="home">8172937</Phone>
    <Phone where="work">1233121</Phone>
    <s>
    	<d>
    		<e>some thing ddep</e>
    	</d>
    </s>
</Person>
	`

	var s struct {
		Id  int `xml:"Id"`
		Foo xml.Attr `xml:"foo,attr"`
		Phone []struct {
			Where  string `xml:"where,attr"`
			Number string `xml:",chardata"`
		}

		Hello string `xml:",any,attr"`
		Hello1 string `xml:",any,attr"`
		Hello2 string `xml:",any,attr"`

		Deep string `xml:"s>d>e"`
	}

	if err := xml.Unmarshal([]byte(x), &s); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v\n", s)

		fmt.Println(s.Foo.Value)
	}
}
