
### Utilizando Kind

Kind é uma ferramenta que instala o k8s localmente através de containers docker.

**Comandos Básicos**

```bash
# criar um novo cluster, vai criar uma 'maquina'
kind create cluster 

# credenciais do cluster
cat ~/.kube/config

# se conectar com o cluster do kind
kubctl cluster-info --context kind-kind 

# verificar container do kind rodando
docker ps
kubectl get nodes # listar os nodes do cluster k8s

# pegar todos os nodes criados utilizando o kind
kind get clusters

# deletar cluster
kind delete clusters kind # por padrao o nome do cluster é definido como 'kind'

```

**Kind com múltiplos nós**

kind.yaml

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4

nodes:
# papel no node
- role: control-plane # vai gerenciar o k8s, master
- role: worker
- role: worker
- role: worker
```

```bash
kind create cluster --config=kind.yaml --name=multipleclusters

###
❯ kubectl get nodes
NAME                             STATUS     ROLES                  AGE   VERSION
multipleclusters-control-plane   Ready      control-plane,master   46s   v1.23.4
multipleclusters-worker          Ready      <none>                 25s   v1.23.4
multipleclusters-worker2         Ready      <none>                 25s   v1.23.4
multipleclusters-worker3         NotReady   <none>                 13s   v1.23.4
###
```
