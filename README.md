# CLI tool to check weather using Golang

```bash
git clone https://github.com/e-for-eshaan/weather-cli.git sun && cd sun && go mod tidy && touch .env && code .env
```

Paste your weather-api key in the .env (get one from https://www.weatherapi.com/)
```
WEATHER_API_KEY=<API_KEY>
```

Save and close `.env`

In the `sun` folder's terminal, run:

```bash
go build && sudo mv sun /usr/local/bin
```

# Usage

Get weather of location:

```bash
# sun <location>

sun london
```

Get weather forecast with skips:
```bash
# sun <location> <skip-hours>

sun london 6
```

Get weather forecast with skips for `n` days
```bash
# sun <location> <skip-hours> <days>
# forecast for 7 days, with 3 hour skips

sun london 3 7 
```
 
