// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	apb "google.golang.org/genproto/googleapis/api/annotations"
	lrpb "google.golang.org/genproto/googleapis/longrunning"
)

// GetFieldBehavior returns a slice of FieldBehavior annotations for
// the given field.
func GetFieldBehavior(f *desc.FieldDescriptor) []apb.FieldBehavior {
	opts := f.GetFieldOptions()
	if x, err := proto.GetExtension(opts, apb.E_FieldBehavior); err == nil {
		return x.([]apb.FieldBehavior)
	}
	return nil
}

// GetOperationInfo returns the LRO annotation.
func GetOperationInfo(m *desc.MethodDescriptor) *lrpb.OperationInfo {
	opts := m.GetMethodOptions()
	if x, err := proto.GetExtension(opts, lrpb.E_OperationInfo); err == nil {
		return x.(*lrpb.OperationInfo)
	}
	return nil
}
