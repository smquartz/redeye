FROM alpine:3.2
ADD redeye-srv /redeye-srv
ENTRYPOINT [ "/redeye-srv" ]
