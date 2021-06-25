# What is this?
Just an nginx server listening on 443 with a self signed cert.

## Runs these shell commands before applying the yaml:

KEY=/tmp/nginx.key
CERT=/tmp/nginx.crt

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${KEY} -out ${CERT} -subj "/CN=nginxsvc/O=nginxsvc"

kubectl create secret tls nginx-certs --key ${KEY} --cert ${CERT}
