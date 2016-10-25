# Alea - Deis Backing Services Manager

[![](https://quay.io/repository/codaisseur/alea-controller/status)]https://quay.io/repository/codaisseur/alea-controller)

At [Codaisseur](https://www.codaisseur.com) we want to provide a Heroku like environment for students. We <3 [Deis Workflow](https://deis.com/) as do our students, as you can see in this picture.

[![](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)

## Anyway..

We use Rails a lot, and PostgreSQL, so we needed an easy way for students to set up their apps despite the fact that Deis does not provide a PostgreSQL service for them like Heroku does. This manager app takes care of that.

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
  - Generate a `SECRET_KEY_BASE` token for the Rails Services API: `docker run --rm quay.io/codaisseur/alea-controller rails secret` and put it in your
    `settings.yaml`:
  - Set up SSL for the Controller Ingress (see below)

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

### API (`controller`)

Key | Default Value | Description
--- | ------------- | -----------
`imageTag` | `v0.1.0` | Tag for the quay.io/codaisseur/alea-controller image.
`imagePullPolicy` | `"Always"` | Pull policy for the quay.io/codaisseur/alea-controller image.
`secretKeyBase` | `""` | **Create a secret by running: `docker run --rm quay.io/codaisseur/alea-controller rails secret`**

## Setting up SSL for the Controller Ingress

The Alea Controller Ingress needs an SSL certificate. To set this up, create a yaml file, `controler-ssl.yaml`, and put in the following:

```yaml
# controler-ssl.yaml

apiVersion: v1
kind: Secret
metadata:
  name: controller-ssl-cert
  namespace: services
type: Opaque
data:
  tls.crt: LS0tLS1CR...
  tls.key: LS0tLS1CR...
```

Put in your crt and key bas64 encoded:

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

Then create the Secret:

```bash
kubectl create -f controller-ssl.yaml
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
