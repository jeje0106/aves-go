// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://57ff4a3b6a40e5d930b524185587088e0d413977797d9a994f46a7ad2945a577425c91e948af4ad488c2a1e52e34a6ec7d746cc29a0e9b61c99dabeb69a0add5@195.231.86.211:30303",
	"enode://d6ddd57c8b87805189f45e1648a0b070beb3b7639ad68c6afb0f1b8297df302b5dd4646416e33991051f53814906ab347c3a6dbf808b386dc765d4a944626b31@83.78.96.42:30403",
        "enode://028ab37370b9794188f71804ea71f9113b5865c51c764ade7eebc62bf4ecf6835aff76b92f2ee14497b15dbfca2af8bcbde846c2612ef4bea86f6448852c03d6@75.119.138.139:30305",
        "enode://58ecb257378eb8d01e3e2016df0b606b52647a542faa5dd40c9637558ff741527a0e5904c6fce97480d70c57fa81d4640a75e652d48d6e56a1746d925a2cc535@89.232.109.242:31302",
        "enode://cc853691e880f76937ddbbd09f7055e8aa5ca2313d3fa54b4a3d6e722c324e63d8c4fd75485ac96d4cc5bdc219b450969c31c226a50c3076f46bdd05e3fdd2f6@78.47.129.247:30303",
        "enode://1d64d936f68b470e9085880072a1978fd9e15fddc80b23bfcd6c6df756aee50e8fceccfb323f07d919fa33b352575717d84456c9d0bb9ee8178088afbf954318@51.161.116.66:40611",
        "enode://57ff4a3b6a40e5d930b524185587088e0d413977797d9a994f46a7ad2945a577425c91e948af4ad488c2a1e52e34a6ec7d746cc29a0e9b61c99dabeb69a0add5@195.231.86.211:30303",
}


var V5Bootnodes = []string{
	// We dont use
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "aves"
	default:
		return ""
	}
	return net
}
