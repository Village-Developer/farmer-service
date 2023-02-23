FROM golang

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

# COPY *.go ./

COPY . ./

RUN go build -o /farmer-service

EXPOSE 9100

CMD ["/farmer-service"]


