// Copyright © 2016 Codaisseur BV <info@codaisseur.com>
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

package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func DeisConfigSet(envVar string) {
	_, err := exec.Command("deis", "config:set", envVar, "-a", cfg.app).Output()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("deis", "config").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done!\n\n%s\n", out)
}
