package reflection

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBuildQuery_Succeed(t *testing.T) {
	RegisterTestingT(t)

	// Write test cases
	type testCase struct {
		object        interface{}
		expectedQuery string
	}
	var testCases []testCase
	testCases = append(testCases, testCase{
		object: order{
			id:         13,
			customerId: 1313,
		},
		expectedQuery: "INSERT INTO order VALUES(13, 1313)",
	})
	testCases = append(testCases, testCase{
		object: customer{
			id:      13,
			name:    "LocTD",
			address: "Dai Tu",
		},
		expectedQuery: "INSERT INTO customer VALUES(13, LocTD, Dai Tu)",
	})

	// Run test cases
	for _, c := range testCases {
		query, err := buildQuery(c.object)
		Expect(err).Should(Succeed())
		Expect(query).Should(Equal(c.expectedQuery))
	}
}

func TestBuildQuery_Fail(t *testing.T) {
	RegisterTestingT(t)

	// Write test cases
	type testCase struct {
		object        interface{}
		expectedError string
	}
	var testCases []testCase
	testCases = append(testCases, testCase{
		object:        13,
		expectedError: "unsupported kind of object: int",
	})
	testCases = append(testCases, testCase{
		object: testCase{
			object:        13,
			expectedError: "",
		},
		expectedError: "unsupported kind of field: interface",
	})

	// Run test cases
	for _, c := range testCases {
		_, err := buildQuery(c.object)
		Expect(err).ShouldNot(Succeed())
		Expect(err.Error()).Should(Equal(c.expectedError))
	}
}
