FROM docker.io/golang:1.25

WORKDIR /app
COPY main.go go.mod /app/
COPY cmd /app/cmd

RUN go build main.go

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
