/*
 * ZGrab Copyright 2015 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package ssh

// HandshakeLog contains detailed information about each step of the
// SSH handshake, and can be encoded to JSON.
type HandshakeLog struct {
	Banner             string            `json:"banner"`
	ServerID           *EndpointId       `json:"server_id"`
	ClientID           *EndpointId       `json:"client_id"`
	ServerKex          *KexInitMsg       `json:"server_key_exchange"`
	ClientKex          *KexInitMsg       `json:"client_key_exchange"`
	AlgorithmSelection *algorithms       `json:"algorithm_selection"`
	KeyExchange        kexAlgorithm      `json:"key_exchange"`
	Extensions         map[string][]byte `json:"extensions"`
	UserAuth           []string          `json:"userauth"`
	Crypto             *kexResult        `json:"crypto"`
}

type EndpointId struct {
	Raw             string `json:"raw"`
	ProtoVersion    string `json:"version"`
	SoftwareVersion string `json:"software"`
	Comment         string `json:"comment"`
}
