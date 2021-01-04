/*
 Copyright 2020 The Kubernetes Authors.

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

package config

import (
	nsxtcfg "k8s.io/cloud-provider-vsphere/pkg/nsxt/config"
)

// Config is used to read and store information from the cloud configuration file
type Config struct {
	Route RouteConfig
	NSXT  nsxtcfg.NsxtConfig
}

// RouteConfig contains the configuration for the route itself
type RouteConfig struct {
	RouterPath string
}
