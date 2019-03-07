FROM golang

# Fetch dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Add project directory to Docker image.
ADD . /go/src/github.com/Madredix/theboats
WORKDIR /go/src/github.com/Madredix/theboats

RUN dep ensure
RUN go build -o app src/main.go

EXPOSE 2222
CMD ./app