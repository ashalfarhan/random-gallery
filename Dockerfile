FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./
COPY . ./

RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "./random-gallery"]

CMD [ "./random-gallery" ]

# FROM alpine

# WORKDIR /www

# COPY --from=builder /app/random-gallery ./
