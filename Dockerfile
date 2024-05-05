## Stage 1 Build binary file
FROM golang:1.18.5-alpine as build

ARG APP
# ARG APPENV

WORKDIR /go/${APP}
COPY ${APP} ./

# Add Config file
# ADD ${APP} ${APPENV}

# Build 
# WORKDIR ${APP}
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${APP} main.go

EXPOSE ${PORT}

CMD ./${APP}

## Stage 2 Copy binary file 
FROM alpine:3.16.2

LABEL Author="webmaster@gps.id"
LABEL Company="PT SUPER SPRING" Website="https://superspring.co.id"

RUN apk add --no-cache libc6-compat

ARG APP
ENV APP=${APP}


COPY --from=build /go/${APP} /srv/
RUN chmod +x /srv/${APP}

WORKDIR /srv/

EXPOSE ${PORT}

CMD ./${APP}