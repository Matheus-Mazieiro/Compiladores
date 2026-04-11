ANTLR_JAR=antlr-4.13.2-complete.jar

gen:
	java -jar ../../$(ANTLR_JAR) -Dlanguage=Go -o parser CalcLexer.g4 

test:
	java -jar ../../compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "go run main.go" gcc ../../temp ../../casos-de-teste/ "811508 821055 812050" t2