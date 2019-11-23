# Simple Go App

[![CircleCI](https://circleci.com/gh/jemisonf/simple_go_app.svg?style=svg)](https://circleci.com/gh/jemisonf/simple_go_app)

This is a simple go app that I built because I wanted a containerizable application I could use if I ran into something container-related that I wanted to test.

## Requirements

Uses Go 1.13, or Docker.

## How to use it

### For local development:
```
go run app.go
```

This will use port 8080 by default.

### Using Docker:

```
docker run -p 8080:8080 docker.io/jemisonf/simple_go_app
```

This will start the server on port 8080 of your local machine.
