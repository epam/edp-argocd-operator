<a name="unreleased"></a>
## [Unreleased]


<a name="v0.2.0"></a>
## [v0.2.0] - 2022-08-26
### Features

- Populate operator metadata in logs during operator start [EPMDEDP-10403](https://jiraeu.epam.com/browse/EPMDEDP-10403)

### Bug Fixes

- Command `make manifest` should generate Role instead of ClusterRole [EPMDEDP-10408](https://jiraeu.epam.com/browse/EPMDEDP-10408)

### Code Refactoring

- Use repository and tag for image reference in chart [EPMDEDP-10389](https://jiraeu.epam.com/browse/EPMDEDP-10389)

### Routine

- Add CHANGELOG.md file [EPMDEDP-10337](https://jiraeu.epam.com/browse/EPMDEDP-10337)
- Add gcflags for go build artifact [EPMDEDP-10411](https://jiraeu.epam.com/browse/EPMDEDP-10411)

### Documentation

- Align README.md [EPMDEDP-10274](https://jiraeu.epam.com/browse/EPMDEDP-10274)


<a name="v0.1.1"></a>
## [v0.1.1] - 2022-08-08
### Routine

- Update helm-docs version [EPMDEDP-10280](https://jiraeu.epam.com/browse/EPMDEDP-10280)
- Change 'go get' to 'go install' for git-chglog [EPMDEDP-10337](https://jiraeu.epam.com/browse/EPMDEDP-10337)

### Documentation

- Review documentation in README.md [EPMDEDP-10388](https://jiraeu.epam.com/browse/EPMDEDP-10388)


<a name="v0.1.0"></a>
## v0.1.0 - 2022-08-08
### Features

- Add capabilities to provision secrets using ESO [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Allow to re-define VERSION for build [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Align to EDP deployment approach [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Argo CD operator [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)

### Bug Fixes

- Rename secret key for external secret [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Use image.version for helm chart deployment [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)

### Routine

- Add kustomize installation as a part of Makefile [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Fix ArtifactHub image annotation [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)

### Documentation

- Add configuration examples [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)
- Update README.md [EPMDEDP-10225](https://jiraeu.epam.com/browse/EPMDEDP-10225)


[Unreleased]: https://github.com/epam/edp-argocd-operator/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/epam/edp-argocd-operator/compare/v0.1.1...v0.2.0
[v0.1.1]: https://github.com/epam/edp-argocd-operator/compare/v0.1.0...v0.1.1
