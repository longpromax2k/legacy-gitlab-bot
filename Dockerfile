FROM golang:1.18.5-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /gitlabhook

EXPOSE 8080

CMD [ "/gitlabhook" ]