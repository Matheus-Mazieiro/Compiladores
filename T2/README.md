# Compilador para a Linguagem LA
Este projeto é parte do trabalho da disciplina de Compiladores na UFSCar, orientado pelo Prof. Jander. O objetivo é implementar, em etapas, um compilador para a linguagem LA (Linguagem Algorítmica), desenvolvida no âmbito do DC/UFSCar.

## T2 — Analisador Sintático
Nesta etapa, desenvolvemos o analisador sintático da linguagem LA. A função principal desse módulo é apontar onde existe erro sintático, indicando a linha e o lexema que causou a detecção do erro.

### Tecnologias Utilizadas
- [Go (1.22+)](https://go.dev)

- [Antlr](https://www.antlr.org/) — para geração do analisador a partir de uma gramática .g4

A gramática usada pode ser encontrada [neste arquivo](http://github.com/Matheus-Mazieiro/Compiladores/blob/master/T2/go-lexer/CalcLexer.g4), e o código gerado pelo ANTLR está disponível [nesta pasta](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T2/go-lexer/parser).

### Como compilar e executar
Certifique-se de ter o Go instalado em sua máquina.

Entre no diretório do [analisador sintático](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T2/go-lexer) dentro da pasta do T2, e execute o programa:
```
./go run main.go <arquivo_de_entrada> <arquivo_de_saida>
```

### Gerar gramática
Para gerar novas gramáticas, basta escrever a gramática em CalcLexer.g4 e executar o comando gen do Makefile presente em T2/go-lexer
```
make gen
```
Entretanto, para executar este comando em um computador que não tem make, basta executar diretamente
```
java -jar ../../antlr-4.13.2-complete.jar -Dlanguage=Go -o parser CalcLexer.g4 
```

### Rodar os testes
O corretor automático desponibilizado pelo professor para a avaliação dos alunos também se encontra neste repositório. Para executá-lo, basta rodar o comando test do Makefile presente em T1/go-lexer
```
make test
```
Entretanto, para executar este comando em um computador que não tem make, basta executar diretamente
```
java -jar ../../compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "go run main.go" gcc ../../temp ../../casos-de-teste/ "811508 821055 812050" t1
```

---

Trabalho realizado por:
- Carolina Martins Emilio Pelicce (RA: 811508)
- Maristella Ramalho Rangel (RA: 821055)
- Matheus de Almeida Mazieiro (RA: 812050)
