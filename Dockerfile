FROM xhofe/alist:latest
LABEL MAINTAINER="i@nn.ci"

ARG DATABASE_URL

WORKDIR /opt/alist/
ENV DB_TYPE postgres
ENV DB_SLL_MODE require
ADD main /main
RUN chmod +x /main
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]