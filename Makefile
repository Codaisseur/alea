export VERSION=0.1.0

build:
	docker build -t quay.io/codaisseur/alea-controller \
		-t quay.io/codaisseur/alea-controller:v${VERSION} .
	docker build -t quay.io/codaisseur/alea-controller \
		-t quay.io/codaisseur/alea-controller:latest .

push:
	docker push quay.io/codaisseur/alea-controller:v${VERSION}
	docker push quay.io/codaisseur/alea-controller:latest
