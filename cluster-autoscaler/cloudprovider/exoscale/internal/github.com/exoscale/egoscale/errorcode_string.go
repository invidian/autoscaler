/*
Copyright 2020 The Kubernetes Authors.

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

package egoscale

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unauthorized-401]
	_ = x[NotFound-404]
	_ = x[MethodNotAllowed-405]
	_ = x[UnsupportedActionError-422]
	_ = x[APILimitExceeded-429]
	_ = x[MalformedParameterError-430]
	_ = x[ParamError-431]
	_ = x[InternalError-530]
	_ = x[AccountError-531]
	_ = x[AccountResourceLimitError-532]
	_ = x[InsufficientCapacityError-533]
	_ = x[ResourceUnavailableError-534]
	_ = x[ResourceAllocationError-535]
	_ = x[ResourceInUseError-536]
	_ = x[NetworkRuleConflictError-537]
}

const (
	_ErrorCode_name_0 = "Unauthorized"
	_ErrorCode_name_1 = "NotFoundMethodNotAllowed"
	_ErrorCode_name_2 = "UnsupportedActionError"
	_ErrorCode_name_3 = "APILimitExceededMalformedParameterErrorParamError"
	_ErrorCode_name_4 = "InternalErrorAccountErrorAccountResourceLimitErrorInsufficientCapacityErrorResourceUnavailableErrorResourceAllocationErrorResourceInUseErrorNetworkRuleConflictError"
)

var (
	_ErrorCode_index_1 = [...]uint8{0, 8, 24}
	_ErrorCode_index_3 = [...]uint8{0, 16, 39, 49}
	_ErrorCode_index_4 = [...]uint8{0, 13, 25, 50, 75, 99, 122, 140, 164}
)

func (i ErrorCode) String() string {
	switch {
	case i == 401:
		return _ErrorCode_name_0
	case 404 <= i && i <= 405:
		i -= 404
		return _ErrorCode_name_1[_ErrorCode_index_1[i]:_ErrorCode_index_1[i+1]]
	case i == 422:
		return _ErrorCode_name_2
	case 429 <= i && i <= 431:
		i -= 429
		return _ErrorCode_name_3[_ErrorCode_index_3[i]:_ErrorCode_index_3[i+1]]
	case 530 <= i && i <= 537:
		i -= 530
		return _ErrorCode_name_4[_ErrorCode_index_4[i]:_ErrorCode_index_4[i+1]]
	default:
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
