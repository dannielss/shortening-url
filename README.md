<h3 align="center">
   <img src="https://user-images.githubusercontent.com/58083563/193971505-32402d48-42cc-4ac2-be84-9f0a7a470784.svg" alt="Go" width="100" />
</h3>
<h3 align="center">URL Shortener</h3>

<p align="center">
  <a href="#gift-Features">Features</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#wrench-Configuration">Configuration</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#rocket-Technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-Usage">Usage</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

## :gift: Features

- Shorten long URLs
- Redirect users from short URLs to the original links
- Observability with Prometheus and Grafana
- Cache with redis
- Swagger docs

## :wrench: Configuration

`Golang version: 1.24.1`

```
1. Copy .env and add redis address
$ cp .env.example .env

2. Install all dependencies
$ go get . (or go mod tidy)

3. Start application
$ make dev

3. Check all commands
$ make help
```

## :rocket: Technologies

- Golang
- Gin-Gonic
- Grafana
- Promtheus
- Redis
- Swagger

## :memo: Usage

- Shorten a URL
  - Send a POST request to `/shorten` with a JSON payload containing the url to be shorted.
  - Example using `curl`:


  ```bash
    curl -X POST http://localhost:3000/shorten \
    -H "Content-Type: applicaiton/json" \
    -d '{"url": "https://example.com"}'
  ```

- Access the shortened URL
  - Navigate to `http://localhost:3000/{shorted_url}` in your browser or via HTTP request.

- Metrics
  - Access Grafana dashboards at `http://localhost:3001`
  - Access Prometheus metrics at `http://localhost:3000/metrics`

- Docs
  - Access Swagger documentation at `http://localhost:3000/swagger/index.html`
