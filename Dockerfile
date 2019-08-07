FROM alpine:3.9
LABEL Name=hub-quay-agent \
      Release=https://github.com/appvia/hub-quay-agent \
      Maintainer=gambol99@gmail.com \
      Url=https://github.com/appvia/hub-quay-agent \
      Help=https://github.com/appvia/hub-quay-agent/issues

RUN apk add --no-cache ca-certificates curl

COPY bin/hub-quay-agent /hub-quay-agent

USER 65534

ENTRYPOINT [ "/hub-quay-agent" ]
