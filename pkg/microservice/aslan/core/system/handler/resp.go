/*
Copyright 2021 The KodeRover Authors.

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

package handler

type Registry struct {
	ID        string `json:"id"`
	RegAddr   string `json:"reg_addr"`
	IsDefault bool   `json:"is_default"`
	Namespace string `json:"namespace"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	TLSCert   string `json:"tls_cert"`
	TLSKey    string `json:"tls_key"`
}
