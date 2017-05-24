# golearn
Learning Go
I have just started learning Go.
This repository contains very basic programs in Go.
Some points to be noted:
* Check for your GOPATH and put your program there only.
* If you want to run the program from your directory then change the GOPATH to that directory.

# docker_learn
Steps to follow:
* install docker
* First build the binaries using `docker built -t imagename .`
* Now run using `docker run -t imagename:latest arg1` {for example pass http://www.yourwebsite.com as arg1}
* To display the images use `docker images`
* To display the active containers `docker ps -a`
* To remove the image run `docker image -rm imagename` or `docker image -rm -f imagename`
