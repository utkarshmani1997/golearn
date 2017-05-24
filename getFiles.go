package main

import (
    "fmt"
    "time"
    "os/exec"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func openMusic(t time.Time) {
     
    exec.Command("xdg-open", "/home/infinity/Music/01 - Sanam Re - Sanam Re [DJMaza.Info].mp3").Start()

}

func main() {
  var dir string
  fmt.Println("Enter Music Location:") 
  fmt.Scanln(&dir)                         //for exp: /home/infinity/Music/01 - Sanam Re - Sanam Re [DJMaza.Info].mp3
  exec.Command("xdg-open", dir).Start()
  doEvery(2*time.Hour, openMusic)
  
}
