package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type ListData struct {
	XMLName xml.Name    `xml:"list"`
	List    []UsersList `xml:"users"`
}

type UsersList struct {
	XMLName  xml.Name `xml:"users"`
	UserList []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
}

func main() {
	xmlFilePointer, _ := os.Open("read-xml.xml")
	fileContentByte, _ := ioutil.ReadAll(xmlFilePointer)
	var listDataComplete ListData
	xml.Unmarshal(fileContentByte, &listDataComplete)
	xmlFilePointer.Close()

	byteData, _ := xml.Marshal(listDataComplete)
	fmt.Println(string(byteData))
}
