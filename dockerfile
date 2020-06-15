FROM golang:alpine3.12 as builder

RUN mkdir /app
RUN chmod 700 /app

COPY . /app

RUN go get -u github.com/helmutkemper/iotmaker.docker
RUN go get -u github.com/helmutkemper/iotmaker.network.stability.pygocentrus

# import golang packages to be used inside image "scratch"
ARG CGO_ENABLED=0
RUN go build -o /app/main /app/main.go

WORKDIR /app

CMD ["/main"]