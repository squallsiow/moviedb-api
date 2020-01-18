FROM golang:latest 
ADD . /go/src/github.com/moviedb-api
WORKDIR  /go/src/github.com/moviedb-api



ENV ADMIN_SECRET 'CflFPa89BzSiVdamikDavDBpKtC9A2zk'
ENV API_KEY 'b95d785d64a4e396406586a175e7955c'
ENV DEFAULT_IMAGE_FOLDER 'Data/Gallery'
ENV DEFAULT_DATASTORE_FILEPATH 'Data/DB'
ENV DEFAULT_DATASTORE_FILE 'moviedb.db'

RUN make setenv
RUN make build

ENTRYPOINT ./moviedb-api

