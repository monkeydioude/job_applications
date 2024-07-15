OUTPUT_FILE=./test_bin_output.txt

TEST_GENERATOR_SCRIPT=./generate_tests.sh
TEST_RESULTS_FILE=./test_results.txt
TEST_INPUT_FILE=./test_max_cases.txt

.PHONY: test
test:
	sh $(TEST_GENERATOR_SCRIPT) ; go run main.go < $(TEST_INPUT_FILE) > $(OUTPUT_FILE); diff -y --suppress-common-lines  $(OUTPUT_FILE) $(TEST_RESULTS_FILE)