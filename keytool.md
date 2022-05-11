#### To Create a certificate

keytool --genkeypair -alias firstcertificate -keypass jiten@123 -validity 30 -storepass jiten@123456

#### All keytool commands are in one file

https://dzone.com/articles/the-most-common-java-keytool-keystore-commands


kubectl create namespace dev

openssl genrsa -out employee.key 2048

openssl req -new -key employee.key -out employee.csr -subj "/CN=employee/O=developers"

openssl x509 -req -in employee.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out employee.crt -days 500

#### create a directory to keep the keys

kubectl config set-credentials employee --client-certificate=/Users/jpalapar/employee/.certs/employee.crt  --client-key=/Users/jpalapar/employee/.certs/employee.key
kubectl config set-context employee-context --cluster=minikube --namespace=dev --user=employee

kubectl --context=employee-context get pods
// Kubernetes always has a concept call close to open

# test after running role and role bind

kubectl --context=employee-context run --image jpalaparthi/fileuploader demo
kubectl --context=employee-context get pods