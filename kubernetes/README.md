# Backing Services on GCE Kubernetes

## Etcd Cluster

First off, you will want to set up the etcd cluster. You can do so with:

```
kubectl create -f etcd.yml
```

## HA PostgreSQL cluster with Stolon with persistent storage

We use [Stolon](https://github.com/sorintlab/stolon) to provide a HA PostgreSQL
service. The following steps roughly follow the ones in their example README,
but we changed the keeper configuration a little bit to make it use persistent
disk storage.

## Set up the persistent Postgres data disks

First get an instance name from your cluster by running:

```
gcloud compute instances list
```

Then create the disks, and format them with the script provided:

```
bin/create_data_disk
```

This will create 2 disks, one for each stolon-keeper:

  - pg-data-disk-1 (200GB)
  - pg-data-disk-2 (200GB)

## Stolon Cluster setup

These example points to a single node etcd cluster on `10.115.252.12:2379` without tls. You can change the ST${COMPONENT}_STORE_ENDPOINTS environment variables in the definitions to point to the right etcd cluster.

### Create the sentinel(s)

```
kubectl create -f stolon-sentinel.yaml
```

This will create a replication controller with one pod executing the stolon sentinel. You can also increase the number of replicas for stolon sentinels in the rc definition or do it later.

### Create the keeper's password secret

This creates a password secret that can be used by the keeper to set up the initial database user. This example uses the value 'password1' but you will want to replace the value with a Base64-encoded password of your choice.

```
kubectl create -f secret.yaml
```

### Create the stolon keepers

Note: In this example the stolon keeper is a replication controller that, for every pod replica, uses a volume for stolon and PostgreSQL data of from the data claims we created. So it'll **NOT** go away when the related pod is destroyed. Actually (kubernetes 1.0), for working with persistent volumes we _need_ to define a different replication controller with `replicas=1` for every keeper instance. So make sure to keep those replicas to be 1 for the keepers! :)

```
kubectl create -f stolon-keeper-1.yaml
```

This will create a replication controller that will create one pod  executing the stolon keeper.
The first keeper will initialize an empty PostgreSQL instance and the sentinel will elect it as the master.

Once the leader sentinel has elected the first master and created the initial cluster view you can add additional stolon keepers. Will do this later.

### Setup superuser password

Now, you should add a password for the `stolon` user (since it's the os user used by the image for initializing the database). In future this step should be automated.

#### Get the pod id
```
kubectl get pods

NAME                       READY     STATUS    RESTARTS   AGE
stolon-keeper-rc-qpqp9     1/1       Running   0          1m
```

### Create the proxies

```
kubectl create -f stolon-proxy.yaml
```
Also the proxies can be created from the start with multiple replicas.

### Create the proxy service

The proxy service is used as an entry point with a fixed ip and dns name for accessing the proxies.

```
kubectl create -f stolon-proxy-service.yaml
```


### Add another keeper

```
kubectl create -f stolon-keeper-2.yaml
```

you'll have a situation like this:

```
kubectl get pods
NAME                         READY     STATUS    RESTARTS   AGE
stolon-keeper-rc-1-2fvw1     1/1       Running   0          1m
stolon-keeper-rc-2-qpqp9     1/1       Running   0          5m
stolon-proxy-rc-up3x0        1/1       Running   0          5m
stolon-sentinel-rc-9cvxm     1/1       Running   0          5m
```

you should wait some seconds (take a look at the pod's logs) to let the postgresql in the new keeper pod to sync with the master

### Scale your cluster

you can also add additional stolon keepers and also increase/decrease the number of stolon sentinels and proxies:

```
kubectl scale --replicas=2 rc stolon-sentinel-rc
```

```
kubectl scale --replicas=2 rc stolon-proxy-rc
```
