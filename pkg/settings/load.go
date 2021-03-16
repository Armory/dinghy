/*
* Copyright 2019 Armory, Inc.

* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at

*    http://www.apache.org/licenses/LICENSE-2.0

* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Package settings is a single place to put all of the application settings.
package settings

import (
	"flag"
	"github.com/armory/dinghy/pkg/settings/source"
	"github.com/armory/dinghy/pkg/settings/source/local"
	"github.com/armory/dinghy/pkg/settings/source/remote"
)

type DinghyMode string

const (
	SingleTenant DinghyMode = "single-tenant"
	MultiTenant  DinghyMode = "multi-tenant"
)

// LoadSettings loads the Spring config from the default Spinnaker paths
// and merges default settings with the loaded settings
func LoadSettings() (source.Source, error) {
	mode := flag.String("mode", string(SingleTenant), "Dinghy mode")
	flag.Parse()
	var s source.Source = local.NewLocalSource()

	if *mode == string(MultiTenant) { // multi-tenant
		s = remote.NewRemoteSource()
	}

	return s, nil
}
