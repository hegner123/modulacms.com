FROM node:25-alpine AS css
WORKDIR /src
COPY package.json package-lock.json ./
RUN npm ci
COPY styles/ styles/
COPY components/ components/
RUN npx @tailwindcss/cli -i styles/app.css -o static/app.css --minify

FROM golang:1.25-alpine AS build
RUN go install github.com/a-h/templ/cmd/templ@v0.3.977
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=css /src/static/app.css static/app.css
RUN templ generate
RUN CGO_ENABLED=0 go build -o /bin/modulacms-site .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/modulacms-site /bin/modulacms-site
COPY --from=build /src/static /static
WORKDIR /
EXPOSE 5050
ENTRYPOINT ["/bin/modulacms-site"]
