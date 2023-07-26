## Setup

### Siga os passos para poder fazer a instalação e usar o software


```shell
git clone -b develop https://github.com/angolarti/pluralsiscloud.git
```

```shell
cd pluralsiscloud && go mod tidy
```

Para activar o Live Reload

```shell
go install github.com/cosmtrek/air@latest 
```

Activar convetional commit com commitlint

```shell
npm install -g @commitlint/cli @commitlint/config-conventional
```

```shell
npm install
```

Activar o track de binários

```shell
apt install -y git-lfs
git lfs install
```

Rodar o aplicaçáo

```shell
npm run dev:live
```

Existe uma pasta http com com arquivo RestClisnt para executar as operações Rest. Para que o mesmo seja executado no VSCode deves instalar a extenssão **Rest Client**