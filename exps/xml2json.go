package exps

import (
	"fmt"
	"encoding/xml"
	"log"
	"strings"
)

const j = `
<?xml version="1.0" encoding="UTF-8" ?>
<Data>
	<Series>
		<id>83462</id>
		<Actors>|Nathan Fillion|Stana Katic|Molly C. Quinn|Jon Huertas|Seamus Dever|Tamala Jones|Susan Sullivan|Ruben Santiago-Hudson|Monet Mazur|</Actors>
	</Series>
	<Episode>
		<id>398671</id>
		<EpisodeName>Flowers for Your Grave</EpisodeName>
	</Episode>
	<Episode>
		<id>424159</id>
		<EpisodeName>Nanny McDead</EpisodeName>
	</Episode>
</Data>
`

type query struct {
	Series *show
	Episodes []*episode `xml:"Episode"`
}

type show struct {
	Id int `xml:"id"`
	Actors string `xml:"Actors"`
}

type episode struct {
	Id int `xml:"id"`
	Name string `xml:"EpisodeName"`
}

func (s show) String() string {
	return fmt.Sprintf("%s", s.Actors)
}

func (e episode) String() string {
	return fmt.Sprintf("%s", e.Name)
}

func XML2json() {
	r := strings.NewReader(j)

	q := new(query)

	decoder := xml.NewDecoder(r)
	err := decoder.Decode(q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}