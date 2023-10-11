FROM xhofe/alist:v3.28.0
LABEL MAINTAINER="i@nn.ci"

ARG DATABASE_URL

WORKDIR /opt/alist/
ENV DB_TYPE postgres
ENV DB_SSL_MODE require
ADD main /main
RUN chmod +x /main
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]
