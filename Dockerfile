FROM golang:1.16-alpine

ARG PORT=3000

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE $PORT

CMD [ "/go/bin/app" ]