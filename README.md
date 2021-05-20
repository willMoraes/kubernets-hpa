# Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

## Desafio Kubernetes e hpa

Para acessar a imagem no Docker Hub, [clique aqui](https://hub.docker.com/r/willmsmoraes/go-hpa) 

Para acessar o PR aprovado pelo processo de CI, [clique aqui](https://github.com/willMoraes/kubernets-hpa/pull/1/checks)

Para gerar um pod que faz requisições para o nosso service, rode o seguinte comando:

`
kubectl run -it load-generator --image=busybox /bin/sh -c "while sleep 0.01; do wget -q -O- http://go-hpa.default.svc.cluster.local; done"
`
