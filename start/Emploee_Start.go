/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	// "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var EmpNo, EmpName  string  // Entities
//	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	// Initialize the chaincode
	EmpNo = args[0]
	EmpName = args[1]
	
	// err := stub.PutState(EmpNo, []byte(EmpName))
    if err != nil {
        fmt.Println("Could not save loan application to ledger", err)
        return nil, err
    }
	
	fmt.Println("Successfully saved Employee Emp" + EmpNo)
	fmt.Println("Successfully saved Employee EmpName" + EmpName)
	
	
	// fmt.Printf("Employee Number = %d, Employee Name = %d\n", args)

	

	return nil, nil
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

    fmt.Println("Entering Employee Details")
 
    if len(args) < 2 {
        fmt.Println("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for Employee ")
    }
 
    //var EmpNo   = args[0]
    //var EmpName = args[1]
 
    /* err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
    if err != nil {
        fmt.Println("Could not save loan application to ledger", err)
        return nil, err
    } */
 
    fmt.Println("Successfully saved loan application")
    return nil, nil
}



// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var EmpNo string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	EmpNo = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(EmpNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get employee Details for " + EmpNo + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	jsonResp := "{\"EmpNumber\":\"" + EmpNo + "\",\"EmpName\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}