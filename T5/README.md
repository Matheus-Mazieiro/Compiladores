# Compilador para a Linguagem LA
Este projeto é parte do trabalho da disciplina de Compiladores na UFSCar, orientado pelo Prof. Jander. O objetivo é implementar, em etapas, um compilador para a linguagem LA (Linguagem Algorítmica), desenvolvida no âmbito do DC/UFSCar.

## T5 — Gerador de Código
Nesta etapa desenvolvemos um gerador de código para a linguagem LA (Linguagem Algorítmica) desenvolvida pelo prof. Jander, no âmbito do DC/UFSCar. O gerador de código deverá produzir código executável em C equivalente ao programa de entrada.


### Tecnologias Utilizadas
- [Go (1.22+)](https://go.dev)

- [Antlr](https://www.antlr.org/) — para geração do analisador a partir de uma gramática .g4

A gramática usada pode ser encontrada [neste arquivo](http://github.com/Matheus-Mazieiro/Compiladores/blob/master/T5/go-lexer/CalcLexer.g4), e o código gerado pelo ANTLR está disponível [nesta pasta](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T5/go-lexer).

### Como compilar e executar
Certifique-se de ter o Go instalado em sua máquina.
Antes de rodar o projeto, é necessário baixar e organizar as dependências:

```
go mod tidy
```

Entre no diretório do [analisador semântico](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T5/go-lexer) dentro da pasta do T5, e compile o programa:
```
go build -o <executavel> *.go
```
Caso `make` esteja instalado, pode executar `make build` para compilar o programa em `semantico`.

Para então executar
```
.\<executavel> <entrada> <saida>
```

### Gerar gramática
Para gerar novas gramáticas, basta escrever a gramática em CalcLexer.g4 e executar o comando gen do Makefile presente em T5/go-lexer
```
make gen
```

### Rodar os testes
O corretor automático desponibilizado pelo professor para a avaliação dos alunos também se encontra neste repositório. Para executá-lo, basta rodar os comandos build e test do Makefile presente em T5/go-lexer
```
make build
make test
```

Entretanto, para executar este comando em um computador que não tem make, basta executar diretamente

```
go build -o semantico *.go
java -jar ../../compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "go run main.go" gcc ../../temp ../../casos-de-teste/ "811508 821055 812050" t5

```

---

Trabalho realizado por:
- Carolina Martins Emilio Pelicce (RA: 811508)
- Maristella Ramalho Rangel (RA: 821055)
- Matheus de Almeida Mazieiro (RA: 812050)
