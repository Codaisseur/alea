# Deis Backing Services Manager

[![](https://quay.io/repository/codaisseur/deis-backing-services-api/status)]https://quay.io/repository/codaisseur/deis-backing-services-api)

At [Codaisseur](https://www.codaisseur.com) we wanted to provide a Heroku like environment for students. We <3 [Deis Workflow](https://deis.com/) as do our students, as you can see in this picture.

[![](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)

## Anyway..

We use Rails a lot, and PostgreSQL, so we needed an easy way for students to set up their apps despite the fact that Deis does not provide a PostgreSQL service for
them like Heroku does. This manager app takes care of that.

It also has a few other of our favorite services:

  - Redis
  - MongoDB
  - Memcached

## Prerequisites

  - A working Deis workflow cluster on k8s
  - [Helm][helm] installed and set up
  - Deis client installed and set up

## Step-by-Step Guide

  - Create a file called `settings.yaml` and set variables you would like to
    change. See **Supported Values** section below.
  - Generate a `SECRET_KEY_BASE` token for the Rails Services API: `docker run --rm quay.io/codaisseur/deis-backing-services-api rails secret` and put it in your
    `settings.yaml`:

```yaml
# settings.yaml

api:
  secretKeyBase: "ae20e..."
```

## Supported Values

### Storage (`storage`)

Key | Default Value | Description
--- | ------------- | -----------
`standardClassName` | `slow` | Storage class name for standard class persistent storage used by Redis by default.
`ssdClassName` | `fast` | Storage class name for SSD class persistent storage (fast) used by PostgreSQL and MongoDB by default.
`provisioner` | `kubernetes.io/gce-pd` | Storage class provisioner.
`standardType` | `pd-standard` | Standard class persistent storage type.
`ssdType` | `pd-ssd` | SSD class persistent storage type.
`zone` | `europe-west1-b` | Default availability zone in which the services will be deployed. **This should match your container cluster's zone, and it should be the same as where Deis runs.**

### MongoDB (`mongo`)

Key | Default Value | Description
--- | ------------- | -----------
`diskName` | `mongodb-data-disk` | Name of the mongodb data disk (should be unique per cluster, thus configurable).
`storageClassName` | `fast` | The class name of the storage type to use (`fast` or `slow`, see `storage.ssdClassName` and `storage.standardClassName`).
`diskSize` | `500Gi` | Disk size for the mongodb disk.
`imageTag` | `3.2.9-r2` | Image tag to use for the bitnami/mongodb docker image.
`imagePullPolicy` | `IfNotPresent` | Image pull policy for the bitnami/docker image.
`dbRootPassword` | `"root"` | Password for the mongodb `root` user.
`dbUsername` | `"api"` | Mongodb username for the app user.
`dbPassword` | `"mypass"` | Mongodb password for the app user.
`dbDatabase` | `"backing-services-api"` | Name of the database to use for the app.

### PostgreSQL (`postgres`)

Key | Default Value | Description
--- | ------------- | -----------
`imageTag` | `latest` | Image tag for the paunin/postgresql-cluster-pgpool and paunin/postgresql-cluster-pgsql images.
`imagePullPolicy` | `IfNotPresent` | Pull policy for above images.
`diskName` | `pg-data-disk` | Name of postgres primary's data disk (should be unique per cluster, thus configurable).
`storageClassName` | `fast` | The class name of the storage type to use (`fast` or `slow`, see `storage.ssdClassName` and `storage.standardClassName`).
`diskSize` | `200Gi` | Disk size for postgres primary's disk.
`username` | `postgres` | Username for the postgres (root) user.
`password` | `password` | Password for the postgres (root) user.
`database` | `backing_services` | Name of the API's database. The API uses this database to store the services it gave out to users of your cluster.

### Redis (`redis`)

Key | Default Value | Description
--- | ------------- | -----------
`diskName` | `redis-data-disk` | Name of the redis data disk (should be unique per cluster, thus configurable).
`storageClassName` | `slow` | The class name of the storage type to use (`fast` or `slow`, see `storage.ssdClassName` and `storage.standardClassName`).
`diskSize` | `200Gi` | Disk size of the Redis persistent disk.

### API (`api`)

Key | Default Value | Description
--- | ------------- | -----------
`imageTag` | `v0.1.0` | Tag for the quay.io/codaisseur/deis-backing-services-api image.
`imagePullPolicy` | `"Always"` | Pull policy for the quay.io/codaisseur/deis-backing-services-api image.
`secretKeyBase` | `""` | **Create a secret by running: `docker run --rm quay.io/codaisseur/deis-backing-services-api rails secret`**

## SSL for the API Ingress

```bash
$ cat certificate-file.crt
-----BEGIN CERTIFICATE-----
/ * your SSL certificate here */
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
/* any intermediate certificates */
-----END CERTIFICATE-----
$ cat certificate-file.crt | base64 -e
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi8gKiB5b3VyIFNTTCBjZXJ0aWZpY2F0ZSBoZXJlICovCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi8qIGFueSBpbnRlcm1lZGlhdGUgY2VydGlmaWNhdGVzICovCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
$ cat certificate.key
-----BEGIN RSA PRIVATE KEY-----
/* your unencrypted private key here */
-----END RSA PRIVATE KEY-----
$ cat certificate.key | base64 -e
LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQovKiB5b3VyIHVuZW5jcnlwdGVkIHByaXZhdGUga2V5IGhlcmUgKi8KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K
```

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
  - √ ~~Create Helm charts for the entire cluster~~

Let us know which services you are missing and we will try to add them.

Feel free to help us out or leave any feedback in the issues :)

## Changelog

  - **2016-07-26** Initial project with PostgreSQL service
  - **2016-07-31** Added Redis, MongoDB, and Memcached services
  - **2016-08-08** Fixed MongoDB issues, running from Dockerfile now
  - **2016-10-24** Moved away from Stolon and to a setup by @paunin with pgpool2
  - **2016-10-24** Moved away from Helm Classic and to the new Helm

## Thanks to

  - The [Deis](https://deis.com/) Team for the awesomeness that is our own PaaS!
  - The Bitnami team for their awesome list of [helm charts](https://github.com/bitnami/charts)
  - @paunin for his [postgresql cluster setup](https://github.com/paunin/postgres-docker-cluster)

[helm]: https://github.com/kubernetes/helm
