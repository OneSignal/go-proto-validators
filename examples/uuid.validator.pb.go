// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/uuid.proto

package validator_examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/OneSignal/go-proto-validators"
	_ "github.com/OneSignal/schema.protobuf/gen/go/uuid/v1"
	regexp "regexp"
	github_com_OneSignal_go_proto_validators "github.com/OneSignal/go-proto-validators"
	github_com_gofrs_uuid "github.com/gofrs/uuid"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_UUIDMsg_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *UUIDMsg) Validate() error {
	if !_regex_UUIDMsg_UserId.MatchString(this.UserId) {
		return github_com_OneSignal_go_proto_validators.FieldError("user_id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_OneSignal_go_proto_validators.FieldError("user_id", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	for index, item := range this.Xyz {
		if nil == item {
			return github_com_OneSignal_go_proto_validators.FieldError("xyz", fmt.Errorf("nil UUID at index %d", index))
		} else {
			if _, err := github_com_gofrs_uuid.FromBytes(item.Data); err != nil {
				return github_com_OneSignal_go_proto_validators.FieldError("xyz", fmt.Errorf("%v at index %d", err, index))
			}
		}
		if item != nil {
			if err := github_com_OneSignal_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_OneSignal_go_proto_validators.FieldError("xyz", err)
			}
		}
		// index += 1 is some dummy code to trick the compiler in case we don't use
		// idex as part of an error message
		index += 1
	}
	for index, item := range this.Abc {
		if item == "" {
			return github_com_OneSignal_go_proto_validators.FieldError("abc", fmt.Errorf(`value '%v' must not be an empty string at index %d`, item, index))
		}
		// index += 1 is some dummy code to trick the compiler in case we don't use
		// idex as part of an error message
		index += 1
	}
	return nil
}
func (this *Foo) Validate() error {
	return nil
}
