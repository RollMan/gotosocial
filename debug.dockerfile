FROM golang:1.21

ENV DEBUG=1

RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /go/src/github.com/superseriousbusiness/gotosocial
WORKDIR /go/src/github.com/superseriousbusiness/gotosocial
RUN DEBUG=1 scripts/build.sh
CMD ["sh", "-c", "dlv debug ./cmd/gotosocial -- testrig start"]
