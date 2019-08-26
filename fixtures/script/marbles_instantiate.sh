peer chaincode instantiate -o orderer.hf.chainhero.io:7050 --tls $CORE_PEER_TLS_ENABLED -cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/hf.chainhero.io/orderers/
orderer.hf.chainhero.io/msp/cacerts/ca.hf.chainhero.io-cert.pem -C $CHANNEL_NAME -n marbles -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org0MSP.member','Org0MSP.member')"
