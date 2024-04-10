// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/stakater/provider-databricks/internal/controller/cluster/cluster"
	policy "github.com/stakater/provider-databricks/internal/controller/cluster_policy/policy"
	pool "github.com/stakater/provider-databricks/internal/controller/instance_pool/pool"
	job "github.com/stakater/provider-databricks/internal/controller/job/job"
	notebook "github.com/stakater/provider-databricks/internal/controller/notebook/notebook"
	providerconfig "github.com/stakater/provider-databricks/internal/controller/providerconfig"
	secret "github.com/stakater/provider-databricks/internal/controller/secret/secret"
	scope "github.com/stakater/provider-databricks/internal/controller/secret_scope/scope"
	token "github.com/stakater/provider-databricks/internal/controller/token/token"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		policy.Setup,
		pool.Setup,
		job.Setup,
		notebook.Setup,
		providerconfig.Setup,
		secret.Setup,
		scope.Setup,
		token.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
