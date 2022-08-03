# API Reference

Packages:

- [v1.edp.epam.com/v1alpha1](#v1edpepamcomv1alpha1)

# v1.edp.epam.com/v1alpha1

Resource Types:

- [ArgoApplication](#argoapplication)




## ArgoApplication
<sup><sup>[↩ Parent](#v1edpepamcomv1alpha1 )</sup></sup>






ArgoApplication is the Schema for the argoapplications API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v1.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>ArgoApplication</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspec">spec</a></b></td>
        <td>object</td>
        <td>
          ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationstatus">status</a></b></td>
        <td>object</td>
        <td>
          ApplicationStatus contains status information for the application<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec
<sup><sup>[↩ Parent](#argoapplication)</sup></sup>



ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecdestination">destination</a></b></td>
        <td>object</td>
        <td>
          Destination is a reference to the target Kubernetes server and namespace<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>project</b></td>
        <td>string</td>
        <td>
          Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsource">source</a></b></td>
        <td>object</td>
        <td>
          Source is a reference to the location of the application's manifests or chart<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecignoredifferencesindex">ignoreDifferences</a></b></td>
        <td>[]object</td>
        <td>
          IgnoreDifferences is a list of resources and their fields which should be ignored during comparison<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecinfoindex">info</a></b></td>
        <td>[]object</td>
        <td>
          Info contains a list of information (URLs, email addresses, and plain text) that relates to the application<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>revisionHistoryLimit</b></td>
        <td>integer</td>
        <td>
          RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsyncpolicy">syncPolicy</a></b></td>
        <td>object</td>
        <td>
          SyncPolicy controls when and how a sync will be performed<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.destination
<sup><sup>[↩ Parent](#argoapplicationspec)</sup></sup>



Destination is a reference to the target Kubernetes server and namespace

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is an alternate way of specifying the target cluster by its symbolic name<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>server</b></td>
        <td>string</td>
        <td>
          Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source
<sup><sup>[↩ Parent](#argoapplicationspec)</sup></sup>



Source is a reference to the location of the application's manifests or chart

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>repoURL</b></td>
        <td>string</td>
        <td>
          RepoURL is the URL to the repository (Git or Helm) that contains the application manifests<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>chart</b></td>
        <td>string</td>
        <td>
          Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcedirectory">directory</a></b></td>
        <td>object</td>
        <td>
          Directory holds path/directory specific options<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcehelm">helm</a></b></td>
        <td>object</td>
        <td>
          Helm holds helm specific options<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcekustomize">kustomize</a></b></td>
        <td>object</td>
        <td>
          Kustomize holds kustomize specific options<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>path</b></td>
        <td>string</td>
        <td>
          Path is a directory path within the Git repository, and is only valid for applications sourced from Git.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourceplugin">plugin</a></b></td>
        <td>object</td>
        <td>
          ConfigManagementPlugin holds config management plugin specific options<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>targetRevision</b></td>
        <td>string</td>
        <td>
          TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.directory
<sup><sup>[↩ Parent](#argoapplicationspecsource)</sup></sup>



Directory holds path/directory specific options

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>exclude</b></td>
        <td>string</td>
        <td>
          Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>include</b></td>
        <td>string</td>
        <td>
          Include contains a glob pattern to match paths against that should be explicitly included during manifest generation<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcedirectoryjsonnet">jsonnet</a></b></td>
        <td>object</td>
        <td>
          Jsonnet holds options specific to Jsonnet<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>recurse</b></td>
        <td>boolean</td>
        <td>
          Recurse specifies whether to scan a directory recursively for manifests<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.directory.jsonnet
<sup><sup>[↩ Parent](#argoapplicationspecsourcedirectory)</sup></sup>



Jsonnet holds options specific to Jsonnet

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecsourcedirectoryjsonnetextvarsindex">extVars</a></b></td>
        <td>[]object</td>
        <td>
          ExtVars is a list of Jsonnet External Variables<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>libs</b></td>
        <td>[]string</td>
        <td>
          Additional library search dirs<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcedirectoryjsonnettlasindex">tlas</a></b></td>
        <td>[]object</td>
        <td>
          TLAS is a list of Jsonnet Top-level Arguments<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.directory.jsonnet.extVars[index]
<sup><sup>[↩ Parent](#argoapplicationspecsourcedirectoryjsonnet)</sup></sup>



JsonnetVar represents a variable to be passed to jsonnet during manifest generation

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>code</b></td>
        <td>boolean</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.directory.jsonnet.tlas[index]
<sup><sup>[↩ Parent](#argoapplicationspecsourcedirectoryjsonnet)</sup></sup>



JsonnetVar represents a variable to be passed to jsonnet during manifest generation

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>code</b></td>
        <td>boolean</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.helm
<sup><sup>[↩ Parent](#argoapplicationspecsource)</sup></sup>



Helm holds helm specific options

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecsourcehelmfileparametersindex">fileParameters</a></b></td>
        <td>[]object</td>
        <td>
          FileParameters are file parameters to the helm template<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ignoreMissingValueFiles</b></td>
        <td>boolean</td>
        <td>
          IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsourcehelmparametersindex">parameters</a></b></td>
        <td>[]object</td>
        <td>
          Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>passCredentials</b></td>
        <td>boolean</td>
        <td>
          PassCredentials pass credentials to all domains (Helm's --pass-credentials)<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>releaseName</b></td>
        <td>string</td>
        <td>
          ReleaseName is the Helm release name to use. If omitted it will use the application name<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>skipCrds</b></td>
        <td>boolean</td>
        <td>
          SkipCrds skips custom resource definition installation step (Helm's --skip-crds)<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>valueFiles</b></td>
        <td>[]string</td>
        <td>
          ValuesFiles is a list of Helm value files to use when generating a template<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>string</td>
        <td>
          Values specifies Helm values to be passed to helm template, typically defined as a block<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          Version is the Helm version to use for templating ("3")<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.helm.fileParameters[index]
<sup><sup>[↩ Parent](#argoapplicationspecsourcehelm)</sup></sup>



HelmFileParameter is a file parameter that's passed to helm template during manifest generation

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is the name of the Helm parameter<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>path</b></td>
        <td>string</td>
        <td>
          Path is the path to the file containing the values for the Helm parameter<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.helm.parameters[index]
<sup><sup>[↩ Parent](#argoapplicationspecsourcehelm)</sup></sup>



HelmParameter is a parameter that's passed to helm template during manifest generation

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>forceString</b></td>
        <td>boolean</td>
        <td>
          ForceString determines whether to tell Helm to interpret booleans and numbers as strings<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is the name of the Helm parameter<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          Value is the value for the Helm parameter<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.kustomize
<sup><sup>[↩ Parent](#argoapplicationspecsource)</sup></sup>



Kustomize holds kustomize specific options

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>commonAnnotations</b></td>
        <td>map[string]string</td>
        <td>
          CommonAnnotations is a list of additional annotations to add to rendered manifests<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>commonLabels</b></td>
        <td>map[string]string</td>
        <td>
          CommonLabels is a list of additional labels to add to rendered manifests<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>forceCommonAnnotations</b></td>
        <td>boolean</td>
        <td>
          ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>forceCommonLabels</b></td>
        <td>boolean</td>
        <td>
          ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>images</b></td>
        <td>[]string</td>
        <td>
          Images is a list of Kustomize image override specifications<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namePrefix</b></td>
        <td>string</td>
        <td>
          NamePrefix is a prefix appended to resources for Kustomize apps<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>nameSuffix</b></td>
        <td>string</td>
        <td>
          NameSuffix is a suffix appended to resources for Kustomize apps<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          Version controls which version of Kustomize to use for rendering manifests<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.plugin
<sup><sup>[↩ Parent](#argoapplicationspecsource)</sup></sup>



ConfigManagementPlugin holds config management plugin specific options

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecsourcepluginenvindex">env</a></b></td>
        <td>[]object</td>
        <td>
          Env is a list of environment variable entries<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.source.plugin.env[index]
<sup><sup>[↩ Parent](#argoapplicationspecsourceplugin)</sup></sup>



EnvEntry represents an entry in the application's environment

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is the name of the variable, usually expressed in uppercase<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          Value is the value of the variable<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.ignoreDifferences[index]
<sup><sup>[↩ Parent](#argoapplicationspec)</sup></sup>



ResourceIgnoreDifferences contains resource filter and list of json paths which should be ignored during comparison with live state.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>group</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>jqPathExpressions</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>jsonPointers</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>managedFieldsManagers</b></td>
        <td>[]string</td>
        <td>
          ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.info[index]
<sup><sup>[↩ Parent](#argoapplicationspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.syncPolicy
<sup><sup>[↩ Parent](#argoapplicationspec)</sup></sup>



SyncPolicy controls when and how a sync will be performed

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecsyncpolicyautomated">automated</a></b></td>
        <td>object</td>
        <td>
          Automated will keep an application synced to the target revision<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#argoapplicationspecsyncpolicyretry">retry</a></b></td>
        <td>object</td>
        <td>
          Retry controls failed sync retry behavior<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>syncOptions</b></td>
        <td>[]string</td>
        <td>
          Options allow you to specify whole app sync-options<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.syncPolicy.automated
<sup><sup>[↩ Parent](#argoapplicationspecsyncpolicy)</sup></sup>



Automated will keep an application synced to the target revision

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>allowEmpty</b></td>
        <td>boolean</td>
        <td>
          AllowEmpty allows apps have zero live resources (default: false)<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>prune</b></td>
        <td>boolean</td>
        <td>
          Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>selfHeal</b></td>
        <td>boolean</td>
        <td>
          SelfHeal specifies whether to revert resources back to their desired state upon modification in the cluster (default: false)<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.syncPolicy.retry
<sup><sup>[↩ Parent](#argoapplicationspecsyncpolicy)</sup></sup>



Retry controls failed sync retry behavior

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#argoapplicationspecsyncpolicyretrybackoff">backoff</a></b></td>
        <td>object</td>
        <td>
          Backoff controls how to backoff on subsequent retries of failed syncs<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>limit</b></td>
        <td>integer</td>
        <td>
          Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.spec.syncPolicy.retry.backoff
<sup><sup>[↩ Parent](#argoapplicationspecsyncpolicyretry)</sup></sup>



Backoff controls how to backoff on subsequent retries of failed syncs

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>duration</b></td>
        <td>string</td>
        <td>
          Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. "2m", "1h")<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>factor</b></td>
        <td>integer</td>
        <td>
          Factor is a factor to multiply the base duration after each failed retry<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>maxDuration</b></td>
        <td>string</td>
        <td>
          MaxDuration is the maximum amount of time allowed for the backoff strategy<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### ArgoApplication.status
<sup><sup>[↩ Parent](#argoapplication)</sup></sup>



ApplicationStatus contains status information for the application

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td>
          Error show detailed error message if reconciliation fails<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>
          Status contains the result of reconciliation<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>