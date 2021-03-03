# openboard/back

## Getting Started

### Easy Mode

The following scripts should be working for Bash on Linux, WSL, and Darwin. When
installing the database, the suggested defaults for this project are:
dbname = "openeug_openb_dev", and dbuser = "openeug_openbdev".

```sh
# from {project_root}
cd back/tools
./install-go
./install-tools
./install-mariadb # local install (optional)
```

```sh
# alternate database setup via container (optional - skip if using mariadb "local install")
# from {project_root}
cd back/tools/iso/
./dev up # subcommands [up|dn|ip|clean] (default: up)
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
cd {your_source_code_dir}
mkdir -p github.com/protocolbuffers/protobuf
cd github.com/protocolbuffers/protobuf
git clone https://github.com/protocolbuffers/protobuf.git .
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
mkdir -p github.com/OpenEugene/openboard
cd github.com/OpenEugene/openboard
git clone https://github.com/OpenEugene/openboard .
```

### Build and Execute

```sh
# from {project_root}
cd back/cmd/openbsrv
go build
./openbsrv --dbpass={your_dbpass} --migrate
# be careful not to add/commit the executable
```

Please refer to the [openbsrv readme](./cmd/openbsrv/README.md) for more details
about usage and flags (e.g. -dbname, -dbuser, etc.).

### Database Migration Management

```sh
# from {project_root}
cd back/cmd/openbsrv
./openbsrv --dbpass={your_dbpass} --rollback
^C # Ctrl-C will send the system signal "SIGINT" and halts the program
./openbsrv --dbpass={your_dbpass} --migrate
```

### Run Tests

#### Unit Tests

None at this time.

#### Surface Tests

None at this time.

#### End-to-end Tests

A convenience script has been provided to build and run the executable, and also
run the end-to-end tests.

```sh
# from {project_root}
cd back/tests/openbsrv
./run-tests
```

## Contributing

- [Start here](../docs/CONTRIBUTING.md)
