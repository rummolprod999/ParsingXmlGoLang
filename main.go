package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Export struct {
	//P        xml.Name `xml:"export1"`
	/*Protocol struct {
		InnerXML string `xml:",innerxml"`
	} `xml:",any"`*/
	Protocol string `xml:",innerxml"`
}
type Attachment struct {
	FileName string `xml:"fileName"`
}
type ProtocolD struct {
	PurchaseNumber string       `xml:"purchaseNumber"`
	ID             string       `xml:"id"`
	Attachments    []Attachment `xml:"attachments>attachment"`
}

var file = "fcsProtocolDeviation_0187300006516001329_11984634.xml"

func TestEmpty(s string) (string, error) {
	if s != "" {
		return s, nil
	}
	return s, fmt.Errorf("empty string")
}
func main() {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(b))
	var prot Export
	var pr ProtocolD
	if err := xml.Unmarshal([]byte(b), &prot); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := TestEmpty(string(prot.Protocol)); err != nil {
		log.Fatalln("zdc")
	}
	prot.Protocol = strings.Replace(prot.Protocol, "ns2:", "", -1)
	prot.Protocol = strings.TrimSpace(prot.Protocol)

	if err := xml.Unmarshal([]byte(prot.Protocol), &pr); err != nil {
		fmt.Println("hren", err)
		return
	}
	/*xmlReader := bytes.NewReader([]byte(prot.Protocol))
	yourPinnacleLineFeed := new(Protocol)
	if err := xml.NewDecoder(xmlReader).Decode(yourPinnacleLineFeed); err != nil {
		return // or log.Panic(err.Error()) if in main
	}*/
	//fmt.Println(pr)

	fmt.Println(pr.PurchaseNumber)
	fmt.Println(pr.ID)
	for _, r := range pr.Attachments {
		fmt.Println(r.FileName)
	}
}
