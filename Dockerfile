FROM golang:1.21
LABEL authors="shivanshtamrakar"
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /kahawatein
EXPOSE 8080

ENTRYPOINT /kahawatein



