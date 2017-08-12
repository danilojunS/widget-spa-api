FROM golang

ENV PORT 4000

WORKDIR /go/src/github.com/danilojunS/widgets-spa-api/
ADD . /go/src/github.com/danilojunS/widgets-spa-api/

RUN make install
RUN make build

EXPOSE 4000

CMD make start-prod