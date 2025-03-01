// Copyright Istio Authors
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

package kubetypes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type InformerOptions struct {
	// A selector to restrict the list of returned objects by their labels.
	LabelSelector string
	// A selector to restrict the list of returned objects by their fields.
	FieldSelector string
}

// WriteAPI exposes a generic API for a client go type for write operations.
type WriteAPI[T runtime.Object] interface {
	Create(ctx context.Context, object T, opts metav1.CreateOptions) (T, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result T, err error)
	Update(ctx context.Context, object T, opts metav1.UpdateOptions) (T, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
}

// WriteAPI exposes a generic API for a client go type for status operations.
// Not all types have status, so they need to be split out
type WriteStatusAPI[T runtime.Object] interface {
	UpdateStatus(ctx context.Context, object T, opts metav1.UpdateOptions) (T, error)
}

// ReadAPI exposes a generic API for a client go type for read operations.
type ReadAPI[T runtime.Object, TL runtime.Object] interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (T, error)
	List(ctx context.Context, opts metav1.ListOptions) (TL, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

// ApplyAPI exposes a generic API for a client go type for apply operations.
type ApplyAPI[T runtime.Object, TA runtime.Object] interface {
	Apply(ctx context.Context, secret TA, opts metav1.ApplyOptions) (result T, err error)
}

// FullAPI exposes a generic API for a client go type for all operations.
// Note each type can also have per-type specific "Expansions" not covered here.
type FullAPI[T runtime.Object, TL runtime.Object, TA runtime.Object] interface {
	ReadAPI[T, TL]
	WriteAPI[T]
	ApplyAPI[T, TA]
}
