# kubernetes
Orquestrador de container.

Produto Open Source utilizado para automatizar a implantação, o dimensionamento e o gerenciamento de aplicativos em contâiners;

Google é a empresa que criou o Kubernetes por isso é o melhor cloud provider para colocarmos nossos serviços.

**Pontos Importantes:**

- K8S é disponibilizado através de um conjunto de APIs;
- Normalmente acessamos a API usando a CLI: **kubectl;**
- Tudo é baseado em estado. Você configura o estado de cada objeto;
- Kubernetes Master: Controla todo o processo do que os outros nós vão fazer;
    - Kube-apiserver
    - Kube-controller-manager
    - Kube-scheduler
- Outros Nodes:
    - Kubelet
    - Kubeproxy
    

### **Dinâmica “superficial”**

**Cluster:** Conjunto de máquinas (Nodes - Nó)

Cada máquina possui uma quantidade de vCPU e Memória

Podemos ter vários nós.

**Pods:** Unidade que contém os containers provisionados

O Pod representa os processos rodando no cluster. Um pode geralmente representa um container porém podemos ter mais de um container rodando no mesmo pod (apesar de não ser muito comúm)

**Deployment:** Tem o objetivo de provisionar os pods, vai definir quantas replicas queremos de cada Pod

Ex: B = Backend ⇒ 3 réplicas
      F = Frontend ⇒ 2 réplicas

Se um pod cair o K8S vai criar outro automáticamente

![image](https://user-images.githubusercontent.com/63565773/166736051-4a81a184-2ca5-4020-bc71-94e6417e981a.png)

Se pedirmos para adicionar mais um Pod de Frontend no cluster e ele não tiver mais recurso a task vai ficar pendente até ter mais recurso computacional ou criarmos um outro nó(cluster) para jogar essa task no próximo nó disponível.

o k8s fica monitorando a saúde dos cluster e pods para porder recriar caso não estejam mais disponíveis

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

**Mudança de contexto e extensão VSC**

```bash
# verificar clusters disponíveis em todos os ambientes (inclusive cloud (digital ocean, aws, gcp)
kubectl config get-clusters

##kubectl config get-clusters
##NAME
##kind-multipleclusters

# mudar context
kubectl config use-context <context name>
```

Para facilidar isso podemos instalar a extensão do VSC: “kubernetes”, essa extensão vai criar uma aba no nosso VSC para verificar nossos clusters e podermos alterar de forma mais fácil
