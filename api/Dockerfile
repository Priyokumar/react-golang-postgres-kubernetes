FROM golang:1.21.5-bookworm as build

WORKDIR /app
COPY routes ./routes/
COPY db ./db/
COPY app.go ./
COPY go.mod ./
COPY go.sum ./
RUN go build -o api
RUN ls
EXPOSE 8000
CMD [ "./api" ]