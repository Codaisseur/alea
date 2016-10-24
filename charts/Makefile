lint:
	helm lint deis-backing-services

build: lint
	helm package deis-backing-services

configure:
	./configure.sh

install: build
	helm install --namespace services --name services --values=custom.yaml deis-backing-services
	if [ -f api-ssl.yaml ]; then kubectl create -f api-ssl.yaml; fi

upgrade: build
	helm upgrade services --namespace services --values=custom.yaml deis-backing-services

clean:
	helm delete services --purge
	kubectl delete -n services ns services

ssl:
	kubectl create -f api-ssl.yaml --namespace=services

run:
	kubectl run api --image=quay.io/codaisseur/deis-backing-services-api -i --tty -n services /bin/bash
