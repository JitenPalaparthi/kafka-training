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