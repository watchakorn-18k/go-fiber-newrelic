<div align="center">

# go-fiber-newrelic

![img](docs/image.png)

</div>

## Install golang package

```bash
go mod tidy
```

# Start APP

```sh
go run . || go run main.go
```

## Air

```sh
go install github.com/cosmtrek/air@latest
```

```sh
air
```

# Podman

```
podman build -t fiber-test .
```

```
podman run --rm -it -p 3000:3000 fiber-test
```

# New Relic
- Go to [New Relic](https://docs.newrelic.com/)
- Get License Key
- Go to tab Traces follow the function