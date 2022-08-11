FROM golang:1.18.5-bullseye

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o /gitlabhook

EXPOSE 8080

CMD [ "/gitlabhook" ]