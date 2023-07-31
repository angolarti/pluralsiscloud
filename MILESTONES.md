# Docker

1. Adicionar os seguintes usecases
1.1. docker_list.go
1.2. docker_start.go
1.3. docker_rm.go
1.4. docker_stop.go
1.5. docker_remove.go

2. Criar as APIS para cada usecase mencionado no ponto 1

3. Criação de VM [libvirt](https://pkg.go.dev/libvirt.org/go/libvirt#hdr-Example_usage)
3.1. Criar uma VMs usando API do Libvrit
3.2. TODO

- [github.com/libvirt/libvirt-go](https://github.com/libvirt/libvirt-go)

4. Criar o fluxo de gestão de um Cluster Kubernetes
4.1. Cria um Deployments
4.2. Criar um Pods
4.3. TODO

5. Terraform Provider
5.1. Desenvolver um terraform provider "pcloud"
5.2. O Acesso o provider deve ser por intermédio de um token e url da plataforma