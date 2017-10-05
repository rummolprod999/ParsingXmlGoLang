package main

import "github.com/jlaffaye/ftp"
import "fmt"



func Ftp() error {
	client, err := ftp.Dial("ftp.zakupki.gov.ru:21")
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := client.Login("free", "free"); err != nil {
		fmt.Println(err)
		return err
	}
	client.ChangeDir("fcs_regions")
	l, _ := client.NameList("")
	for _, t := range l {
		fmt.Println(t)
		//fmt.Println(a)
	}
	return nil
}
