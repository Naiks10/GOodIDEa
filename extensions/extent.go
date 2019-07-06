package extensions

import (
	"fmt"
)

type Extension struct{
	Name string
	Description string
	Link string
}

func (m Extension) add() bool{
	script := "INSERT INTO EXTENSIONS (NAME_, DESC_, LINK_ ) VALUES ("+m.Name+","+m.Description+","+m.Link+")"
	fmt.Println(script + " executed")
	return true
}