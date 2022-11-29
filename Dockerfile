FROM golang:1.19.3-alpine

COPY . ./app
WORKDIR ./app
RUN go build -o /ipblock

CMD [ "/ipblock" ]

EXPOSE 50051:50051