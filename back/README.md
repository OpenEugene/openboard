# openboard/back

## Getting Started

### Easy Mode

The following scripts should be working for Bash on Linux, WSL, and Darwin. When
installing the database, the suggested defaults for this project are:
dbname = "openeug_openb_dev", and dbuser = "openeug_openbdev".

```sh
./tools/install-go
./tools/install-tools
./tools/install-mariadb # local install (optional)
```

```sh
# alternate database setup via container (optional - skip if using mariadb "local install")
pushd ./tools/iso/
./dev up # subcommands [up|dn|ip|clean] (default: up)
popd
```

### Normal Mode

- [Install Go](https://golang.org/doc/install)
- [Install Tools](./tools/install-tools)
- [Install MariaDB (>=10.2.22)](https://www.google.com/search?q=install+mariadb+stable+on)

### Advanced Options

#### Protobuf/protoc

Protocol buffers are Google's language-neutral, platform-neutral, extensible
mechanism for serializing structured data.

[Learn More](https://developers.google.com/protocol-buffers/) |
[Releases](https://github.com/protocolbuffers/protobuf/releases) |
[Install From
Source](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)

```sh
# may require changes for your OS/ENV
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

## Project Source

`openbsrv` will open three ports (4242, 4243, 4244). 4242 is used to serve gRPC
endpoints. gRPC endpoints can be accessed directly, or via an HTTP gateway on
port 4243. In order to view the API endpoints, please visit
http://localhost:4243/v/docs. The frontend assets are served on port 4244.

### Clone

```sh
cd {your_source_code_dir}
mkdir -p OpenEugene/openboard
cd $_ # "go into" newly created directory
git clone https://github.com/OpenEugene/openboard .
```

### Build and Execute

```sh
go build
./openbsrv -frontdir=../../../front/public/ --dbpass={your_dbpass}
# be careful not to add/commit the executable
```

### Run Tests

A convenience script has been provided to build and run the executable, and also
run the end-to-end tests.

From the cmd/openbsrv directory:

```sh
cd ../../tests/openbsrv
./run-tests
```

Please refer to the [openbsrv readme](./cmd/openbsrv/README.md) for more details
about usage and flags (e.g. -dbname, -dbuser, etc.).

## Contributing

- [Start here](../docs/CONTRIBUTING.md)
