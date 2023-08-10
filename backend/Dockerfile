FROM golang:1.20.5

WORKDIR /app/server
COPY . .

EXPOSE 443
RUN go mod tidy
CMD ["go", "run", "."]