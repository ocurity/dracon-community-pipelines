# Run Pipelines

This directory contains a small go program which when given the `results.txt` output from "createDTProjects" and the location of the file `pipelinerun.yaml.tmpl` generates one pipelinerun.yaml per entry in the "results.txt" file.

You can then execute all these pipelineruns as such:

```bash

 for run in $(ls pipelineruns/); do
  kubectl -n dracon create -f pipelineruns/$run; done
```

This will create a large number of pipeline runs in your cluster so be careful if you are running this on a limited resources environment.