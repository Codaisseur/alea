export VERSION=0.1.0

build:
	docker build -t quay.io/codaisseur/deis-backing-services-api \
		-t quay.io/codaisseur/deis-backing-services-api:v${VERSION} .
		docker build -t quay.io/codaisseur/deis-backing-services-api \
			-t quay.io/codaisseur/deis-backing-services-api:latest .

push:
	docker push quay.io/codaisseur/deis-backing-services-api:v${VERSION}
	docker push quay.io/codaisseur/deis-backing-services-api:latest
