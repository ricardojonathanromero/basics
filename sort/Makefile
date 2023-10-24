# Define the command to run for each sub-folder
SUB_FOLDER_CMD = echo "Processing folder:" $@
RUN_TESTS = echo && echo "=======================" && echo && echo "Runing tests" && go test -v
RUN_BENCHMARK_TESTS = echo && echo "=======================" && echo && echo "Runing benchmark tests" && go test -bench=. -benchmem -v
RUN_FUZZ_TESTS = echo && echo "=======================" && echo && echo "Runing fuzz tests" && go test -fuzz=Fuzz -fuzztime=20s

# List all sub-folders in the project's root directory
SUB_FOLDERS := $(wildcard */)

# Makefile rule to iterate through sub-folders and run the command
all: $(SUB_FOLDERS)

$(SUB_FOLDERS):
	@cd $@ && $(SUB_FOLDER_CMD) && $(RUN_TESTS) && $(RUN_BENCHMARK_TESTS) && $(RUN_FUZZ_TESTS)

.PHONY: all $(SUB_FOLDERS)
