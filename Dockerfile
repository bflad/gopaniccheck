FROM golang:1.13-stretch
WORKDIR /src
COPY gopaniccheck /usr/bin/gopaniccheck
ENTRYPOINT ["/usr/bin/gopaniccheck"]
CMD ["./..."]
