FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./
COPY . ./

RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "./random-gallery"]

FROM alpine

WORKDIR /

COPY --from=builder ./app/random-gallery /bin

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT [ "/bin/random-gallery" ]
