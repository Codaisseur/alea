# Deis Backing Services Manager

At [Codaisseur](https://www.codaisseur.com) we wanted to provide a Heroku like environment for students. We <3 [Deis Workflow](https://deis.com/) as do our students, as you can see in this picture.

[![](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)

## Anyway..

We use Rails a lot, and PostgreSQL, so we needed an easy way for students to set up their apps despite the fact that Deis does not provide a PostgreSQL service for
them like Heroku does. This manager app takes care of that.

## Install a Redis cluster with Helm

We'll use [Helm](http://helm.sh/) to set up a HA Redis cluster for us.
Refer to the Helm docs for install instructions.

```
helmc install redis-cluster
```

## Set up the Postgres service cluster

Check out the [instructions](/kubernetes) to set up the services cluster on Kubernetes to get you started.

## Setting up the Manager App

### Create a Deis app

Clone this repository and set up a deis app:

```
deis create
```

### Get the proxy service ip

```
kubectl get svc
NAME                   LABELS                                    SELECTOR                                       IP(S)           PORT(S)
redis-sentinel         10.115.253.131   <none>
26379/TCP
stolon-proxy-service   10.115.246.38    <none>        5432/TCP
```

### Deploy the Backing Services Manager App

Use the IP to create the `DATABASE_URL`, we will use the one mentioned in the above output:

```
deis config:set DATABASE_URL=postgresql://stolon:password@10.115.246.38:5432/backing_services
```

And same for `REDIS_URL`:

```
deis config:set
REDIS_URL=redis://10.115.253.131:26379/redis_services
```

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

## Get a Database for new Apps

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

## Roadmap

We will add more services as we go, like:

  - √ ~~PostgreSQL~~
  - √ ~~Redis~~
  - MongoDB
  - Memcached

Feel free to help us out or leave any feedback in the issues :)
