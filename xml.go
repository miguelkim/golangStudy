package main

/*
import (
	"encoding/xml"
	"io/ioutil"
)

type Campus struct {
	XMLName  xml.Name `xml:"campus"`
	Name     string   `xml:"name,attr"`
	Comment  string   `xml:",comment"`
	Building Building
}
type Building struct {
	XMLNAME xml.Name `xml:"building"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
	Device  Device
}
type Device struct {
	XMLNAME xml.Name `xml:"device"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

func main() {
	camp := Campus{
		Name:    "campus1",
		Comment: "building-c",
		Building: Building{
			Name:    "big",
			Comment: "Dev",
			Device: Device{
				Name: "1",
				Type: "r",
			},
		},
	}
	indt, _ := xml.MarshalIndent(camp, "", " ")
	_ = ioutil.WriteFile("file.xml", indt, 0644)
}
*/
