package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct{
	Name string
	Age int
	Meta UserMeta
}

type UserMeta struct{
	Visits int
}

func main() {
	fmt.Println("Experimental main")
	
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil{ 
		panic(err)
	}
	
	user := User{
		Name: "John Smith",
		Age: 49,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	// Anonymous struct
/*	user:= struct {
		Name string
	}{
		Name: "John Smith",
	}*/



	err = t.Execute(os.Stdout, user)
	if err != nil{
		panic(err)
	}
	
}
