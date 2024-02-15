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

### Run locally

#### Requirements

* [Docker](https://docs.docker.com/engine/install/)
* [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* [Helm](https://helm.sh/docs/intro/install/)

#### Steps

```bash
make deploy-local
```

Run `kubectl port-forward svc/parcellab-sre-challenge 5000:5000` to create a port forward that will allow you access the app on `http://localhost:5000`

## CI/CD

The app is built, tested and deployed using GitHub Actions workflows which can be found in the `.github/workflows` folder.
