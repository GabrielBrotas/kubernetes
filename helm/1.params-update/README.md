
```bash
helm install mydb bitnami/mysql --set auth.rootPassword = test123456 # not recommended

helm install mydb bitnami/mysql --values values.yaml

helm upgrade mydb bitnami/mysql --values

helm upgrade mydb bitnami/mysql --reuse-values

```