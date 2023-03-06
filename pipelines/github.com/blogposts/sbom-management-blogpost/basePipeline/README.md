# Base SBOM Pipeline Template

The files in this directory setup and execute a Dracon pipeline which runs [aquasec/trivy](github.com/aquasec/trivy) against a target repository, enriches the results with licensing information from Google's Open Source Insights Repository and finally, uploads the results in Dependency Track.

You can install the pipeline with
`kubectl -n dracon apply -k .`
Which will apply kustomization to all the objects described in the `kustomization.yaml` file and install the pipeline in the namespace "dracon" in the k8s cluster described by your default context.

The `pipelinerun.yaml.tmpl` file is the suggested template to provide to the `runpipelines` project for generating pipelineruns.

