// Copyright 2022 Arduino SA
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package browser

import "os/exec"

// Find will find the browser
func Find(process string) ([]byte, error) {
	return []byte(process), nil
}

// Kill will kill a process
func Kill(process string) ([]byte, error) {
	cmd := exec.Command("Taskkill", "/F", "/IM", process+".exe")
	return cmd.Output()
}

// Start will start a command
func Start(command []byte, url string) ([]byte, error) {
	cmd := exec.Command("cmd", "/C", "start", string(command), url)
	return cmd.Output()
}
