// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"context"
	"time"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/model"
	"github.com/derailed/k9s/internal/model1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	unlockedIC = "🖍"
	lockedIC   = "🔑"
)

// Namespaceable tracks namespaces.
type Namespaceable interface {
	// ClusterWide returns true if the model represents resource in all namespaces.
	ClusterWide() bool

	// GetNamespace returns the model namespace.
	GetNamespace() string

	// SetNamespace changes the model namespace.
	SetNamespace(string)

	// InNamespace check if current namespace matches models.
	InNamespace(string) bool
}

// Lister tracks resource getter.
type Lister interface {
	// Get returns a resource instance.
	Get(ctx context.Context, path string) (runtime.Object, error)
}

// Tabular represents a tabular model.
type Tabular interface {
	Namespaceable
	Lister

	// SetInstance sets parent resource path.
	SetInstance(string)

	// SetLabelSelector sets the label selector.
	SetLabelSelector(labels.Selector)

	// GetLabelSelector fetch the label filter.
	GetLabelSelector() labels.Selector

	// Empty returns true if model has no data.
	Empty() bool

	// RowCount returns the model data count.
	RowCount() int

	// Peek returns current model data.
	Peek() *model1.TableData

	// Watch watches a given resource for changes.
	Watch(context.Context) error

	// Refresh forces a new refresh.
	Refresh(context.Context) error

	// SetRefreshRate sets the model watch loop rate.
	SetRefreshRate(time.Duration)

	// AddListener registers a model listener.
	AddListener(model.TableListener)

	// RemoveListener unregister a model listener.
	RemoveListener(model.TableListener)

	// Delete a resource.
	Delete(context.Context, string, *metav1.DeletionPropagation, dao.Grace) error

	// SetViewSetting injects custom cols specification.
	SetViewSetting(context.Context, *config.ViewSetting)
}
