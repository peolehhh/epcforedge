// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
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

package af

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func getAllSubscriptions(afCtx *afContext,
	cliCtx context.Context) ([]TrafficInfluSub, *http.Response, error) {

	cliCfg := NewConfiguration()
	cli := NewClient(cliCfg)

	tsResp, resp, err := cli.TrafficInfluSubGetAllAPI.SubscriptionsGetAll(
		cliCtx, afCtx.cfg.AfId)

	if err != nil {

		log.Errf("AF Traffic Influance Subscription GET: %s", err.Error())

		return nil, nil, err
	}
	return tsResp, resp, nil

}

//GetAllSubscriptions function
func GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		tsResp     []TrafficInfluSub
		resp       *http.Response
		transID    int
		tsRespJSON []byte
	)

	afCtx := r.Context().Value(string("af-ctx")).(*afContext)
	cliCtx, cancel := context.WithCancel(context.Background())

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-osSignals
		log.Infof("Received signal: %#v", sig)
		cancel()
	}()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	transID, err = genTransactionID(afCtx)
	if err != nil {

		log.Errf("Traffic Influance Subscription PUT %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	afCtx.transactions[transID] = TrafficInfluSub{}
	tsResp, resp, err = getAllSubscriptions(afCtx, cliCtx)
	delete(afCtx.transactions, transID)
	if err != nil {
		log.Errf("Traffic Influence Subscription create : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tsRespJSON, err = json.Marshal(tsResp)
	if err != nil {
		log.Errf("Traffic Influence Subscription GET : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(tsRespJSON)
}
