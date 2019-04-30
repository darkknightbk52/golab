package json

import (
	"encoding/json"
	"fmt"
	. "github.com/onsi/gomega"
	"io"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	RegisterTestingT(t)

	type Customer struct {
		ID   int
		Name string
		Age  int
	}

	type testCase struct {
		jsonDetail       string
		expectedCustomer Customer
	}

	testCases := make(map[int]testCase)
	var jsonDetails []string
	{
		tc := testCase{
			jsonDetail: `{"ID": 1, "Name": "LocTD", "Age": 13}`,
			expectedCustomer: Customer{
				Name: "LocTD",
				Age:  13,
			},
		}
		testCases[1] = tc
		jsonDetails = append(jsonDetails, tc.jsonDetail)
	}

	{
		tc := testCase{
			jsonDetail: `{"ID": 2, "Name": "HaiTT", "Age": 13}`,
			expectedCustomer: Customer{
				Name: "HaiTT",
				Age:  13,
			},
		}
		testCases[2] = tc
		jsonDetails = append(jsonDetails, tc.jsonDetail)
	}

	jsonStream := strings.Join(jsonDetails, "\n")
	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var c Customer
		err := decoder.Decode(&c)
		if err == io.EOF {
			break
		}
		Expect(err).Should(Succeed())
		Expect(c.Name).Should(Equal(testCases[c.ID].expectedCustomer.Name))
		Expect(c.Age).Should(Equal(testCases[c.ID].expectedCustomer.Age))
	}
}

func TestUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	type Engine struct {
		ID   int
		Code string
	}

	type Car struct {
		ID      int
		Model   string
		Engines []Engine
	}

	type testCase struct {
		serializedData []byte
		expectedCar    Car
	}

	var testCases []testCase
	{
		tc := testCase{
			serializedData: []byte(`{"ID": 13, "Model": "Mercedes", "Engines": [{"ID": 13, "Code": "Front"}, {"ID": 14, "Code": "Behind"}]}`),
			expectedCar: Car{
				ID:    13,
				Model: "Mercedes",
			},
		}
		tc.expectedCar.Engines = append(tc.expectedCar.Engines, Engine{
			ID:   13,
			Code: "Front",
		})
		tc.expectedCar.Engines = append(tc.expectedCar.Engines, Engine{
			ID:   14,
			Code: "Behind",
		})
		testCases = append(testCases, tc)
	}

	{
		tc := testCase{
			serializedData: []byte(`{"ID": 113, "Model": "Bentley", "Engines": [{"ID": 113, "Code": "Above"}, {"ID": 114, "Code": "Under"}]}`),
			expectedCar: Car{
				ID:    113,
				Model: "Bentley",
			},
		}
		tc.expectedCar.Engines = append(tc.expectedCar.Engines, Engine{
			ID:   113,
			Code: "Above",
		})
		tc.expectedCar.Engines = append(tc.expectedCar.Engines, Engine{
			ID:   114,
			Code: "Under",
		})
		testCases = append(testCases, tc)
	}

	compare := func(source, expected *Car) bool {
		if source.ID != expected.ID {
			return false
		}
		if source.Model != expected.Model {
			return false
		}
		for _, s := range source.Engines {
			found := false
			for _, e := range expected.Engines {
				if s.ID == e.ID {
					found = true
					if s.Code != e.Code {
						return false
					}
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	for _, tc := range testCases {
		var car Car
		err := json.Unmarshal(tc.serializedData, &car)
		Expect(err).Should(Succeed())
		same := compare(&car, &tc.expectedCar)
		Expect(same).Should(BeTrue())
		fmt.Println(car)
	}
}
