package main    //The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.

import (
  "fmt"         //importing the package that contains some predefined methods.           
  "net/http"    //importing this package for retrieving the source-code of the webpage requested.
  "io/ioutil"   //importing this package for reading and printing the source-code fetched.
)               

func main() {
  resp, err := http.Get("https://utkarshmani1997.github.io/Aavesh/") //two variables declared:one is for getting the response or 
                                                                     //the data  from the requested url and other is for getting the error
                                                                     //message in the case if it has encountered any error so that need to be 
                                                                     //resolved there itself.  
                                                  
  fmt.Println("http transport error is:", err)     

  body, err := ioutil.ReadAll(resp.Body)           //read all the source-code in the body.  
                                          
                                          
  fmt.Println("read error is:", err)              

  fmt.Println(string(body))                         
} 
