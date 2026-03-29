# Compilador para a Linguagem LA
Este projeto é parte do trabalho da disciplina de Compiladores na UFSCar, orientado pelo Prof. Jander. O objetivo é implementar, em etapas, um compilador para a linguagem LA (Linguagem Algorítmica), desenvolvida no âmbito do DC/UFSCar.

## T1 — Analisador Léxico
Nesta etapa, desenvolvemos o analisador léxico da linguagem LA. A função principal desse módulo é tokenizar o código-fonte, identificando os diferentes elementos léxicos da linguagem.

Além de identificar corretamente os tokens válidos, o analisador também deve capturar erros léxicos, como:

- Símbolos não reconhecidos;

- Comentários abertos e não fechados;

- Cadeias de caracteres não finalizadas.

### Tecnologias Utilizadas
- [Go (1.18+)](https://go.dev)

- [Antlr](https://www.antlr.org/) — para geração do analisador a partir de uma gramática .g4

A gramática usada pode ser encontrada [neste arquivo](https://github.com/Matheus-Mazieiro/Compiladores/blob/master/T1%20-%20Analisador%20Lexico/go-lexer/CalcLexer.g4), e o código gerado pelo ANTLR está disponível [nesta pasta](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T1%20-%20Analisador%20Lexico/go-lexer/lexer).

### Como compilar e executar
Certifique-se de ter o Go instalado em sua máquina.

Entre no diretório do [analisador lexico](https://github.com/Matheus-Mazieiro/Compiladores/tree/master/T1%20-%20Analisador%20Lexico/go-lexer) dentro da pasta do T1, e execute o programa:
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
