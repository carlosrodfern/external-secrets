/*
Copyright © 2025 ESO Maintainer Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package logs provides functionality for contextual logging in the external-secrets system.
package logs

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NameAppends contains the names to be added to the logger.
type NameAppends []string

// CtxLog creates a new logger from context, or ctrl.Log if no present.
func CtxLog(ctx context.Context) logr.Logger {
	var log logr.Logger
	if found, err := logr.FromContext(ctx); err == nil {
		log = found
	} else {
		log = ctrl.Log
	}
	return log
}

// CtxLogWithNames creates a new logger from context, or ctrl.Log if no present, and appends the names via WithName.
func CtxLogWithNames(ctx context.Context, names NameAppends) logr.Logger {
	log := CtxLog(ctx)
	for _, name := range names {
		log = log.WithName(name)
	}
	return log
}
