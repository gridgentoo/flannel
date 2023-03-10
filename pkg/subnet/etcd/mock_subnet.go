// Copyright 2015 flannel authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcd

import (
	"github.com/flannel-io/flannel/pkg/ip"
	"github.com/flannel-io/flannel/pkg/subnet"
)

func NewMockManager(registry *MockSubnetRegistry) subnet.Manager {
	return newLocalManager(registry, ip.IP4Net{}, ip.IP6Net{}, 60)
}

func NewMockManagerWithSubnet(registry *MockSubnetRegistry, sn ip.IP4Net, sn6 ip.IP6Net) subnet.Manager {
	return newLocalManager(registry, sn, sn6, 60)
}
