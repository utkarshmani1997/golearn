FROM golang:1.7-alpine

ADD . /home
        
WORKDIR /home
# This builds a binary first
RUN ["go", "build"]
# Now you can run the executable and pass arguments at the run time.
ENTRYPOINT ["./home"]
