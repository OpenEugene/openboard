# openboard/back

## Getting started

After ensuring that git is installed, follow these directions making changes for
your OS/ENV and skipping unnecessary steps:

```shell
# set vars
gorel="go1.12.1.linux-amd64.tar.gz"
repo="champagneabuelo/openboard"
# get and decompress go
wget https://dl.google.com/go/${gorel}
tar -C /usr/local -zxf ${gorel}
rm ${gorel}
# set envvars
export PATH=${PATH}:/usr/local/go/bin
export GO111MODULE=on
# get source
cd ${your_source_dir}
mkdir ${repo}
cd $_
git clone https://github.com/${repo} .
# build openbsrv
cd back
go build -o ${your_bin_dir}/openbsrv ./cmd/openbsrv/*.go
# run openbsrv
openbsrv -frontdir=../front/public
```

## What just happened?

`openbsrv` will open three ports (4242, 4243, 4244). 4242 is used to serve gRPC
endpoints. gRPC endpoints can be accessed directly, or via an HTTP gateway on
port 4243. In order to view the API endpoints, please visit
http://localhost:4243/v/docs. The frontend assets are served on port 4244.

## openbsrv

[openbsrv readme](./cmd/openbsrv/README.md)

## Contributing

### Please be sure to:

- Communicate before opening a PR.
- Run gometalinter and fix raised errors/warnings.
  - If required fixes seem wrong or odd, please include a note in the PR.
- Run tests and fix incorrect structure/behavior as needed.
- Add new passing tests when useful.
- Generate code when modifying protobufs (`./gengo` from the root back dir).

### Protobuf

Protocol buffers are Google's language-neutral, platform-neutral, extensible
mechanism for serializing structured data.
[Learn More](https://developers.google.com/protocol-buffers/)
[Releases](https://github.com/protocolbuffers/protobuf/releases)
[Install From
Source](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)

```shell
// basic install from source; may require changes for your OS/ENV...
git clone https://github.com/protocolbuffers/protobuf.git
cd protobuf
git submodule update --init --recursive
./autogen.sh
./configure
make
make check
sudo make install
sudo ldconfig # refresh shared library cache.
```

### protoc Plugins

```shell
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

### go-bindata

In order to embed data from files into the compiled artifact, go-bindata is
called by multiple `go generate` commands.

```shell
go get -u github.com/go-bindata/go-bindata
```

### withdraw

```shell
go get -u github.com/codemodus/withdraw
```
