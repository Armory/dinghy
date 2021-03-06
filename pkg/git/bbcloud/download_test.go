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

package bbcloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUrl(t *testing.T) {
	cases := []struct {
		endpoint string
		owner    string
		repo     string
		path     string
		branch   string
		url      string
	}{
		{
			endpoint: "https://api.github.com",
			owner:    "armory",
			repo:     "myrepo",
			path:     "my/path.yml",
			branch:   "mybranch",
			url:      "https://api.github.com/repositories/armory/myrepo/src/mybranch/my/path.yml?raw",
		},
	}

	for _, c := range cases {
		downloader := &FileService{
			Config: Config{
				Endpoint: c.endpoint,
			},
		}
		url := downloader.EncodeURL(c.owner, c.repo, c.path, c.branch)

		assert.Equal(t, c.url, url)
	}
}

func TestDecodeUrl(t *testing.T) {
	cases := []struct {
		endpoint string
		owner    string
		repo     string
		path     string
		branch   string
		url      string
	}{
		{
			endpoint: "https://api.github.com",
			owner:    "armory",
			repo:     "myrepo",
			path:     "my/path.yml",
			branch:   "mybranch",
			url:      "/repositories/armory/myrepo/src/mybranch/my/path.yml?raw",
		},
	}

	for _, c := range cases {
		downloader := &FileService{
			Config: Config{
				Endpoint: c.endpoint,
			},
		}
		org, repo, path, branch := downloader.DecodeURL(c.url)

		assert.Equal(t, c.owner, org)
		assert.Equal(t, c.repo, repo)
		assert.Equal(t, c.path, path)
		assert.Equal(t, c.branch, branch)
	}
}
