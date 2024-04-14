FROM golang

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /server

EXPOSE 5000
CMD [ "/server" ]