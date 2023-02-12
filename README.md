# dracon-community-pipelines
A collection of community contributed Dracon Pipelines look here for inspiration or solutions

## Quickstart

### Prerequisites

- Access to a Kubernetes cluster w/ Tekton Pipelines installed. See below for a local version.

#### Local Kubernetes w/ Tekton Pipelines via [KinD](https://kind.sigs.k8s.io/)

1. Create KinD cluster named `dracon-demo`. For more info, see [official documentation](https://kind.sigs.k8s.io/docs/user/quick-start/#creating-a-cluster):

    ```bash
    $ kind create cluster --name dracon-demo
    ```

2. Install Tekton Pipelines. For more info, see [official documentation](https://tekton.dev/docs/installation/pipelines/#installing-tekton-pipelines-on-kubernetes).

    ```bash
    $ kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
    ```

3. **Optional**: Install Tekton Dashboard for a Web UI. For more info, see [official documentation](https://github.com/tektoncd/dashboard/blob/main/docs/install.md).

    ```bash
    $ curl -sL https://raw.githubusercontent.com/tektoncd/dashboard/main/scripts/release-installer | \
        bash -s -- install latest --read-write
    # Use `kubectl proxy` so you can access Kubernetes services on your local machine.
    $ kubectl proxy
    # Tekton Dashboard is now available at: http://localhost:8001/api/v1/namespaces/tekton-pipelines/services/tekton-dashboard:http/proxy/#/about
    ```

4. **Optional**: Install ECK and create an Elasticsearch + Kibana Dashboards. For more info, see [official documentation](https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-deploy-eck.html).

    ```bash
    # Create ECK CRDs.
    $ kubectl create -f https://download.elastic.co/downloads/eck/2.6.1/crds.yaml
    # Apply ECK operator resources.
    $ kubectl apply -f https://download.elastic.co/downloads/eck/2.6.1/operator.yaml
    # Create Elasticsearch.
    $ cat <<EOF | kubectl apply -f -
    apiVersion: elasticsearch.k8s.elastic.co/v1
    kind: Elasticsearch
    metadata:
    name: quickstart
    spec:
    version: 8.6.1
    nodeSets:
    - name: default
        count: 1
        config:
        node.store.allow_mmap: false
    EOF
    # Create Kibana.
    $ cat <<EOF | kubectl apply -f -
    apiVersion: kibana.k8s.elastic.co/v1
    kind: Kibana
    metadata:
    name: quickstart
    spec:
    version: 8.6.1
    count: 1
    elasticsearchRef:
        name: quickstart
    EOF
    ```

### Running a Community Pipeline

1. Apply the Pipeline resources output via Kustomize to your cluster:

    ```bash
    $ kustomize build pipelines/github.com/kubernetes/kubernetes/ | kubectl apply -f -
    ```

2. Create a PipelineRun

    ```bash
    $ kubectl create -f pipelines/github.com/kubernetes/kubernetes/pipelinerun.yaml
    ```
