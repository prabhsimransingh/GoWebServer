# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/git.soma.salesforce.com/prabh-singh/GoWebServer

# Build the golang-docker command inside the container.
RUN go install git.soma.salesforce.com/prabh-singh/GoWebServer

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/GoWebServer

# http server listens on port 8080.
EXPOSE 8080