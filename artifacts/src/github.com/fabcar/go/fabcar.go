package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// HumanCenteredData :  Define the HumanCenteredData structure, with properties.  Structure tags are used by encoding/json library
type HumanCentered struct {
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Timestamp  string `json:"timestamp"`
	Barometer  string `json:"barometer"`
	Heartrate  string `json:"heartrate"`
	Name       string `json:"name"`
}

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("fabcar_cc")

// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "queryHumanCentered":
		return s.queryHumanCentered(APIstub, args)
	case "initLedger":
		return s.initLedger(APIstub)
	case "createHumanCentered":
		return s.createHumanCentered(APIstub, args)
	case "queryAllHumanCentereds":
		return s.queryAllHumanCentereds(APIstub, args)
	case "getHistoryForAsset":
		return s.getHistoryForAsset(APIstub, args)
	case "restictedMethod":
		return s.restictedMethod(APIstub, args)
	case "test":
		return s.test(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}

	// return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryHumanCentered(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	humanAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(humanAsBytes)
}

func (s *SmartContract) test(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	HumanCenteredAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(HumanCenteredAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	humans := []HumanCentered{
		HumanCentered{Latitude: "37.6330061", Longitude: "126.9171758",Timestamp: "1637938938368", Barometer: "102655",Heartrate: "91", Name : "Hakpyeong Kim"},
		HumanCentered{Latitude: "37.6330061", Longitude: "126.9171758",Timestamp: "1637938932378", Barometer: "102655",Heartrate: "87", Name : "Hakpyeong Kim"},
		HumanCentered{Latitude: "37.5619743", Longitude: "126.9365936",Timestamp: "1637846346663", Barometer: "101783",Heartrate: "72", Name : "Choi heeju"},
	}

	i := 0
	for i < len(humans) {
		humanAsBytes, _ := json.Marshal(humans[i])
		APIstub.PutState("HUMAN"+strconv.Itoa(i), humanAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createHumanCentered(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var human = HumanCentered{Latitude: args[1], Longitude: args[2], Timestamp: args[3], Barometer: args[4], Heartrate: args[5], Name: args[6]}

	humanAsBytes, _ := json.Marshal(human)
	APIstub.PutState(args[0], humanAsBytes)


	return shim.Success(humanAsBytes)
}

func (s *SmartContract) queryAllHumanCentereds(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 : UserName / TimestampStart(Small) / TimestampEnd(Big)")
	}

	UserName       := args[0]
	TimestampStart := args[1]
	TimestampEnd   := args[2]

	startKey := UserName+TimestampStart
	endKey := UserName+TimestampEnd

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllHumanCentereds:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// func (s *SmartContract) queryAllHumanCentereds(APIstub shim.ChaincodeStubInterface) sc.Response {

// 	startKey := "HUMAN0"
// 	endKey := "HUMAN999"

// 	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(queryResponse.Key)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.WriteString(string(queryResponse.Value))
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf("- queryAllHumanCentereds:\n%s\n", buffer.String())

// 	return shim.Success(buffer.Bytes())
// }

func (s *SmartContract) restictedMethod(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// get an ID for the client which is guaranteed to be unique within the MSP
	//id, err := cid.GetID(APIstub) -

	// get the MSP ID of the client's identity
	//mspid, err := cid.GetMSPID(APIstub) -

	// get the value of the attribute
	//val, ok, err := cid.GetAttributeValue(APIstub, "attr1") -

	// get the X509 certificate of the client, or nil if the client's identity was not based on an X509 certificate
	//cert, err := cid.GetX509Certificate(APIstub) -

	val, ok, err := cid.GetAttributeValue(APIstub, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		shim.Error("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		shim.Error("Client identity doesnot posses the attribute")
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return shim.Error("Only user with role as APPROVER have access this method!")
	} else {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		humanAsBytes, _ := APIstub.GetState(args[0])
		return shim.Success(humanAsBytes)
	}

}

func (t *SmartContract) getHistoryForAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	humanName := args[0]

	resultsIterator, err := stub.GetHistoryForKey(humanName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForAsset returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
