/*
Copyright 2023 The Kubernetes Authors.

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

package v1beta1

import (
	"context"

	"github.com/openshift/api/machine/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

type MachineSetWebhook struct{}

func (r *MachineSetWebhook) SetupWebhookWithManager(mgr manager.Manager) error {
	return builder.WebhookManagedBy(mgr).
		For(&v1beta1.MachineSet{}).
		WithValidator(r).
		Complete()
}

var _ webhook.CustomValidator = &MachineSetWebhook{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *MachineSetWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return nil
}

func (r *MachineSetWebhook) ValidateUpdate(ctx context.Context, oldRaw runtime.Object, newRaw runtime.Object) error {
	return nil
}

func (r *MachineSetWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
