FROM alpine:latest

ADD simple /simple

ENTRYPOINT [ "/simple" ]
