[![codecov](https://codecov.io/gh/epam/edp-argocd-operator/branch/master/graph/badge.svg?token=VNMBQGDKJ3)](https://codecov.io/gh/epam/edp-argocd-operator)

# EDP Argo CD Operator

| :heavy_exclamation_mark: Please refer to [EDP documentation](https://epam.github.io/edp-install/) to get the notion of the main concepts and guidelines. |
| --- |

Get acquainted with the EDP Argo CD Operator, its installation process, local development, and architectural scheme.

## Overview

EDP Argo CD Operator is an EDP operator that manages the EDP Argo CD Tenants:

* Argo CD is deployed in a separate `argocd` namespace.
* Argo CD uses a `cluster-admin` role for managing cluster resources.
* The `control-plane` application is created using App of Apps approach, and its code is managed by the `control-plane` members.
* The `control-plane` is used to onboard new Argo CD Tenants (App Projects).
* The `control-plane admin` provides `JWT Token` for each `EDP Tenant`.
* The `EDP Tenant` deploys `edp-argocd-operator` in its `edpTenant` EDP namespace, and uses `JWT Token` provided by `control-plane admin`.
* The `EDP Tenant Member` manages `Argo CD Repositories` and `Argo CD Applications` using `kind: Secret` and `kind: ArgoApplication` in the `edpTenant` namespace.

See a diagram below for the details:

![edpTenant](docs/assets/edpTenant.png)

## Deployment Example

Repository:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: demo
  labels:
  # must be type of repository
    argocd.edp.epam.com/secret-type: repository
stringData:
  type: git
  url: ssh://argocd@gerrit.edp-delivery-sk-delivery-dev:30007/demo.git
  project: team-foo
#  Use insecure to work with privately hosted Git services over SSH.
#  If true, it is the same as use --insecure-skip-server-verification.
#  Optional, default - "false".
#  See: https://argo-cd.readthedocs.io/en/release-1.8/user-guide/private-repositories/#unknown-ssh-hosts
  insecure: "true"
  sshPrivateKey: |
    -----BEGIN OPENSSH PRIVATE KEY-----
    YOUR_PRIVATE_SSH_KEY
    -----END OPENSSH PRIVATE KEY-----
```

EDP Argo CD Application has the same specification as the native Argo CD Application:

```yaml
apiVersion: v1.edp.epam.com/v1alpha1
kind: ArgoApplication
metadata:
  name: demo
spec:
  project: team-foo
  destination:
    namespace: team-foo-demo
    server: https://kubernetes.default.svc
  source:
    helm:
      parameters:
        - name: image.tag
          value: master-0.1.0-1
        - name: image.repository
          value: image-repo
    path: deploy-templates
    repoURL: ssh://argocd@gerrit.edp-delivery-sk-delivery-dev:30007/demo.git
    targetRevision: master
  syncPolicy:
    syncOptions:
      - CreateNamespace=true
    automated:
      selfHeal: true
      prune: true
```

![example](docs/assets/example.png)

## Prerequisites

1. Linux machine or Windows Subsystem for Linux instance with the installed [Helm v.3](https://helm.sh/docs/intro/install/).
2. Admin access to the EDP namespace.
3. Access to Argo CD, including the generated JSON Web Token.
4. Deployed EDP project/namespace. Please follow the [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/) instruction.

## Installation

In order to install the EDP Argo CD Operator, follow the steps below:

1. To add the Helm EPAMEDP charts for a local client, run the `helm repo add` command:

    ```bash
    helm repo add epmdedp https://epam.github.io/edp-helm-charts/stable
    ```

2. Select the available Helm chart version:

    ```bash
    helm search repo epmdedp/edp-argocd-operator -l

    NAME                          CHART VERSION   APP VERSION       DESCRIPTION
    epmdedp/edp-argocd-operator   0.2.0           0.2.0             A Helm chart for EDP Argo CD Operator
    ```

    _**NOTE:** It is highly recommended to use the latest version of the Helm chart._

3. Deploy the EDP Argo CD Operator. Chart parameters are specified in the [deploy-templates/README](deploy-templates/README.md) file.

4. Install the EDP Argo CD Operator to the <edp-project> namespace with the following Helm command:

    ```bash
    helm install edp-argocd-operator epamedp/edp-argocd-operator \
      --version <chart_version> --namespace <edp-project>
    ```

5. Check that the deployment of your EDP Argo CD Operator is running in the <edp-project> namespace.

## Local Development

In order to develop the operator, first set up a local environment. For details, please refer to the [Local Development](https://epam.github.io/edp-install/developer-guide/local-development/) page.

Development versions are also available, please refer to the [snapshot helm chart repository](https://epam.github.io/edp-helm-charts/snapshot/) page.

### Related Articles

* [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/)
