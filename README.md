# simple http/https golang server

## run
```
# http
go run http.go

# https
openssl genrsa -out self-signed.key 2048
openssl req -new -x509 -sha256 -key self-signed.key -out self-signed.crt -days 36500
go run https.go
```
