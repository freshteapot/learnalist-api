# Setting up kubernetes with k3s

# Notes
- You need a server ;)
- commands below work on debian
- setup ~/.ssh/config, to make your life easier

# Step 1
- log onto the server

## Install some useful tools
```
sudo apt-get install nmap lsof rsync jq
```

## Check what ports are open
```
nmap XXX.XXX.XXX.XXX
```

## Install k3s
- dont install traefik, we are going to use nginx-ingress
```sh
curl -sLS https://get.k3s.io | INSTALL_K3S_EXEC='server --tls-san 127.0.0.1 --no-deploy servicelb --no-deploy traefik' sh -
```

## Add server ip as path to registry.devbox in /etc/hosts
- faking the ownership of the domain, allows us to run an insecure docker registry,

```sh
XXX.XXX.XXX.XXX registry.devbox
```

## Download the kubeconfig

```sh
ssh lal01.learnalist.net sudo cat /etc/rancher/k3s/k3s.yaml > /Users/tinkerbell/.k3s/lal01.learnalist.net.yaml
```

## Set KUBECONFIG
- setting this allows kubectl to just work.

```sh
export KUBECONFIG="/Users/tinkerbell/.k3s/lal01.learnalist.net.yaml"
```


## Helm installs
- I am opting for fetch and templates, instead of installing helm in the cluster

## Download the packages

```sh
helm fetch --untar stable/docker-registry
helm fetch --untar stable/nginx-ingress
```

## Setup and apply nginx-ingress

```sh
rm -rf output/nginx-ingress/
helm template nginx-ingress --name frontdoor -f custom/nginx-ingress.yaml --output-dir ./output
kubectl apply -f output/nginx-ingress/templates/
```

## Setup and docker-registry

```sh
rm -rf output/docker-registry/
helm template docker-registry  -f custom/docker-registry.yaml --output-dir ./output
kubectl apply -f output/docker-registry/templates/
```

