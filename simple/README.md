# go-simple

## Deployment
```
$ kubectl apply -f go-simple-deployment.yaml
$ kubectl port-forward deployment/go-simple 9000:9000
```
localhost:9000 にアクセスする。

## Service
```
$ kubectl apply -f go-simple-service.yaml
$ kubectl get pods,services

NAME                             READY   STATUS    RESTARTS   AGE
pod/go-simple-84d7489bfb-jfhxx   1/1     Running   0          9m39s
pod/go-simple-84d7489bfb-nsmkc   1/1     Running   0          9m39s

NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
service/go-simple    NodePort    10.106.109.180   <none>        80:32582/TCP   5s
service/kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        46d
```
localshot:32582 にアクセスする。
