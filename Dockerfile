FROM golang
#Add our application to golang base image
ADD . /go/src/github.com/craguilar/cidm
#Install the application in bin
RUN go install github.com/craguilar/cidm/cmd/cidm-server
# Run the cidm=server command by default when the container starts on specified port
ENTRYPOINT ["/go/bin/cidm-server", "--port", "30030"]
# Document that the service listens on port 30030
EXPOSE 30030