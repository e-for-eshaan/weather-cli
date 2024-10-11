# CLI tool to check weather using Golang

```
mkdir sun && cd sun
git clone https://github.com/e-for-eshaan/weather-cli.git .
go mod tidy
go build
touch .env
code .env
```

Paste your weather-api key

Save and close the editor that opened `.env`

In the `sun` folder's terminal, run:

```
go build && sudo mv /usr/local/bin
```