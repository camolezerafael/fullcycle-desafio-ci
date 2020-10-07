FROM scratch

COPY gopath/bin/src/desafio-go /src/desafio-go

ENTRYPOINT ["/src/desafio-go"]