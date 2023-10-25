# Define the command to run for each sub-folder
NEW_LINE = echo "======================="
SUB_FOLDER_CMD = echo "Processing folder:" $@
RUN_TESTS = echo && $(NEW_LINE) && echo && echo "Runing tests" && go test -v -parallel=5
RUN_COV = echo && $(NEW_LINE) && echo && echo "Runing cov report" && go test -parallel=5 -covermode=count -coverprofile coverage && go tool cover -func=coverage && go tool cover -html=coverage && rm coverage
RUN_BENCHMARK_TESTS = echo && $(NEW_LINE) && echo && echo "Runing benchmark tests" && go test -bench=. -benchmem -v -count=10 -run=Bench > 10_runs_bench.txt && go test -parallel=20 -bench=. -benchmem -v -count=20 -run=Bench > 20_runs_bench.txt && benchstat 10_runs_bench.txt 20_runs_bench.txt && rm 10_runs_bench.txt && rm 20_runs_bench.txt
RUN_FUZZ_TESTS = echo && $(NEW_LINE) && echo && echo "Runing fuzz tests" && go test -fuzz Fuzz -fuzztime 20s -parallel=5

# List all sub-folders in the project's root directory
SUB_FOLDERS := $(wildcard */)

# Makefile rule to iterate through sub-folders and run the command
all: $(SUB_FOLDERS)

$(SUB_FOLDERS):
	@cd $@ && $(SUB_FOLDER_CMD) && $(RUN_TESTS) && $(RUN_COV) && $(RUN_BENCHMARK_TESTS) && $(RUN_FUZZ_TESTS)

.PHONY: all $(SUB_FOLDERS)
