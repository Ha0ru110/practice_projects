FROM alpine:3.15
RUN apk add --no-cache tzdata ca-certificates
COPY ./url /opt/
CMD /opt/url
