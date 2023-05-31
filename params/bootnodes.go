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
	// Ethereum Foundation Go Bootnodes//TODO
	"enode://eb565f4c270e45ee13aaf4e27e452ee11591c2dc4ab795a8987f705a50fd3604f751f7c0d8b669287b2ba1b3713effea13bb0c25d245066b1a703bd0c477f86a@35.73.253.58:32668",
	"enode://98eda17a83decd4a9b606b926f12b4c965b7b6b851a277e895861338eb7eebae8d4b48523eee33fafee5ad9a5f8d3fb14e56ad75ea70e80de815dc98962d9904@52.192.142.14:32668",
	"enode://d7ee8747aa8a0e4ab61ef27e18045a36014f1458a018394f587411d4959c9cf41742bd1402f693e9eb8a76b9fdc4e4d12b5a49c0000e7b08092e7f357b72a4c1@3.115.0.208:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://163fcba9e550fb31baf895153111fc212a2ef73ca904c4437cf164f69f0a2046b5b2090d2a7ab2dcec9dbc1e9787bb94468c39b7ea1c56dfb55664971924e47b@34.64.149.189:32668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
