prep:
	dnf -y install nginx podman podman-compose

build:
	go build -o _build/pocketbase main.go

build-container:
	podman build -t lmsqueezy .

build-in-container:
	podman build -t lmsqueezy-builder -f Dockerfile.build . &&\
	podman run -it --rm -v ${PWD}:/opt/app-root/src/app/:z -v ${HOME}/.cache/go-build/:/opt/app-root/src/.cache/go-build/:Z -u 0 lmsqueezy-builder &&\
	podman-compose build