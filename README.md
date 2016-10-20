# Deis Backing Services Manager

[![](https://quay.io/repository/codaisseur/deis-backing-services-api/status)]https://quay.io/repository/codaisseur/deis-backing-services-api)

At [Codaisseur](https://www.codaisseur.com) we wanted to provide a Heroku like environment for students. We <3 [Deis Workflow](https://deis.com/) as do our students, as you can see in this picture.

[![](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)

## Anyway..

We use Rails a lot, and PostgreSQL, so we needed an easy way for students to set up their apps despite the fact that Deis does not provide a PostgreSQL service for
them like Heroku does. This manager app takes care of that.

## Prerequisites

  - A working Deis workflow cluster on k8s
  - Helm classic installed and set up
  - Deis client installed and set up

## Install a Redis cluster with Helm

We'll use [Helm](http://helm.sh/) to set up a HA Redis cluster for us.
Refer to the Helm docs for install instructions.

```
helmc install redis-cluster
```

## Set up the Postgres service cluster and MongoDB data disk

Check out the [instructions](/kubernetes) to set up the services cluster on Kubernetes to get you started.

## Install MongoDB

```
helmc repo add bitnami https://github.com/bitnami/charts
helmc fetch bitnami/mongodb
helmc edit mongodb
```

Update `tpl/values.toml` to match your needs.

Then run the generators:

```
helmc generate --force mongodb
```

And install:

```
helmc install mongodb
```

## Installing Memcached

Fetch the `memcached` chart with `helmc`

```
helmc fetch memcached
```

Then make it a little easier to use by adding 2 services

```
helmc edit memcached
```

```yaml
# manifests/memcached-rc.yaml

apiVersion: v1
kind: Service
metadata:
  name: memcached-1
spec:
  ports:
  - name: memcached-port-1
    port: 11211
    protocol: TCP
    targetPort: 11211
  selector:
    app: memcached-1

---

apiVersion: v1
kind: Service
metadata:
  name: memcached-2
spec:
  ports:
  - name: memcached-port-2
    port: 11211
    protocol: TCP
    targetPort: 11211
  selector:
    app: memcached-2

---

apiVersion: v1
kind: ReplicationController
metadata:
  name: memcached-1
  labels:
    app: memcached-1
    heritage: helm
spec:
  replicas: 1
  selector:
    name: memcached-1
    mode: cluster
    provider: memcached
  template:
    metadata:
      labels:
        name: memcached-1
        mode: cluster
        provider: memcached
    spec:
      containers:
      - name: memcached
        image: "memcached:1.4.24"
        ports:
        - containerPort: 11211
---
apiVersion: v1
kind: ReplicationController
metadata:
  name: memcached-2
  labels:
    heritage: helm
    app: memcached-2
spec:
  replicas: 1
  selector:
    name: memcached-2
    mode: cluster
    provider: memcached
  template:
    metadata:
      labels:
        name: memcached-2
        mode: cluster
        provider: memcached
    spec:
      containers:
      - name: memcached
        image: "memcached:1.4.24"
        ports:
        - containerPort: 11211
```

Then install it

```
helmc install memcached
```

---
apiVersion: v1

## Setting up the Manager App

### Create a Deis app

Clone this repository and set up a deis app:

```
deis create
```

### Get the proxy service ip

```
kubectl get svc

NAME                   CLUSTER-IP       EXTERNAL-IP   PORT(S)
memcached-1            10.115.249.176   <none>        11211/TCP
memcached-2            10.115.244.86    <none>        11211/TCP
mongodb                10.115.240.18    <none>        27017/TCP
redis-sentinel         10.115.253.131   <none>        26379/TCP
stolon-proxy-service   10.115.246.38    <none>        5432/TCP
```

### Deploy the Backing Services Manager App

The app runs from a Docker container, and you will need to set a
`SECRET_KEY_BASE`:

```
deis config:set SECRET_KEY_BASE=$(rails secret)
```

Use the IP to create the `DATABASE_URL`, we will use the one mentioned in the above output:

```
deis config:set DATABASE_URL=postgresql://stolon:password@10.115.246.38:5432/backing_services
```

And same for `REDIS_URL`:

```
deis config:set REDIS_URL=redis://10.115.253.131:26379/redis_services
```

Then the 2 `MEMCACHED_SERVERS` separated by commas:

```
deis config:set MEMCACHED_SERVERS=10.115.249.176,10.115.244.86
```

And finally `MONGODB_URL`:

```
deis config:set MONGODB_URL=mongodb://root:rootPassword@10.115.240.18:27017/admin
```

Use the root user and root password from the `values.toml` file that you edited with `helmc`.

Then deploy the Rails app:

```
git push deis master
```

And run the rake task to create the `backing_services` database for this app:

```
deis run rake deis:create_database
```

Then migrate the database:

```
deis run rake db:migrate
```

## Get a Postgres database for new Apps

When you are setting up a new app that needs to use PostgreSQL, you can create a database by `POST`ing to the Manager app's Postgres endpoint:

```
curl -XPOST http://watery-fowls.xxx.xxx.xx.xxx.nip.io/postgres_databases
```

This will return your new `DATABASE_URL`:

```
DATABASE_URL=postgres://meagan:itXA7CiKj33R7T4cS8s4@10.xxx.xxx.xx:5432/compress_program
```

## Get a Redis db for new Apps

Similarly, we can get a REDIS_URL for new apps:

```
curl -XPOST http://watery-fowls.xxx.xxx.xx.xxx.nip.io/redis_services
```

Which will return something like:

```
REDIS_URL=redis://10.xxx.xxx.xx:26379/copy_port
```

## Get a Mongo db for new Apps

And unsurprisingly this works the same for MongoDB:

```
curl -XPOST http://watery-fowls.xxx.xxx.xx.xxx.nip.io/mongodb_services
```

Which will return something like:

```
MONGODB_URL=mongodb://10.xxx.xxx.xx:27017/persistence_much
```

## Get a Memcached namespace for new Apps

Memcached is configured with servers and a namespace:

```
curl -XPOST http://watery-fowls.xxx.xxx.xx.xxx.nip.io/memcached_services
```

Which will return something like:

```
MEMCACHED_SERVERS=10.115.244.86,10.115.249.176 MEMCACHED_NAMESPACE=navigate_card
```

## Roadmap

  - We will add more services as we go, like:
    - √ ~~PostgreSQL~~
    - √ ~~Redis~~
    - √ ~~MongoDB~~
    - √ ~~Memcached~~
  - Create a Helm chart for Stolon.
  - Create Helm charts for the entire cluster

Let us know which services you are missing and we will try to add them.

Feel free to help us out or leave any feedback in the issues :)

## Changelog

  - **2016-07-26** Initial project with PostgreSQL service
  - **2016-07-31** Added Redis, MongoDB, and Memcached services
  - **2016-08-08** Fixed MongoDB issues, running from Dockerfile now

## Thanks to

  - The [Deis](https://deis.com/) Team for the awesomeness that is our own PaaS!
  - The Bitnami team for their awesome list of [helmc charts](https://github.com/bitnami/charts)
  - The Storint.lab team for their super duper HA Postgres solution [Stolon](https://github.com/sorintlab/stolon)
