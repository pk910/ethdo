// Copyright © 2019, 2020 Weald Technology Trading
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

package blockinfo

import (
	"context"
	"time"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethdo/util"
)

type dataIn struct {
	// System.
	timeout time.Duration
	quiet   bool
	verbose bool
	debug   bool
	// Operation.
	eth2Client eth2client.Service
	jsonOutput bool
	sszOutput  bool
	// Chain information.
	blockID string
	stream  bool
}

func input(ctx context.Context) (*dataIn, error) {
	data := &dataIn{}

	if viper.GetDuration("timeout") == 0 {
		return nil, errors.New("timeout is required")
	}
	data.timeout = viper.GetDuration("timeout")
	data.quiet = viper.GetBool("quiet")
	data.verbose = viper.GetBool("verbose")
	data.debug = viper.GetBool("debug")
	data.jsonOutput = viper.GetBool("json")
	data.sszOutput = viper.GetBool("ssz")

	data.stream = viper.GetBool("stream")

	var err error
	data.eth2Client, err = util.ConnectToBeaconNode(ctx, viper.GetString("connection"), viper.GetDuration("timeout"), viper.GetBool("allow-insecure-connections"))
	if err != nil {
		return nil, err
	}

	if viper.GetString("blockid") == "" {
		data.blockID = "head"
	} else {
		// Specific slot.
		data.blockID = viper.GetString("blockid")
	}

	return data, nil
}
