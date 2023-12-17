FROM golang:1.20
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY question_api/ ./question_api

RUN go build -o /quizknow

EXPOSE 8080

CMD ["/quizknow"]

