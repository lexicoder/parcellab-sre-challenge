[![wakatime](https://wakatime.com/badge/user/f0ba8fe5-0102-41e9-aa49-1864bfbd9cf8/project/018d61f9-2990-481c-8f6f-a93e61ce9285.svg)](https://wakatime.com/badge/user/f0ba8fe5-0102-41e9-aa49-1864bfbd9cf8/project/018d61f9-2990-481c-8f6f-a93e61ce9285)
![example workflow](https://github.com/lexicoder/parcellab-sre-challenge/actions/workflows/build.yaml/badge.svg)

# parcellab-sre-challenge

Requires Go 1.22 or higher.

Please follow the [install instructions](https://golang.org/doc/install) for relevant version.

## Build & Test

### Build

```bash
make
```

### Test

```bash
make test
```

## Run locally

### Requirements

- [Docker](https://docs.docker.com/engine/install/)
- [Helm](https://helm.sh/docs/intro/install/)
- A local kubernetes cluster if you intend to deploy locally. [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) coudl be used for this.

### Run binary

```bash
make run
```

Access the service from `http://localhost:5000`

### Run in kubernetes cluster

```bash
make deploy-local
```

Run `kubectl port-forward svc/parcellab-sre-challenge 5000:5000` to create a port forward that will allow you test the endpoint on `http://localhost:5000`

### Configuration

The service is primarily configured through a set of environment variables.

| Environment Variable | Description                                                                             | Required | Default |
| -------------------- | --------------------------------------------------------------------------------------- | -------- | ------- |
| SERVER_PORT          | This configures the port the service listens on.                                        | `false`  | `8080`  |
| SALUTATION           | This configures the saluttation returned by the service                                 | `false`  | `Hi`    |
| SERVER_IDLE_TIMEOUT  | [IdleTimeout](https://pkg.go.dev/net/http#Server.IdleTimeout) setting for http server   | `false`  | `30s`   |
| SERVER_READ_TIMEOUT  | [ReadTimeout](https://pkg.go.dev/net/http#Server.ReadTimeout) setting for http server   | `false`  | `15s`   |
| SERVER_WRITE_TIMEOUT | [WriteTimeout](https://pkg.go.dev/net/http#Server.WriteTimeout) setting for http server | `false`  | `15s`   |

## CI/CD

The app is built, tested and deployed using GitHub Actions workflows which can be found in the [.github/workflows](.github/workflows) folder.
