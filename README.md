# Golearn
## Learning Go <br>
I have just started learning Go.<br>
This repository contains very basic programs in Go.<br>
Some points to be noted:
* Check for your GOPATH and put your program there only.
* Use `go env` to check you GOPATH.
* If you want to run the program from your directory then change the GOPATH to that directory.

# [docker_learn](https://github.com/utkarshmani1997/golearn/tree/master/docker_learn)
Steps to follow:
* Install docker
* Make a Dockerfile for your progam.
* First build the binaries using `sudo docker built -t imagename .` ('.' refers your current directory)
* To display the images use `sudo docker images`
* To display the running containers `sudo docker ps`
* To remove the image run `sudo docker image rm imagename` or `sudo docker image rm -f imagename`
* To stop the container run `sudo docker stop container_id`

## [fetchSource](https://github.com/utkarshmani1997/golearn/tree/master/docker_learn/fetchSource)
* Now start container using `sudo docker run -t imagename:latest arg1 arg2 ...` {for example pass http://www.yourwebsite.com as arg1}
## [SampleApp](https://github.com/utkarshmani1997/golearn/tree/master/docker_learn/SampleApp)
* Start container using `sudo docker run -p 8080:8080 imagename:latest`
