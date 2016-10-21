lint:
	helm lint deis-backing-services

build:
	helm package deis-backing-services

install:
	helm install --namespace apps --name services --values=dev.yaml deis-backing-services

clean:
	helm delete services --purge
