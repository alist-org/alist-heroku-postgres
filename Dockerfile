FROM xhofe/alist:latest
LABEL MAINTAINER="i@nn.ci"

WORKDIR /opt/alist/
ENV DB_SLL_MODE require
ADD main /main
RUN chmod +x /main
RUN /main

CMD [ "./alist", "-docker" ]