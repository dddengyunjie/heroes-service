package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"strconv"
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

// GetMarblesByRange
func (setup *FabricSetup) GetMarblesByRange(startKey string, endKey string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "getMarblesByRange")
	args = append(args, startKey)
	args = append(args, endKey)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}

// GetMarblesByRangeWithPagination
func (setup *FabricSetup) GetMarblesByRangeWithPagination(startKey string, endKey string, pageSize int, bookMark string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "getMarblesByRangeWithPagination")
	args = append(args, startKey)
	args = append(args, endKey)
	args = append(args, strconv.Itoa(pageSize))
	args = append(args, bookMark)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}

// QueryMarbles
func (setup *FabricSetup) QueryMarbles(queryString string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "queryMarbles")
	args = append(args, queryString)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}

// GetHistoryForMarble
func (setup *FabricSetup) GetHistoryForMarble(name string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "getHistoryForMarble")
	args = append(args, name)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}

// QueryMarblesWithPagination
func (setup *FabricSetup) QueryMarblesWithPagination(queryString string, pageSize int, bookMark string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "queryMarblesWithPagination")
	args = append(args, queryString)
	args = append(args, strconv.Itoa(pageSize))
	args = append(args, bookMark)

	// Create a request (proposal) and send it
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4])}})
	if err != nil {
		return "", fmt.Errorf("failed to query funds: %v", err)
	}

	return string(response.Payload), nil
}
