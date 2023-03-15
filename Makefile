prep:
	dnf -y install nginx podman podman-compose

build:
	go build -o _build/pocketbase main.go

build-container:
	podman build -t lmsqueezy .