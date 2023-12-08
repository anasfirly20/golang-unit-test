# golang-unit-test

1. go test -v -run TestFunctionName (to run test on a specific function)
2. go test -v (to run all tests)
3. go test -run /SubTestName (to run all unit test that has SubTestName)

# Assertion (testify package)

https://github.com/stretchr/testify

# SubTest

1. go test -v -run TestFunctionName/SubTestName

# Run Benchmark

1. go test -v -bench=. (Unit test will also be run)
2. go test -v -run=NotMathUnitTest -bench=.
3. go test -v -run=NotMathUnitTest -bench=BenchmarkTest (run a specific benchmark)
4. go test -v -bench=. ./... (run from upper folder)
5. go test -v -run=NotMathUnitTest -bench=. ./... (run from upper folder only benchmark)
