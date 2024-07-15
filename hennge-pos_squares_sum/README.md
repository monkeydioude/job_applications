Description
- We want you to calculate the sum of squares of given integers, excluding any negatives.
- The first line of the input will be an integer N (1 <= N <= 100), indicating the number of test cases to follow.
- Each of the test cases will consist of a line with an integer X (0 < X <= 100), followed by another line consisting of X number of space-separated integers Yn (-100 <= Yn <= 100).
- For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
- __Note: There should be no output until all the input has been received.__
- __Note 2: Do not put blank lines between test cases solutions.__
- __Note 3: Take input from standard input, and output to standard output.__

Rules
- Write your solution using Go Programming Language or Python Programming Language. Do not submit your solution with both languages at once!
- You may only use standard library packages.
- In addition, extra point is awarded if solution does not declare any global variables.

Specific rules for Go solution:
- Your source code must be a single file
- Do not use any for and goto statement
- Your solution will be tested against Go 1.20 (as of February 2023) or higher

Sample Input
```
2
4
3 -1 1 14
5
9 6 -53 32 16
```
Sample Output
```
206
1397
```