// Copyright 2020 MongoDB Inc
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

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_private_endpoints.go -package=mocks github.com/mongodb/mongocli/internal/store PrivateEndpointLister,PrivateEndpointDescriber,PrivateEndpointCreator,PrivateEndpointDeleter,InterfaceEndpointDescriber,InterfaceEndpointCreator,InterfaceEndpointDeleter

type PrivateEndpointLister interface {
	PrivateEndpoints(string, *atlas.ListOptions) ([]atlas.PrivateEndpointConnection, error)
}

type PrivateEndpointDescriber interface {
	PrivateEndpoint(string, string) (*atlas.PrivateEndpointConnection, error)
}

type PrivateEndpointCreator interface {
	CreatePrivateEndpoint(string, *atlas.PrivateEndpointConnection) (*atlas.PrivateEndpointConnection, error)
}

type PrivateEndpointDeleter interface {
	DeletePrivateEndpoint(string, string) error
}

type InterfaceEndpointDescriber interface {
	InterfaceEndpoint(string, string, string) (*atlas.InterfaceEndpointConnection, error)
}

type InterfaceEndpointCreator interface {
	CreateInterfaceEndpoint(string, string, string) (*atlas.InterfaceEndpointConnection, error)
}

type InterfaceEndpointDeleter interface {
	DeleteInterfaceEndpoint(string, string, string) error
}

// PrivateEndpoints encapsulates the logic to manage different cloud providers
func (s *Store) PrivateEndpoints(projectID string, opts *atlas.ListOptions) ([]atlas.PrivateEndpointConnection, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).PrivateEndpoints.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// PrivateEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) PrivateEndpoint(projectID, privateLinkID string) (*atlas.PrivateEndpointConnection, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).PrivateEndpoints.Get(context.Background(), projectID, privateLinkID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreatePrivateEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) CreatePrivateEndpoint(projectID string, r *atlas.PrivateEndpointConnection) (*atlas.PrivateEndpointConnection, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).PrivateEndpoints.Create(context.Background(), projectID, r)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) DeletePrivateEndpoint(projectID, privateLinkID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).PrivateEndpoints.Delete(context.Background(), projectID, privateLinkID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateInterfaceEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) CreateInterfaceEndpoint(projectID, privateLinkID, interfaceEndpointID string) (*atlas.InterfaceEndpointConnection, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).PrivateEndpoints.AddOneInterfaceEndpoint(context.Background(), projectID, privateLinkID, interfaceEndpointID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// InterfaceEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) InterfaceEndpoint(projectID, privateLinkID, interfaceEndpointID string) (*atlas.InterfaceEndpointConnection, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).PrivateEndpoints.GetOneInterfaceEndpoint(context.Background(), projectID, privateLinkID, interfaceEndpointID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteInterfaceEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) DeleteInterfaceEndpoint(projectID, privateLinkID, interfaceEndpointID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).PrivateEndpoints.DeleteOneInterfaceEndpoint(context.Background(), projectID, privateLinkID, interfaceEndpointID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
