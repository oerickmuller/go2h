# Go em (aproximadamente) 2 horas

## Ambiente de desenvolvimento

### Editor de código

Qualquer editor de código que suporte Go, como o [VSCode](https://code.visualstudio.com/) ou o Neovim. 

### Linux

Instale o [go](https://golang.org/dl/), de acordo com o guia de instalação, ou usando o gerenciador de pacotes [mise](https://mise.jdx.dev/getting-started.html)

Instale o **make**, de acordo com o seu gerenciador de pacotes.  

Configure o ambiente de execução do Go:

- Adicione o Go ao seu PATH
- Crie o diretório de trabalho

```bash
mkdir ~/go/{bin,src,pkg}
```

### Windows

Instale o [go](https://golang.org/dl/)

Depois do Go instalado, instale o Task (é como um Makefile, mas nativo em go)
```
go install github.com/go-task/task/v3/cmd/task@latest
```

Instale também o make
```
choco install make
```

## Extensões recomendadas

- VS Code

    - [ms-vscode.makefile-tools](https://marketplace.visualstudio.com/items?itemName=ms-vscode.makefile-tools)
    - [golang.go](https://marketplace.visualstudio.com/items?itemName=golang.go)

- Neovim

    - [ray-x go](https://github.com/ray-x/go.nvim)
