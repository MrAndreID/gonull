package gonull

import (
    "time"
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

type NullInt32 struct {
	sql.NullInt32
}

type NullInt64 struct {
	sql.NullInt64
}

type NullBool struct {
	sql.NullBool
}

type NullTime struct {
	sql.NullTime
}

func (response NullString) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.String)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullString) UnmarshalJSON(data []byte) error {
    var str *string

    if err := json.Unmarshal(data, &str); err != nil {
        return err
    }

    if str != nil {
        if *str != "" {
            response.Valid = true
            response.String = *str
        } else {
            response.Valid = false
        }
    } else {
        response.Valid = false
    }

    return nil
}

func (response NullInt32) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.Int32)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullInt32) UnmarshalJSON(data []byte) error {
    var i *int32

    if err := json.Unmarshal(data, &i); err != nil {
        return err
    }

    if i != nil {
        response.Valid = true
        response.Int32 = *i
    } else {
        response.Valid = false
    }

    return nil
}

func (response NullInt64) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullInt64) UnmarshalJSON(data []byte) error {
    var i *int64

    if err := json.Unmarshal(data, &i); err != nil {
        return err
    }

    if i != nil {
        response.Valid = true
        response.Int64 = *i
    } else {
        response.Valid = false
    }

    return nil
}

func (response NullBool) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullBool) UnmarshalJSON(data []byte) error {
    var b *bool

    if err := json.Unmarshal(data, &b); err != nil {
        return err
    }

    if b != nil {
        response.Valid = true
        response.Bool = *b
    } else {
        response.Valid = false
    }

    return nil
}

func (response NullTime) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullTime) UnmarshalJSON(data []byte) error {
    var t *time.Time

    if err := json.Unmarshal(data, &t); err != nil {
        return err
    }

    if t != nil {
        response.Valid = true
        response.Time = *t
    } else {
        response.Valid = false
    }

    return nil
}