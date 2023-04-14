FROM golang:1.17.6-alpine as Build

COPY . .

#build our binary at root location. Binary name will be main. We are using go modules so gpath env variable should be empty.
# CGO_ENABLED=0 is required to build a static binary, otherwise the shared libraries might not be present and it will fail to run
# with: no such file or directory
RUN CGO_ENABLED=0 GOPATH= go build -o /main main.go
# if we don't install the CA certs all the connections to any https endpoint will fail complaining the cert is signed by
# an unknown authority
RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=Build /main /go/bin/main
COPY --from=Build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD [ "/go/bin/main" ]