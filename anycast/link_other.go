// +build !linux

/**
 * Tenta DNS Server
 *
 *    Copyright 2017 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * link_other.go: Placeholder for non-linux link management
 */

package anycast

import (
	"github.com/tenta-browser/tenta-dns/common"

	"github.com/sirupsen/logrus"
)

func addLink(dummies *[]string, b common.Netblock, cnt *uint, lg *logrus.Entry) {
	panic("Link adding not supported on this platform")
}

func removeLinks(dummies *[]string, lg *logrus.Entry) {
	lg.Errorf("Cannot remove links, not supported on this system type")
}
