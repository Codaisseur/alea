lint:
	helm lint alea

build: lint
	helm package alea

configure:
	./configure.sh

install: build
	helm install --namespace services --name alea --values=settings.yaml alea
	if [ -f controller-ssl.yaml ]; then kubectl create -f controller-ssl.yaml; fi

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
	git fetch
	git checkout gh-pages
	#git pull origin gh-pages
	mkdir -p ../releases
	mv alea-*.tgz ../releases/
	helm repo index releases --url https://alea.github.io/repo
