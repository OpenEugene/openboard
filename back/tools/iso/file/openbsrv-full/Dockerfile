FROM alpine

RUN apk add --update git libc6-compat

RUN wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go*.tar.gz
RUN rm go*.tar.gz
ENV PATH=${PATH}:/usr/local/go/bin
ENV GO111MODULE=on

ARG codedir=/code
ARG codebin=${codedir}/bin
RUN mkdir -p ${codebin}
ENV PATH=${PATH}:${codebin}

ARG repo=champagneabuelo/openboard
ARG codesrc=${codedir}/src
RUN mkdir -p ${codesrc}/${repo}
RUN git clone https://github.com/${repo} ${codesrc}/${repo}
WORKDIR ${codesrc}/${repo}/back/cmd/openbsrv
RUN go build -o ${codebin}/openbsrv

ENV DBADDR=""
ENV DBPORT=""
ENV DBNAME=""
ENV DBUSER=""
ENV DBPASS=""
ENV MIGRATETYPE=""
ENV FRONTDIR=${codesrc}/${repo}/front/public

ENTRYPOINT openbsrv \
	${DBADDR/$DBADDR/-dbaddr=$DBADDR} \
	${DBPORT/$DBPORT/-dbport=$DBPORT} \
	${DBNAME/$DBNAME/-dbname=$DBNAME} \
	${DBUSER/$DBUSER/-dbuser=$DBUSER} \
	${DBPASS/$DBPASS/-dbpass=$DBPASS} \
	${MIGRATETYPE/$MIGRATETYPE/-$MIGRATETYPE} \
	-frontdir=${FRONTDIR}

EXPOSE 4243 4244
