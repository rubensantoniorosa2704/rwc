# Rwd Code Challenge

Esta aplicação é um desafio de código que lê arquivos de texto e realiza operações como contagem de linhas, bytes e palavras.

## Como Executar

Para executar a aplicação, use o comando abaixo, substituindo `[OPTIONS]` pela opção desejada e `test.txt` pelo caminho do arquivo que você deseja processar:

```sh
go run cmd/main.go [OPTIONS] test.txt
```

## Opções Disponíveis:
- `-c`: Conta o número de bytes no arquivo.
- `-l`: Conta o número de linhas no arquivo.
- `-w`: Conta o número de palavras no arquivo.