# Deis Backing Services Manager

At [Codaisseur](https://www.codaisseur.com) we wanted to provide a Heroku like environment for students. We <3 [Deis Workflow](https://deis.com/) as do our students, as you can see in this picture.

[![](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)](http://cd.sseu.re/Cksa06GXIAA7cDa.jpg)

## Anyway..

We use Rails a lot, and PostgreSQL, so we needed an easy way for students to set up their apps despite the fact that Deis does not provide a PostgreSQL service for
them like Heroku does. This manager app takes care of that.

## Set up the services cluster first

Check out the [instructions](/tree/master/kubernetes) to set up the services cluster on Kubernetes to get you started.

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
stolon-proxy-service   <none>                                    stolon-cluster=kube-stolon,stolon-proxy=true   10.247.50.217   5432/TCP
```

### Deploy the Backing Services Manager App

Use the IP to create the `DATABASE_URL`, we will use the one mentioned in the above output:

```
deis config:set DATABASE_URL=postgresql://stolon:password@10.247.50.217:5432/backing_services
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

## Roadmap

We will add more services as we go, like:

  - Redis
  - MongoDB
  - Memcached

Feel free to help us out :)
