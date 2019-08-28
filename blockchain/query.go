package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello(isHistory bool) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "query")
	if isHistory {
		args = append(args, "queryHistory")
	} else {
		args = append(args, "query")
	}
	args = append(args, "hello")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// ReadMarble
func (setup *FabricSetup) ReadMarble(name string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "readMarble")
	args = append(args, name)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}

// QueryMarblesByOwner
func (setup *FabricSetup) QueryMarblesByOwner(owner string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "queryMarblesByOwner")
	args = append(args, owner)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}
