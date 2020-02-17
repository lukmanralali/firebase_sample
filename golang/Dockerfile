FROM golang:1.12-alpine

RUN apk update && \
	apk add bash git && \
    apk add --update make

# We create an /app directory within our image that will hold our application source files
RUN mkdir /go/src/firebase_sample
# We copy everything in the root directory into our /app directory
COPY . /go/src/firebase_sample
# We specify that we now wish to execute any further commands inside our /app directory
WORKDIR /go/src/firebase_sample
# add project maintainer
RUN go get -u github.com/golang/dep/cmd/dep
# add auto reloading app when change available
RUN go get github.com/markbates/refresh
# RUN dep init
# RUN dep ensure
# we run go build to compile the binary executable of our Go program
# RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
# CMD ["make","run"]