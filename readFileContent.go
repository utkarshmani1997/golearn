package main

import "os"   
import "fmt"


func main() {

 var i string
 fmt.Printf("Enter the Filename:")
 fmt.Scanf("%s ",&i)    //enter the filename with extension
 
file, err:=os.Open(i) //opens the file name passed into the argument.
 
 if err != nil{
   panic(err)
}
 
data := make([]byte, 300)      

 count, err := file.Read(data)  //reading the data from file

 if err != nil {
   panic(err)
}
 
 fmt.Printf("read %d bytes: %q\n", count, data[:count])  //displaying the data of the file
 
 _, err := file.Close()
 if err != nil {
		log.Fatal(err)
	}

}
