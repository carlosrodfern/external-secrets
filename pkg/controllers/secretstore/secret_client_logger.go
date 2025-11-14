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

package secretstore

import (
	"context"

	esv1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1"
	"github.com/external-secrets/external-secrets/runtime/logs"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
)

var (
	prefixNames = logs.NameAppends{"provider"}
)

type secretClientLogger struct {
	client esv1.SecretsClient
}

func (s *secretClientLogger) GetSecret(ctx context.Context, ref esv1.ExternalSecretDataRemoteRef) ([]byte, error) {
	log := s.getLog(ctx)
	log.WithValues("key", ref.Key).Info("getting secret")
	return s.client.GetSecret(logr.NewContext(ctx, log), ref)
}

func (s *secretClientLogger) PushSecret(ctx context.Context, secret *corev1.Secret, data esv1.PushSecretData) error {
	log := s.getLog(ctx)
	log.WithValues("secret", secret.GetName(), "key", data.GetSecretKey()).Info("pushing secret")
	return s.client.PushSecret(logr.NewContext(ctx, log), secret, data)
}

func (s *secretClientLogger) DeleteSecret(ctx context.Context, remoteRef esv1.PushSecretRemoteRef) error {
	log := s.getLog(ctx)
	log.WithValues("key", remoteRef.GetRemoteKey()).Info("deleting secret")
	return s.client.DeleteSecret(logr.NewContext(ctx, log), remoteRef)
}

func (s *secretClientLogger) SecretExists(ctx context.Context, remoteRef esv1.PushSecretRemoteRef) (bool, error) {
	log := s.getLog(ctx)
	log.WithValues("key", remoteRef.GetRemoteKey()).Info("checking secret existence")
	return s.client.SecretExists(logr.NewContext(ctx, log), remoteRef)
}

func (s *secretClientLogger) Validate() (esv1.ValidationResult, error) {
	return s.client.Validate()
}

func (s *secretClientLogger) GetSecretMap(ctx context.Context, ref esv1.ExternalSecretDataRemoteRef) (map[string][]byte, error) {
	log := s.getLog(ctx)
	log.WithValues("key", ref.Key).Info("getting secret map")
	return s.client.GetSecretMap(logr.NewContext(ctx, log), ref)
}

func (s *secretClientLogger) GetAllSecrets(ctx context.Context, ref esv1.ExternalSecretFind) (map[string][]byte, error) {
	log := s.getLog(ctx)
	if ref.Name != nil {
		log = log.WithValues("secretRegex", ref.Name.RegExp)
	} else if ref.Path != nil {
		log = log.WithValues("secretPath", ref.Path)
	}
	log.Info("getting all secrets")
	return s.client.GetAllSecrets(logr.NewContext(ctx, log), ref)
}

func (s *secretClientLogger) Close(ctx context.Context) error {
	log := s.getLog(ctx)
	return s.client.Close(logr.NewContext(ctx, log))
}

func (s *secretClientLogger) GetNameAppends() logs.NameAppends {
	return s.client.GetNameAppends()
}

func (s *secretClientLogger) getLog(ctx context.Context) logr.Logger {
	names := prefixNames
	names = append(names, s.client.GetNameAppends()...)
	return logs.CtxLogWithNames(ctx, names)
}
