# Compilador para a Linguagem LA
Este projeto é parte do trabalho da disciplina de Compiladores na UFSCar, orientado pelo Prof. Jander. O objetivo é implementar, em etapas, um compilador para a linguagem LA (Linguagem Algorítmica), desenvolvida no âmbito do DC/UFSCar.

## T2 — Analisador Semântico
Nesta etapa, desenvolvemos o analisador semântico da linguagem LA. O analisador semântico verifica se o programa está correto em termos de significado, analisando o uso e a declaração de identificadores, controle de escopo e compatibilidade de tipos em atribuições e expressões. Ele detecta erros como variáveis não declaradas, redefinições no mesmo escopo e incompatibilidades de tipo, sem interromper a execução, permitindo reportar todos os erros encontrados no código.


### Tecnologias Utilizadas
- [Go (1.22+)](https://go.dev)

- [Antlr](https://www.antlr.org/) — para geração do analisador a partir de uma gramática .g4

A gramática usada pode ser encontrada [neste arquivo](http://github.com/Matheus-Mazieiro/Compiladores/blob/master/T2/go-lexer/CalcLexer.g4), e o código gerado pelo ANTLR está disponível [nesta pasta](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T2/go-lexer).

### Como compilar e executar
Certifique-se de ter o Go instalado em sua máquina.

Entre no diretório do [analisador semântico](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T3/go-lexer) dentro da pasta do T3, e execute o programa:
```
./go run main.go <arquivo_de_entrada> <arquivo_de_saida>
```

### Gerar gramática
Para gerar novas gramáticas, basta escrever a gramática em CalcLexer.g4 e executar o comando gen do Makefile presente em T1/go-lexer
```
make gen
```

### Rodar os testes
O corretor automático desponibilizado pelo professor para a avaliação dos alunos também se encontra neste repositório. Para executá-lo, basta rodar o comando test do Makefile presente em T1/go-lexer
```
make test
```

---

Trabalho realizado por:
- Carolina Martins Emilio Pelicce (RA: 811508)
- Maristella Ramalho Rangel (RA: 821055)
- Matheus de Almeida Mazieiro (RA: 812050)
