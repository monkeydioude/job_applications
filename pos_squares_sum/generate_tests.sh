#!/bin/bash

# File names
TEST_FILE="test_max_cases.txt"
RESULT_FILE="test_results.txt"

# Initialize or clear files
echo "100" > "$TEST_FILE"
> "$RESULT_FILE" # Clears or creates the result file

# Generate test cases and results
for i in {1..100}; do
    echo "100" >> "$TEST_FILE"
    nums=() # Initialize array without 'local' since it's outside a function
    for j in {1..100}; do
        rand=$((RANDOM % 20000 - 10000))
        printf "%d " "$rand" >> "$TEST_FILE"
        nums+=("$rand")
    done
    echo "" >> "$TEST_FILE" # New line for next test case

    # Calculate and append the result for the current test case
    sum=0
    for num in "${nums[@]}"; do
        if [ "$num" -gt 0 ]; then
            ((sum+=num*num))
        fi
    done
    echo "$sum" >> "$RESULT_FILE"
done
