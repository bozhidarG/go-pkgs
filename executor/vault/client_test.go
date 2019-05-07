// Copyright 2018 SumUp Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vault

import (
	stdRsa "crypto/rsa"
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"github.com/sumup-oss/go-pkgs/logger"
	"github.com/sumup-oss/go-pkgs/logger/testlogger"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run(
		"it creates a new client with specified logger, vault API client and private key",
		func(t *testing.T) {
			apiClientArg := &api.Client{}
			privKeyArg := &stdRsa.PrivateKey{}
			loggerArg := testlogger.NewTestLogger(logger.DebugLevel)

			actual := NewClient(loggerArg, apiClientArg, privKeyArg)
			assert.Equal(t, loggerArg, actual.logger)
			assert.Equal(t, apiClientArg, actual.vaultClient)
			assert.Equal(t, privKeyArg, actual.privateKey)
		},
	)
}
