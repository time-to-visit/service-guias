FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .


RUN cd cmd  && go  build -o /service-guides

EXPOSE 3003

CMD [ "/service-guides" ]