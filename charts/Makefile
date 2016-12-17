lint:
	helm lint alea

build: lint
	helm package alea

install: build
	helm install --namespace services --name alea --values=settings.yaml alea
	if [ -f controller-ssl-${CLOUDEXT}.yaml ]; then kubectl create -f controller-ssl-${CLOUDEXT}.yaml; fi

upgrade: build
	helm upgrade alea --namespace services --values=custom.yaml alea

clean:
	helm delete alea --purge
	kubectl delete -n services ns services

ssl:
	kubectl create -f controller-ssl.yaml --namespace=services

run:
	kubectl run api --image=quay.io/codaisseur/alea-controller -i --tty -n services /bin/bash

release: build
	helm repo index . --url https://storage.googleapis.com/alea-charts
