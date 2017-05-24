package main

//import "fmt"
import "os/exec"

func main() {
// var count int
// var url string
// fmt.Printf("Enter the Url:")
// fmt.Scanln(&url)    //enter the filename with extension

 exec.Command("cvlc", "/home/infinity/Music/01 - Sanam Re - Sanam Re [DJMaza.Info].mp3").Start() //opens the file name passed into the argumen
}
