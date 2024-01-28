FROM golang:1.21

RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /go/src/github.com/superseriousbusiness/gotosocial
WORKDIR /go/src/github.com/superseriousbusiness/gotosocial
RUN scripts/build.sh
CMD ["/bin/bash"]
