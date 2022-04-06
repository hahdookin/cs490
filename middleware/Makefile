SHELL=/bin/bash
OUTPUT=server.o

all: compile

compile:
	@go build -o $(OUTPUT) server.go

run: compile
	@./$(OUTPUT)

clean:
ifneq (,$(wildcard ./$(OUTPUT)))
	@rm ./$(OUTPUT)
endif