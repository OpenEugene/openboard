# openboard/back

## Getting Started (Easy Mode)

The following scripts should be working for Bash on Linux and WSL. Apple users
should refer to the Normal Mode section.

```shell
./tools/install-go
./tools/install-tools
./tools/install-mariadb # local install (optional)
```

```shell
# skip if using mariadb local install
pushd ./tools/iso/
./dev up # subcommands [up|down|ip|clean] (default: up)
popd
```

## Getting Started (Normal Mode)

Copy the following Bash scripts to your system or enter their contents manually
while making changes as needed.

- [Install Go](./tools/install-go)
- [Install Tools](./tools/install-tools)
- [Install MariaDB](./tools/install-mariadb)

## Protobuf/protoc (Advanced Mode; Optional)

Protocol buffers are Google's language-neutral, platform-neutral, extensible
mechanism for serializing structured data.

[Learn More](https://developers.google.com/protocol-buffers/) |
[Releases](https://github.com/protocolbuffers/protobuf/releases) |
[Install From
Source](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)

```shell
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

```shell
repo="champagneabuelo/openboard"
cd {your_source_dir}
mkdir -p ${repo}
cd $_
git clone https://github.com/${repo} .

cd back/cmd/openbsrv
go build -o {your_bin_dir}/openbsrv

# only include the database address if not 127.0.0.1/localhost.
openbsrv -frontdir=../../../front/public/ --dbpass={your_dbpass} --dbaddr={your_dbaddr}
```

[openbsrv readme](./cmd/openbsrv/README.md)

## Contributing

### Please be sure to:

- Communicate before opening a PR.
- Run gometalinter and fix raised errors/warnings.
  - gometalinter install directions are highly dependent on choice of dev
  environment.
  - If gometalinter suggests odd/wrong fixes, please communicate before applying
  them.
- Run tests and fix incorrect structure/behavior as needed.
- Add new passing tests when useful.
- Generate code when modifying protobufs (`./gengo` from the root back dir).
