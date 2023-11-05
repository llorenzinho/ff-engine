# Feature Flag backend

This app should be deployed in an cluster which is running the [FeatureFlag Operator](https://github.com/llorenzinho/ff-operator).

Feature flag resources can be deployed in the cluster and their status can be queried using this app.

## Usage

This app expose a variety of endpoints to query the status of the feature flags:

The following endpoint can be used to manage the K8s resources itself:

- **GET** `/api/v1/featureflags`: List all the feature flags in the cluster. It will return a list of `FeatureFlag` resources.
- **GET** `/api/v1/featureflags/{name}`: Get the status of a specific feature flag. It will return a `FeatureFlag` resource.
- **POST** `/api/v1/featureflags`: Create a new feature flag. DTO is available [here](./models/crud/ffmodels.go)
- **PUT** `/api/v1/featureflags`: Update the fiven featureflag. DTO is available [here](./models/crud/ffmodels.go)
- **DELETE** `/api/v1/featureflags/{name}`: Delete the given feature flag.

The following endpoints can be used to query the status of the feature flags and **should be used by the applications**:

- **GET** `/api/v1/featureflags/{name}/active`: return a json with the status of the given FF and if it is enabled or not
- **PUT** `/api/v1/featureflags/{name}/enable`: enable a given FF
- **PUT** `/api/v1/featureflags/{name}/disable`: disable the given FF
- **PUT** `/api/v1/featureflags/{name}/activate`: activate the given FF
- **PUT** `/api/v1/featureflags/{name}/deactivate`: deactivate the given FF

All the above endpoints accept the following query parameters:

- `namespace`: The namespace where the feature flag is deployed. If not provided, it will use the default namespace defined in [config.json](./config/config.json) file.

## Glossary

A feature flag can be **enabled** or not and even **active** or not:

- **Enabled**: The feature flag is enabled in the cluster. It means that the feature flag is deployed in the cluster and it is ready to be used. If this value is false, it's status will always be false.
- **Active**: The feature flag is active. It means that the feature flag is enabled and it is ready to be used. This value can be used by the applications to decide what to show to the users.

## Development

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- Running kubernetes cluster (e.g. [minikube](https://minikube.sigs.k8s.io/docs/start/))
- [FeatureFlag Operator](https://github.com/llorenzinho/ff-operator) and crds installed in the cluster

### Run locally

Run `docker-compose up` to run the application locally. You must have a k8s cluster running locally and need to copy the kubeconfig file to the root of the project.

## Build

This app is using the [k8s code generator](https://github.com/kubernetes/code-generator) to build the clientset based on the given CRDs.

You need to install the code-generator tool:

```bash
go get k8s.io/code-generator
```

Rebuild the clientset only if the CRDs are changed:

```bash
./hack/update-codegen.sh
```
