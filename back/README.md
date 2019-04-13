# openboard/back

## Getting started

After ensuring that git is installed, follow these directions making changes for
your OS/ENV and skipping unnecessary steps:

### The Database

#### As Local Install

```shell
sudo su -
apt update
apt install mariadb-server
mysql_secure_installation
systemctl stop mariadb
mysqld_safe --skip-grant-tables &
mysql -u root
```

```sql
MariaDB> USE mysql;
MariaDB> UPDATE qUser SET Plugin='mysql_native_password';
MariaDB> FLUSH PRIVILEGES;
MariaDB> quit;
```

```shell
pkill -9 mysqld
systemctl start mariadb
mysql -u root -p -e "CREATE DATABASE openeug_openb_dev DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;"
mysql -u root -p -e "GRANT ALL PRIVILEGES ON openeug_openb_dev.* TO openeug_openbdev@'localhost' IDENTIFIED BY '{your_db_pass}';"
```

#### As Docker Container

```shell
./iso/dev
# with subcommands [up|down|ip|clean] (default: up)
```

### The Application

```shell
gorel="go1.12.4.linux-amd64.tar.gz"
repo="champagneabuelo/openboard"

cd /usr/local
sudo wget https://dl.google.com/go/${gorel}
sudo tar -C /usr/local -zxf ${gorel}
sudo rm ${gorel}

export PATH=${PATH}:/usr/local/go/bin
export GO111MODULE=on

cd {your_source_code_dir}
mkdir -p ${repo}
cd $_
git clone https://github.com/${repo} .

cd back/cmd/openbsrv
go build -o {your_bin_dir}/openbsrv

# only include the DB address if not 127.0.0.1/localhost.
openbsrv -frontdir=../front/public --dbpass={your_db_pass} --dbaddr={your_db_addr}
```

### What just happened?

`openbsrv` will open three ports (4242, 4243, 4244). 4242 is used to serve gRPC
endpoints. gRPC endpoints can be accessed directly, or via an HTTP gateway on
port 4243. In order to view the API endpoints, please visit
http://localhost:4243/v/docs. The frontend assets are served on port 4244.

### openbsrv

[openbsrv readme](./cmd/openbsrv/README.md)

## Tooling

### Simple Tools

#### go-bindata

In order to embed data from files into the compiled artifact, go-bindata is
called by multiple `go generate` commands.

- https://github.com/go-bindata/go-bindata

#### protoc Plugins

protoc plugins are leveraged by the protoc command to generate powerful code
easily.

- https://github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
- https://github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
- https://github.com/golang/protobuf/protoc-gen-go

#### withdraw

withdraw is a go-based file removal tool to ensure go generate commands are
portable.

- https://github.com/codemodus/withdraw

### Simple Tools Installation

```shell
cd ./tools
./install
```

### Protobuf (protoc)

Protocol buffers are Google's language-neutral, platform-neutral, extensible
mechanism for serializing structured data.

[Learn More](https://developers.google.com/protocol-buffers/) |
[Releases](https://github.com/protocolbuffers/protobuf/releases) |
[Install From
Source](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)

### Protobuf (protoc) Installation

```shell
# basic install from source; may require changes for your OS/ENV...
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

## Contributing

### Please be sure to:

- Communicate before opening a PR.
- Run gometalinter and fix raised errors/warnings.
  - If required fixes seem wrong or odd, please include a note in the PR.
- Run tests and fix incorrect structure/behavior as needed.
- Add new passing tests when useful.
- Generate code when modifying protobufs (`./gengo` from the root back dir).
