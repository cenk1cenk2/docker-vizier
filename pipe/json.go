package pipe

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"time"
)

type JsonDuration struct {
	time.Duration
}

func (field *JsonDuration) UnmarshalJSON(b []byte) error {
	var unmarshalled interface{}

	err := json.Unmarshal(b, &unmarshalled)

	if err != nil {
		return err
	}

	switch value := unmarshalled.(type) {
	case float64:
		field.Duration = time.Duration(value)
	case string:
		field.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid duration: %#v", unmarshalled)
	}

	return nil
}

func (field *VizierPermission) UnmarshalJSON(b []byte) error {
	type A VizierPermission
	t := &struct {
		User  *string `json:"user,omitempty"`
		Group *string `json:"group,omitempty"`
		*A
	}{
		A: (*A)(field),
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	if t.User != nil {
		var u *user.User
		var err error

		if _, e := strconv.ParseUint(*t.User, 10, 32); e != nil {
			u, err = user.Lookup(*t.User)
		} else {
			u, err = user.LookupId(*t.User)
		}

		if err != nil {
			return err
		}

		uid, err := strconv.ParseUint(u.Uid, 10, 32)

		if err != nil {
			return err
		}

		parsed := uint32(uid)
		field.User = &parsed

		if t.Group == nil {
			gid, err := strconv.ParseUint(u.Gid, 10, 32)

			if err != nil {
				return err
			}

			parsed := uint32(gid)
			field.Group = &parsed
		}
	}

	if t.Group != nil {
		var g *user.Group
		var err error

		if _, e := strconv.ParseUint(*t.Group, 10, 32); e != nil {
			g, err = user.LookupGroup(*t.Group)
		} else {
			g, err = user.LookupGroupId(*t.Group)
		}

		if err != nil {
			return err
		}

		gid, err := strconv.ParseUint(g.Gid, 10, 32)

		if err != nil {
			return err
		}

		parsed := uint32(gid)
		field.Group = &parsed
	}

	return nil
}

func (field *VizierChmod) UnmarshalJSON(b []byte) error {
	type A VizierChmod
	t := &struct {
		File *string `json:"file,omitempty"`
		Dir  *string `json:"dir,omitempty"`
		*A
	}{
		A: (*A)(field),
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	if t.File != nil {
		perm, err := strconv.ParseUint(*t.File, 8, 32)

		if err != nil {
			return err
		}

		parsed := os.FileMode(perm)
		field.File = &parsed
	}

	if t.Dir != nil {
		perm, err := strconv.ParseUint(*t.Dir, 8, 32)

		if err != nil {
			return err
		}

		parsed := os.FileMode(perm)
		field.Dir = &parsed
	}

	return nil
}
