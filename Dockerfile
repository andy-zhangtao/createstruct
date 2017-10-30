FROM alpine:latest

RUN apk update && \
    apk add libc6-compat ca-certificates wget openssl&& \
    update-ca-certificates

COPY createstruct /createstruct
COPY README.md /readme.md

CMD ["/createstruct"]