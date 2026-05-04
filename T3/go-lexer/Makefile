ANTLR_JAR=antlr-4.13.2-complete.jar

gen:
	java -jar ../../$(ANTLR_JAR) -Dlanguage=Go -visitor -o parser CalcLexer.g4 

test:
	java -jar ../../compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "/home/matheus-ubunto/Compiladores/T3/go-lexer/semantico" gcc ../../temp ../../casos-de-teste/ "811508 821055 812050" t3

build:
	go build -o semantico *.go
