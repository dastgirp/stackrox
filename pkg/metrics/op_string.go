// Code generated by "stringer -type=Op"; DO NOT EDIT.

package metrics

import "strconv"

const _Op_name = "AddAddManyCountDedupeGetGetAllGetManyGetGroupedListPruneResetRenameRemoveRemoveManySearchUpdateUpsertUpsertAll"

var _Op_index = [...]uint8{0, 3, 10, 15, 21, 24, 30, 37, 47, 51, 56, 61, 67, 73, 83, 89, 95, 101, 110}

func (i Op) String() string {
	if i < 0 || i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
