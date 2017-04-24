package main

import (

	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"strconv"
)

type UserBenefit struct{
	
}
type User struct {
	FfId    string
	Title	string
	Gender	string
	FirstName	string
	LastName	string
	Dob			string
	Email		string
	Country		string
	Address		string
	City 		string
	Zip			string
	CreatedBy	string
	TotalPoint	string
}

func main() {
	err := shim.Start(new(UserBenefit))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


func (t *UserBenefit) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	
	if len(args) <= 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting more then 1")
	}

u:=User{}
u.FfId=args[0]
u.Title=args[1]
u.Gender=args[2]
u.FirstName=args[3]
u.LastName=args[4]
u.Dob=args[5]
u.Email=args[6]	
u.Country=args[7]
u.Address=args[8]
u.City=args[9]
u.Zip=args[10]
u.CreatedBy=args[11]
u.TotalPoint=args[12]

json_byte, err:=json.Marshal(u);
	//hardcoded the key since not using the DB
	err = stub.PutState("user", json_byte)
	if err != nil {
		msg="UnSuccesful"
			return []byte(msg), err
	}
	msg="Success"
	return []byte(msg), nil
}

func (t *UserBenefit) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	// Handle different functions
	if function == "addDelete" {
		return t.addDeletePoints(stub,args)
	} 
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *UserBenefit) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	 if function == "read" {
        	return t.read(stub, args)
    } else if function == "getPoints" {
        return t.getPoints(stub, args)
    } else if function == "getUser"{
        return t.getUser(stub, args)
    }
	
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}


// read - query function to read key/value pair
func (t *UserBenefit) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	src_json:=[]byte(valAsbytes)
		u := User{}
	json.Unmarshal(src_json, &u)
	point:=[]byte(u.TotalPoint)
	

	return point, nil
}

func (t *UserBenefit) getPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key string
	
	var err error
fmt.Print("hi in get Points")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	src_json:=[]byte(valAsbytes)
		u := User{}
	json.Unmarshal(src_json, &u)
	jsonResp := []byte("{\"TotalPoints\":\"" + u.TotalPoint +"\"}")
	

	return jsonResp, nil
}
func (t *UserBenefit) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState("user")
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	return valAsbytes, nil
}

func (t *UserBenefit) addDeletePoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key,key2, jsonResp string
	var err error
	var msg string
	if len(args) <= 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	key2= args[1]
	fmt.Println(key2+" is the new key2")
	valAsbytes, err := stub.GetState("user")
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	src_json:=[]byte(valAsbytes)
u := User{}
	json.Unmarshal(src_json, &u)
	//ponits are hardcoded as of now can be made dynamic
	if key2== "add" {
			i, n := strconv.Atoi(u.TotalPoint)
			if n != nil {
				msg="Points not added"
		panic(n)
	}
			i=i+20
	u.TotalPoint=strconv.Itoa(i)
	}else if key2=="delete"	{
					i, n := strconv.Atoi(u.TotalPoint)
					if n != nil {
						msg="Point not substracted"
		panic(n)
	}
			i=i-20
	u.TotalPoint=strconv.Itoa(i)}
	json_byte, err:=json.Marshal(u);

	fmt.Println(u.TotalPoint+" the total point")
	err = stub.PutState("user", json_byte)
	return []byte(msg), nil
}



