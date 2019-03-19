# openboard/back

## Getting started

After ensuring that git is installed, follow these directions making changes for
your OS/ENV and skipping unnecessary steps:

  # set vars
  gorel="go1.12.linux-amd64.tar.gz"
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
[Learn More](https://developers.google.com/protocol-buffers/).
