FROM alpine

RUN apk add --update git libc6-compat

RUN wget https://dl.google.com/go/go1.12.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go*.tar.gz
RUN rm go*.tar.gz
ENV PATH=${PATH}:/usr/local/go/bin
ENV GO111MODULE=on

ARG src=/code
ARG go_bin=${src}/bin
RUN mkdir -p ${go_bin}
ENV PATH=${PATH}:${go_bin}

ARG repo=champagneabuelo/openboard
RUN mkdir -p ${src}/${repo}
RUN git clone https://github.com/${repo} ${src}/${repo}
WORKDIR ${src}/${repo}/back
RUN go build -o ${go_bin}/openbsrv ./cmd/openbsrv/*.go

ENTRYPOINT openbsrv -frontdir=${src}/${repo}/front/public

EXPOSE 4243 4244
