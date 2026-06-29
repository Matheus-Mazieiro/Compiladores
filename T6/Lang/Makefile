ANTLR_JAR = antlr-4.13.2-complete.jar

TEST_DIR = ../testcases
TEST_OUT = ../saida

ENTRADA = $(TEST_DIR)/0-playground.pl
SAIDA = $(TEST_OUT)/0.cpp

gen:
	java -jar ../../$(ANTLR_JAR) \
		-Dlanguage=Go \
		-visitor \
		-o parser \
		Grammar.g4

build:
	go build -o pl-compiler *.go

test:
	./pl-compiler $(ENTRADA) $(SAIDA)

test-all:
	@mkdir -p $(TEST_OUT)
	@for file in $(wildcard $(TEST_DIR)/*); do \
		name=$$(basename "$$file"); \
		out="$(TEST_OUT)/$${name%.*}.cpp"; \
		echo "==> $$name -> $$(basename "$$out")"; \
		./pl-compiler "$$file" "$$out"; \
	done