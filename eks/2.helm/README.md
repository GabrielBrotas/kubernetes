

```bash
# search charts
helm search hub <package name> from hub
helm search repo <package name> local

helm repo add bitnami https://charts.bitnami.com/bitnami # add bitnami package as metadata
helm search repo nginx

cd eks/2.helm
helm pull bitnami/nginx --untar=true

# install
helm install nginx-app bitnami/nginx
```