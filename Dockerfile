FROM golang:1.23-alpine

RUN apk add --no-cache gcc musl-dev

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  user

WORKDIR /app

COPY . .

RUN apk add --no-cache npm
RUN npm install && npm run build:css

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /server .

RUN mkdir -p /app/data && chown -R user:user /app/data

USER user:user

EXPOSE 8080

CMD ["/server"]
