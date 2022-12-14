# edp-argocd-operator

![Version: 0.3.0](https://img.shields.io/badge/Version-0.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.3.0](https://img.shields.io/badge/AppVersion-0.3.0-informational?style=flat-square)

A Helm chart for EDP Argo CD Operator

**Homepage:** <https://epam.github.io/edp-install/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| epmd-edp | <SupportEPMD-EDP@epam.com> | <https://solutionshub.epam.com/solution/epam-delivery-platform> |
| sergk |  | <https://github.com/SergK> |

## Source Code

* <https://github.com/epam/edp-argocd-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| externalSecrets.enabled | bool | `false` | Enable External Secret Operator usage |
| externalSecrets.secretStoreName | string | `"aws-parameterstore"` | edp-install chart by default provision SecretStore with name `aws-parameterstore` |
| externalSecrets.ssmParameterStoreName | string | `"/edp/deploy-secrets"` | Value name in AWS ParameterStore |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"epamedp/edp-argocd-operator"` | EDP argocd-operator Docker image name. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/edp-argocd-operator) |
| image.tag | string | `""` | EDP argocd-operator Docker image tag. The released image can be found on [Dockerhub](https://hub.docker.com/r/epamedp/edp-argocd-operator/tags) |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext.runAsNonRoot | bool | `true` |  |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"128Mi"` |  |
| resources.requests.cpu | string | `"10m"` |  |
| resources.requests.memory | string | `"64Mi"` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |

