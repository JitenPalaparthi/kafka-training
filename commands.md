# Kubernetes and Kafka

## Minikube

### Install Minikube

```curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64```

```sudo install minikube-linux-amd64 /usr/local/bin/minikube```


### Workloads

- Pod, Deployment, ReplicationController, ReplicaSet, StatefulSet, Job, CronJob
- There any also some CRDs (Custom Resource Definitions) can also be workloads
- One of the examples is Strimzi (Operator)

- Everything is temp. Pods come and go (Mey be lived for short time) Epheimeral.
- The smallest workload in kubernetes is Pod
- Ultimately any workload that is created is a Pod

- Few useful concepts and tools
- K9s --> Wonderful tool for management and operations
- k3s -> Wonderful innovation by Rancher
- Cluster API --> Started by vmware and being used or maintaind by big gaints (Microsoft, Google and Amazon)
- Kyvarno --> Policy management on Kubernetes (Open source project by Nirmata)
- CNI --> Container Netwroking Interface
- CSI --> Container Storage Interface
- CRI --> Container Runtime Interface
  
- Kubernetes Architecture
- Created a Multi node Cluster on Minikube 
- Workloads
  - Pod
  - ReplicaSet
  - Deployment
    - Resource allocations
    - Deploying on specific nodes
  - DaemonSet
  - Job
  - Scale
  - Rollback
    - RollingUpdate
    - Recreate
  
  - Services
    - Loadbalancer
    - NodePort
    - ClusterIP 

- CronJob
- Deeper into Services part
- StatefulSet
- ConfigMap
- Secrets
- Security --> Security in Kafka
- Manually deploy Kafka Cluster on machine
- Deploy Kafka on Kubernetes using scripts provided by commnity and also we try to deploy our own 
- How does Kubernetes operator work/ will create a simple kubernetes operator --> Gives you clear idea about operator and if there is any issue in Strimzi you can be able to debug it.
- Strimzi kafka operator basics 
- Security concepts using Strimzi 

### Create Configmap using command by passing key-values

```kubectl create configmap kafka-config -n test --from-literal deployment-type=test --from-literal deployment-version=v0.0.1```

### Create Configmap using command by passing filenames

```kubectl create cm kafka-config-2 -n test --from-file deployment-type.txt --from-file deployment-version.txt```

### Create namespace thru command

```kubectl create ns test```

### To create a secret from command from literals

```kubectl create secret generic db-creds -n test --from-literal=username=jiten --from-literal=password=jiten@123```

### To create a secret from command with files

```ubectl create secret generic db-creds-file  --from-file=username --from-file=password -n test```