# RBAC in kubernetes

## Generating certificates

### Step-1 : Create a namespace

```kubectl create namespace dev```

### Step-2: Generate a private key

```openssl genrsa -out employee.key 2048```

### Step-3: Generate certificate signing request (CSR) with private key

```openssl req -new -key employee.key -out employee.csr -subj "/CN=employee/O=developers"```

### Step-4: Generate Certificate file from the cluter by using cluster/masternode ca.crt and ca.key file with csr file

```openssl x509 -req -in employee.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out employee.crt -days 500```

### Create credentials, context and user in .kube/config file 

```kubectl config set-credentials employee --client-certificate=/Users/jpalapar/employee/.certs/employee.crt  --client-key=/Users/jpalapar/employee/.certs/employee.key```

```kubectl config set-context employee-context --cluster=minikube --namespace=dev --user=employee```

## Creating role and rolebiding

### Try accessing the cluster with the newly created context

- By default kubernetes works on close to open principle so nothing is accessable unless roles and role binginds are created for the user and namespace

- run the foles in rbac/ directory and try changing the context and get/create/update/delete those resources

```kubectl --context=employee-context run --image jpalaparthi/fileuploader demo```

```kubectl --context=employee-context get pods```

```kubectl --context=employee-context get pods```