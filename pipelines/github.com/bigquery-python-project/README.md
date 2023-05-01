## Bigquery pipeline

This pipelinerun requires a secret that is best kept in an environment variable.
You can run this pipeline as such


   ```bash
    $ kustomize build | kubectl apply -f -
    # Note: you can just run the below to see the generated Tekton Pipeline resources
    # $ kustomize build
    ```
```bash
export CONSUMER_BIGQUERY_TOKEN=$(gcloud auth application-default print-access-token)
sed -i "s#\\\\$CONSUMER_BIGQUERY_TOKEN#$CONSUMER_BIGQUERY_TOKEN#g" pipelinerun.yaml
kubectl create -f pipelinerun.yaml
```