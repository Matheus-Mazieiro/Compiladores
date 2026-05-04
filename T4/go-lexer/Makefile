ANTLR_JAR=antlr-4.13.2-complete.jar

gen:
	java -jar ../../$(ANTLR_JAR) -Dlanguage=Go -visitor -o parser CalcLexer.g4 

test:
	java -jar ../../compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "./semantico" gcc ../../temp ../../casos-de-teste/ "811508 821055 812050" t4

build:
	go build -o semantico *.go
