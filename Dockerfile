
# =============== build and run ===============
# build:  docker build -t travel-api .
# run:    docker run hello-world


# =============== build stage ===============
FROM golang:1.16.5-buster AS build
# env
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"
# dependent
WORKDIR /app
COPY go.* ./
RUN go mod download -x all
# build
COPY . ./
# ldflags:
#  -s: disable symbol table
#  -w: disable DWARF generation
# run "go tool link -help" to get the full list of ldflags
RUN go env && go build -ldflags "-s -w" -o service -v ./main.go

# =============== final stage ===============
FROM alpine:latest AS final
# resources
WORKDIR /app
COPY --from=build /app/service ./
COPY --from=build /app/conf/base.toml ./conf/base.toml
#COPY --from=build /app/conf/prod ./conf/prod

# set time zone
RUN apk add --no-cache tzdata

EXPOSE 80
ENTRYPOINT ["env","GO_ENV=prod","/app/service", "-other", "flags"]

