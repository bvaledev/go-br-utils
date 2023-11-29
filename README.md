# Go BR Utils

## Visão Geral

`Go BR Utils` é um pacote go útil para validar CPF e CNPJ brasileiro. Este pacote fornece uma maneira simples e eficaz de verificar a validade de números de CPF ou CNPJ.

## Instalação

Para usar o `go-br-utils` no seu projeto Go, basta instalar o pacote via `go get`:

```bash
go get github.com/bvaledev/go-br-utils/pkg/validator
```

## Exemplo

```go
package main

import (
	"fmt"

	"github.com/bvaledev/go-br-utils/pkg/validator"
)

func main() {
	fmt.Println(validator.IsCPFCNPJ("714.602.380-01")) // true
	fmt.Println(validator.IsCPFCNPJ("35.040.714/0001-08")) // true
	fmt.Println(validator.IsCPFCNPJ("35.040.7")) //false

	fmt.Println(validator.IsCPF("714.602.380-01")) // true
	fmt.Println(validator.IsCNPJ("35.040.714/0001-08")) // true

	fmt.Println(validator.IsCPF("714.602.380-02")) //false
	fmt.Println(validator.IsCNPJ("35.040.714/0001-18")) //false
}

```