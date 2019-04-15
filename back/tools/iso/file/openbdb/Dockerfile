FROM mariadb:10.2

RUN apt-get update && apt-get dist-upgrade -y

ENV MYSQL_ROOT_PASSWORD=""
ENV MYSQL_DATABASE=""
ENV MYSQL_USER=""
ENV MYSQL_PASSWORD=""
ENV MYSQL_ALLOW_EMPTY_PASSWORD = no
ENV MYSQL_RANDOM_ROOT_PASSWORD = no

ARG mycnf="/etc/mysql/my.cnf"
RUN sed -Ei 's/^(.*)(bind.*ress\s*=\s*)[.0-9]{7,15}/\20.0.0.0/g' ${mycnf}
ARG mycnf="/etc/mysql/conf.d/mariadb.cnf"
RUN sed -Ei 's/^(.*)(char.*rver\s*=\s*)[[:alnum:][:punct:]]*/\2utf8mb4/g' ${mycnf}
RUN sed -Ei 's/^(.*)(coll.*rver\s*=\s*)[[:alnum:][:punct:]]*/\2utf8mb4_unicode_ci/g' ${mycnf}

EXPOSE 3306
