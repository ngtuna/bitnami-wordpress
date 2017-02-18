FROM bitnami/minideb:latest

MAINTAINER ng.tuna@gmail.com

RUN install_packages ca-certificates

WORKDIR /srv/

COPY bitnami-wordpress-aws.tar.gz /srv/

RUN tar zxvf bitnami-wordpress-aws.tar.gz

EXPOSE 9000

ENTRYPOINT ["bash"]

CMD ["run.sh"]
