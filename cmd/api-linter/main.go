// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Binary api-linter is a linter that checks Google APIs according to the API Improvement Proposals
// defined in https://aip.dev
package main

import (
	"log"
	"os"
)

func main() {
	if err := runCLI(rules(), configs(), os.Args); err != nil {
		log.Fatal(err)
	}
}
