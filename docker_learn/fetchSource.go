package main      //The package “main” tells the Go compiler that the package should compile as an executable program 
                  //instead of a shared library.

import (
  "fmt"         //importing the package that contains some predefined methods.
// "flag"        //importing this package to do something with the command line arguments with the predefined methods 
                //existing in this package. 
  "os"          //importing this package for the system calls. 
  "net/http"    //importing this package for retrieving the source-code of the webpage requested.
  "io/ioutil"   //importing this package for reading and printing the source-code fetched.

)

func main() {
                 
    flag.Parse()      
  args := flag.Args()
    fmt.Println(args)       
    if len(args) < 1 {        
    fmt.Println("Please Enter the URL")  
    os.Exit(1)                                
  }
    retrieve(args)        //call the retrieve function                                   
}

func retrieve(url string){          //gives the  source code as output.
    
    resp, err := http.Get(url)         
    if err != nil{
      fmt.Println("read error is:", err)
      return 
    }
    body, err := ioutil.ReadAll(resp.Body);
    if err != nil{ 
      fmt.Println("read error is:", err)
      return 
    } else{
      fmt.Println(string(body))
      
    }
}
