#!/bin/bash
# This script will test the coverage of all packages and push to Codecov

set -e

EXPECTED_COVERAGE="90.0"
COVERAGE_PROFILE="coverage.out"
COMMIT_ID=$(git describe --tags --always)

echo "Running go test and tool coverage for profile 'coverage.out'..."

make test

COVERAGE_PERCENTAGE=$(go tool cover -func=${COVERAGE_PROFILE}  | grep 'total:' | awk '{print $3}' | sed 's/%//')

if (( ${COVERAGE_PERCENTAGE%%.*} < ${EXPECTED_COVERAGE%%.*} || ( ${COVERAGE_PERCENTAGE%%.*} == ${EXPECTED_COVERAGE%%.*} && ${COVERAGE_PERCENTAGE##*.} < ${EXPECTED_COVERAGE##*.} ) )) ; then
	echo "coverage dropped to ${COVERAGE_PERCENTAGE}, expected was ${EXPECTED_COVERAGE}"
	exit 1
fi
echo "Coverage for ${COMMIT_ID}:  ${COVERAGE_PERCENTAGE}"
echo "Sending the report to CodeCov..."
if [[ -n "${CODECOV_REPO_TOKEN}" ]]; then
	bash <(curl -s https://codecov.io/bash) -t "${CODECOV_REPO_TOKEN}"
else
	echo "Error: CODECOV_REPO_TOKEN not set..."
fi
