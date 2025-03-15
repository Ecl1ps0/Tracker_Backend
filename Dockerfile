FROM golang:alpine AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/main.go

FROM python:slim

WORKDIR /Proctor

RUN groupadd -r proctor && useradd -r -g proctor proctor

RUN mkdir -p /Proctor && chown -R proctor:proctor /Proctor

USER proctor

COPY --from=build /app .

RUN apt-get update && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

COPY ./ml_model/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

ENTRYPOINT ["./main"]