FROM golang:1.21.5

WORKDIR /app
COPY . .

RUN go build app/main.go

CMD [ "./main" ]

EXPOSE 3333