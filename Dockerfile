FROM xhofe/alist:latest
LABEL MAINTAINER="i@nn.ci"

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ADD main /main
RUN chmod +x /main

CMD /entrypoint.sh