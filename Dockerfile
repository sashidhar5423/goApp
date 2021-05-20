FROM golang
WORKDIR /app
COPY go.mod, go. ./
RUN go.mod download
COPY . .
RUN go build -o main .
EXPOSE 5252
CMD ["./main"]



