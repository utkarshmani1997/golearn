FROM golang:1.7-alpine

ADD . /home
        
WORKDIR /home

RUN ["go", "build"]

ENTRYPOINT ["./home"]
