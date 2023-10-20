kind create cluster --name dracon-demo --quiet
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
kubectl apply -f https://raw.githubusercontent.com/ocurity/dracon-community-pipelines/main/resources/deduplication-enricher-db.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/2.6.1/crds.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/2.6.1/operator.yaml
kubectl apply -f https://raw.githubusercontent.com/ocurity/dracon-community-pipelines/main/resources/eck-elasticsearch.yaml
kubectl apply -f https://raw.githubusercontent.com/ocurity/dracon-community-pipelines/main/resources/eck-kibana.yaml
