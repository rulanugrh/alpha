FROM golang:1.18-alpine

RUN apk --no-cache add ca-certificates git

WORKDIR /app/run

COPY . .

RUN go mod tidy
RUN go build -o server

EXPOSE 5000

CMD [ "./server" ]