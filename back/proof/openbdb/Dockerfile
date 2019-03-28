FROM mariadb:10.2

RUN apt-get update && apt-get dist-upgrade -y

ENV MYSQL_ROOT_PASSWORD=""
ENV MYSQL_DATABASE=""
ENV MYSQL_USER=""
ENV MYSQL_PASSWORD=""
ENV MYSQL_ALLOW_EMPTY_PASSWORD = no
ENV MYSQL_RANDOM_ROOT_PASSWORD = no

ARG mycnf="/etc/mysql/my.cnf"
RUN sed -Ei 's/^(.*)(bind-address\s+=\s+)[0-9][.0-9]{5,13}[0-9]/\20.0.0.0/g' ${mycnf}

EXPOSE 3306
