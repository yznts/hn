# -------------
# build stage
# -------------
FROM golang:alpine AS build

# Attach sources
WORKDIR /src
ADD . /src

# System deps
RUN apk add --no-cache git npm

# Build
RUN go build -o hn
RUN (cd static; npm i; npm run build)

# -------------
# runtime stage
# -------------
FROM alpine

# Copy app
WORKDIR /app
COPY --from=build /src/hn /app/
COPY --from=build /src/*.html /app/
COPY --from=build /src/static/dist /app/static/dist

# Entrypoint
ENTRYPOINT PORT=25025 ./hn

