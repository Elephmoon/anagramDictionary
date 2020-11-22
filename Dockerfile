FROM golang:1.14-alpine as development

ENV PROJECT_PATH=/anagramDictionary
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=0

RUN apk add --no-cache ca-certificates git make bash protobuf

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN go build -o build/anagramDictionary cmd/anagramDictionary/main.go

FROM alpine:latest AS production

WORKDIR /root/
RUN apk --no-cache add
COPY --from=development /anagramDictionary/build/anagramDictionary .
COPY --from=development /anagramDictionary/internal/static/swagger_dist ./internal/static/swagger_dist
COPY --from=development /anagramDictionary/api/ ./api/
ENTRYPOINT ["./anagramDictionary"]