# Kubernetes Operator Sample

## The Goal

If you're using .NET Core applications which are using ConfigMap in Kubernetes, you must be having a problem with the reflection of configuration changes on runtime.

There are some issues such as [this](https://github.com/dotnet/corefx/issues/27232) on Github about it and they're still open.

The operator will create a **ConfigMap** and it knows which deployments are using this. If there are any changes on ConfigMap, it will update whole linked deployments.

## Create demo environment on Minikube

```bash
# Create minikube
minikube start
# Change LoadBalancerExternalIP for the sample project
minikube ip | xargs -I {} sed -i "" 's|YOUR_MINIKUBE_IP|{}|g' samples/manifest.yaml
#(OPTIONAL) If you're using private container registry, you need to create a secret contains registry credentials.
kubectl create secret docker-registry reg-cred --docker-server=repo.treescale.com --docker-username=[DOCKER_REGISTRY_USER_NAME] --docker-password=[DOCKER_REGISTRY_PASSWORD] --docker-email=[INFORMATION_MAIL]
```

## Build and deploy the operator

```bash
# Build and push your operator
operator-sdk build repo.treescale.com/selcukusta/config-management-operator:1.0.0
docker push repo.treescale.com/selcukusta/config-management-operator:1.0.0

# Change build image for operator deployment
sed -i "" 's|REPLACE_IMAGE|repo.treescale.com/selcukusta/config-management-operator:1.0.0|g' deploy/operator.yaml

# Setup Service Account
kubectl apply -f deploy/service_account.yaml
# Setup RBAC
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/role_binding.yaml
# Setup the CRD
kubectl apply -f deploy/crds/selcukusta.com_netcoreconfigmanagements_crd.yaml
# Deploy the app-operator
kubectl apply -f deploy/operator.yaml

# Create your own Custom Resource
kubectl apply -f deploy/crds/selcukusta.com_v1alpha1_netcoreconfigmanagement_cr.yaml

# Check if ConfigMap is created
kubectl describe cm sampleconfig
##### Expected Output #####
# Data
# ====
# appsettings.Kubernetes.json:
# ----
# {
#   "ConfigMapOptions": {
#       "Hello": "World!"
#   }
# }
##### Expected Output #####
```

## Deploy sample application

```bash
# Sample .NET Core 3.0 application is read static
# config file and mounted config file from ConfigMap.
# When the ConfigMap is changed, deployment should be
# updated via Custom Resource.
kubectl apply -f samples/manifest.yaml
```

## Test

```bash
# Request the application per second
while true; do sleep 1; curl http://YOUR_MINIKUBE_IP:30002/api/values;echo -e "\t$(date)";done
# Initial output should be;
# ["sampleapp-dep-56f8f597c5-v5fsx","Neptune!","1.0.0"]	Thu Nov 28 12:22:15 +03 2019
```

Open `deploy/crds/selcukusta.com_v1alpha1_netcoreconfigmanagement_cr.yaml` file and change the `Neptune!` as `World!`. Then re-run `kubectl apply -f deploy/crds/selcukusta.com_v1alpha1_netcoreconfigmanagement_cr.yaml`.

Turn to your terminal and see the changes on-the-fly!

```json
["sampleapp-dep-55b48c48b9-dz6qh", "World!", "1.0.0"]	Thu Nov 28 12:24:12 +03 2019
```

Thanks to `RollingUpdate`, there is any downtime while configuration changes are applied.

## References

> https://github.com/kubernetes/kubectl/blob/3874cf79897cfe1e070e592391792658c44b78d4/pkg/polymorphichelpers/objectrestarter.go

> https://github.com/operator-framework/operator-sdk/blob/master/doc/user/client.md

> https://medium.com/@shubhomoybiswas/writing-kubernetes-operator-using-operator-sdk-c2e7f845163a

> https://github.com/operator-framework/operator-sdk/blob/master/example/memcached-operator/memcached_controller.go.tmpl
