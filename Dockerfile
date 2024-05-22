FROM golang:1.22

COPY . .

RUN apt-get update && apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash
RUN apt-get install -y nodejs

RUN npm install -g tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download

RUN npx tailwindcss -i ./assets/css/input.css -o ./assets/css/tailwind.css --minify
RUN templ generate
RUN go build -o ./webapp ./src/

EXPOSE 3000
CMD ["./webapp"]