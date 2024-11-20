/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/submariner-io/admiral/pkg/log"
	"github.com/submariner-io/admiral/pkg/resource"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	nodeutil "k8s.io/component-helpers/node/util"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var logger = log.Logger{Logger: logf.Log.WithName("Node")}

var Retry = wait.Backoff{
	Steps:    5,
	Duration: 5 * time.Second,
	Factor:   1.2,
	Jitter:   0.1,
}

func GetLocalNode(clientset kubernetes.Interface) (*v1.Node, error) {
	nodeName, ok := os.LookupEnv("NODE_NAME")
	if !ok {
		return nil, errors.New("error reading the NODE_NAME from the environment")
	}

	var node *v1.Node

	err := retry.OnError(Retry, func(err error) bool {
		logger.Warningf("Error reading the local node - retrying: %v", err)
		return true
	}, func() error {
		var err error

		node, err = clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
		if err != nil {
			return errors.Wrapf(err, "unable to find local node %q", nodeName)
		}

		return nil
	})

	return node, errors.Wrapf(err, "failed to get local node %q", nodeName)
}

func WaitForLocalNodeReady(ctx context.Context, client kubernetes.Interface) {
	// In most cases the node will already be ready; otherwise, wait forever or until the context is cancelled.
	err := wait.PollUntilContextCancel(ctx, time.Second, true, func(_ context.Context) (bool, error) {
		localNode, err := GetLocalNode(client) //nolint:contextcheck // TODO - should pass the context parameter

		if err != nil {
			logger.Error(err, "Error retrieving local node")
		} else {
			_, condition := nodeutil.GetNodeCondition(&localNode.Status, v1.NodeReady)
			if condition != nil && condition.Status == v1.ConditionTrue {
				logger.Info("Local node ready")
				return true, nil
			}

			logger.Infof("Local node not ready - waiting. Conditions: %s", resource.ToJSON(localNode.Status.Conditions))
		}

		return false, nil
	})
	if err != nil {
		logger.Error(err, "Error waiting for local node")
	}
}
