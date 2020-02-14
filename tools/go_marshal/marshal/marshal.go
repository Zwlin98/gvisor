// Copyright 2019 The gVisor Authors.
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

// Package marshal defines the Marshallable interface for
// serialize/deserializing go data structures to/from memory, according to the
// Linux ABI.
//
// Implementations of this interface are typically automatically generated by
// tools/go_marshal. See the go_marshal README for details.
package marshal

import (
	"gvisor.dev/gvisor/pkg/usermem"
)

// Task provides a subset of kernel.Task, used in marshalling. We don't import
// the kernel package directly to avoid circular dependency.
type Task interface {
	// CopyScratchBuffer provides a task goroutine-local scratch buffer. See
	// kernel.CopyScratchBuffer.
	CopyScratchBuffer(size int) []byte

	// CopyOutBytes writes the contents of b to the task's memory. See
	// kernel.CopyOutBytes.
	CopyOutBytes(addr usermem.Addr, b []byte) (int, error)

	// CopyInBytes reads the contents of the task's memory to b. See
	// kernel.CopyInBytes.
	CopyInBytes(addr usermem.Addr, b []byte) (int, error)
}

// Marshallable represents a type that can be marshalled to and from memory.
type Marshallable interface {
	// SizeBytes is the size of the memory representation of a type in
	// marshalled form.
	SizeBytes() int

	// MarshalBytes serializes a copy of a type to dst. dst must be at least
	// SizeBytes() long.
	MarshalBytes(dst []byte)

	// UnmarshalBytes deserializes a type from src. src must be at least
	// SizeBytes() long.
	UnmarshalBytes(src []byte)

	// Packed returns true if the marshalled size of the type is the same as the
	// size it occupies in memory. This happens when the type has no fields
	// starting at unaligned addresses (should always be true by default for ABI
	// structs, verified by automatically generated tests when using
	// go_marshal), and has no fields marked `marshal:"unaligned"`.
	Packed() bool

	// MarshalUnsafe serializes a type by bulk copying its in-memory
	// representation to the dst buffer. This is only safe to do when the type
	// has no implicit padding, see Marshallable.Packed. When Packed would
	// return false, MarshalUnsafe should fall back to the safer but slower
	// MarshalBytes.
	MarshalUnsafe(dst []byte)

	// UnmarshalUnsafe deserializes a type by directly copying to the underlying
	// memory allocated for the object by the runtime.
	//
	// This allows much faster unmarshalling of types which have no implicit
	// padding, see Marshallable.Packed. When Packed would return false,
	// UnmarshalUnsafe should fall back to the safer but slower unmarshal
	// mechanism implemented in UnmarshalBytes.
	UnmarshalUnsafe(src []byte)

	// CopyIn deserializes a Marshallable type from a task's memory. This may
	// only be called from a task goroutine. This is more efficient than calling
	// UnmarshalUnsafe on Marshallable.Packed types, as the type being
	// marshalled does not escape. The implementation should avoid creating
	// extra copies in memory by directly deserializing to the object's
	// underlying memory.
	CopyIn(task Task, addr usermem.Addr) (int, error)

	// CopyOut serializes a Marshallable type to a task's memory. This may only
	// be called from a task goroutine. This is more efficient than calling
	// MarshalUnsafe on Marshallable.Packed types, as the type being serialized
	// does not escape. The implementation should avoid creating extra copies in
	// memory by directly serializing from the object's underlying memory.
	CopyOut(task Task, addr usermem.Addr) (int, error)
}
