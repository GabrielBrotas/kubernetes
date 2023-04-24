```bash

helm create first-chart # create chart, by default is going to be a nginx

```
Chart.yaml -> metadata about the chart
charts/ -> dependencies, if this charts depend on other chart it will be pulled and stored to this folder
templates/ -> k8s manifest files to deploy on k8s cluster
values -> values that goes into template folder, we can override the values

templates/_helper.tpl -> helper functions to use across others files
```bash
# deploy our chart
helm install mychart first-chart # the dry run flag will validate the template generated
```
---

Package Chart
```bash
helm package first-chart 
```

Verify indentation errors
```bash
helm lint first-chart
```

Verify chart template output
```bash
helm template first-chart
```

Get manifest deployed
```bash
helm get manifest mychart
```

Install chart dependencies, the dependencies are installed on Chart.yaml
```bash
helm dependency update first-chart

# verify if is setup properly
helm upgrade mychart first-chart/ --dry-run # the dry run flag will validate the template generated, 
helm upgrade mychart first-chart      
```

Hooks -> we can execute hooks after some event on the Chart, after/before create, delete, upgrade,...

Test
test is always against the release, not the Chart itself
```bash
helm test mychart
```