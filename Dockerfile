FROM alpine:3.8
LABEL Name=hub-quay-agent \
      Release=https://github.com/appvia/hub-quay-agent \
      Maintainer=gambol99@gmail.com \
      Url=https://github.com/appvia/hub-quay-agent \
      Help=https://github.com/appvia/hub-quay-agent/issues

RUN apk add --no-cache ca-certificates curl

ADD bin/hub-quay-agent /hub-quay-agent

USER 65534

ENTRYPOINT [ "/hub-quay-agent" ]
