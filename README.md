[![codecov](https://codecov.io/gh/epam/edp-argocd-operator/branch/master/graph/badge.svg?token=VNMBQGDKJ3)](https://codecov.io/gh/epam/edp-argocd-operator)

# EDP Argo CD Operator

| :heavy_exclamation_mark: Please refer to [EDP documentation](https://epam.github.io/edp-install/) to get the notion of the main concepts and guidelines. |
| --- |

Get acquainted with the EDP ArgoCD Operator and the installation process as well as the local development, and architecture scheme.

## Overview

EDP Argo CD Operator is an EDP operator that is responsible for managing ArgoCD EDP Tenants. See below diagram for details:

* Argo CD is deployed in separate `argocd` namespace
* Argo CD uses `cluster-admin` role for managing cluster resources
* `control-plane` Application is created using AppsOfApps approach and its code is managed by `control-plane` members
* `control-plane` is used to onboard new Argo CD Tenants (AppProjects)
* `control-plane admin` provides `JWT Token` for each `EDP Tenant`
* `EDP Tenant` deploys `edp-argocd-operator` in its EDP namespace `edpTenant` and uses `JWT Token` provided by `control-plane admin`
* `EDP Tenant Member` manages `Argo CD Repositories` and `Argo CD Applications` using `kind: Secret` and `kind: ArgoApplication` in `edpTenant` namespace

![edpTenant](docs/assets/edpTenant.png)

## Deployment example

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

EDP Argo CD Application has the same specification as native Argo CD one

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

1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed
2. Admin access to the EDP Namespace
3. Access to Argo CD (and JWT Token)
4. EDP project/namespace is deployed by following the [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/) instruction

## Installation

In order to install the Operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":

    ```bash
    helm repo add epmdedp https://epam.github.io/edp-helm-charts/stable
    ```

2. Choose available Helm chart version:

    ```bash
    helm search repo epmdedp/edp-argocd-operator ---devel

    NAME                          CHART VERSION   APP VERSION       DESCRIPTION
    epmdedp/edp-argocd-operator   0.1.0  	      0.1.0             A Helm chart for EDP Argo CD Operator
    ```

    _**NOTE:** It is highly recommended to use the latest released version._

3. Deploy operator:

    Chart parameters available in [deploy-templates/README.md](deploy-templates/README.md):

4. Install operator in the <edp_cicd_project> namespace with the helm command; find below the installation command example:

    ```bash
    helm install edp-argocd-operator epamedp/edp-argocd-operator \
      --version <chart_version> --namespace <edp_cicd_project>
    ```

5. Check the <edp_cicd_project> namespace that should contain operator deployment with your operator in a running status.

## Local Development

In order to develop the operator, first set up a local environment. For details, please refer to the [Developer Guide](https://epam.github.io/edp-install/developer-guide/local-development/) page.

### Related Articles

* [Architecture Scheme of Sonar Operator](docs/arch.md)
* [Install EDP](https://epam.github.io/edp-install/operator-guide/install-edp/)
