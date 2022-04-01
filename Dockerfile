FROM xhofe/alist:latest
LABEL MAINTAINER="i@nn.ci"

ARG DATABASE_URL

WORKDIR /opt/alist/
ENV DB_SLL_MODE require
ADD main /main
RUN chmod +x /main
RUN /main

CMD [ "./alist", "-docker" ]