ANTLR_JAR=antlr-4.13.2-complete.jar

ENTRADA="../testcases/1-restricoes_isoladas.pl"
SAIDA="../saida/1.cpp"

gen:
	java -jar ../../$(ANTLR_JAR) \
		-Dlanguage=Go \
		-visitor \
		-o \
		parser \
		Grammar.g4 

test:
	./pl-compiler $(ENTRADA) $(SAIDA)

build:
	go build -o pl-compiler *.go
