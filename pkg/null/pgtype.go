package null

import (
	"time"

	// [参考](github.com/volatiletech/null/v9)
	"github.com/jackc/pgx/v5/pgtype"
)

// NewInt2 creates a new Int2
func NewInt2(i int16, valid bool) pgtype.Int2 {
	return pgtype.Int2{
		Int16: i,
		Valid: valid,
	}
}

// NewInt4 creates a new Int4
func NewInt4(i int32, valid bool) pgtype.Int4 {
	return pgtype.Int4{
		Int32: i,
		Valid: valid,
	}
}

// NewInt8 creates a new Int8
func NewInt8(i int64, valid bool) pgtype.Int8 {
	return pgtype.Int8{
		Int64: i,
		Valid: valid,
	}
}

// NewFloat8 creates a new Float8
func NewFloat8(f float64, valid bool) pgtype.Float8 {
	return pgtype.Float8{
		Float64: f,
		Valid:   valid,
	}
}

// NewText creates a new Text
func NewText(s string, valid bool) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  valid,
	}
}

// NewDate creates a new Date - 2006-01-02
func NewDate(t time.Time, valid bool) pgtype.Date {
	return pgtype.Date{
		Time:             t,
		InfinityModifier: pgtype.Finite,
		Valid:            valid,
	}
}

// NewTimestamptz creates a new Timestamptz - 2006-01-02
func NewTimestamptz(t time.Time, valid bool) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:             t,
		InfinityModifier: pgtype.Finite,
		Valid:            valid,
	}
}

// NewBool creates a new Bool
func NewBool(b bool, valid bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  b,
		Valid: valid,
	}
}

// NewUUID creates a new UUID
func NewUUID(b [16]byte, valid bool) pgtype.UUID {
	return pgtype.UUID{
		Bytes: b,
		Valid: valid,
	}
}

// NewPoint creates a new Point
func NewPoint(x, y float64, valid bool) pgtype.Point {
	return pgtype.Point{
		P:     pgtype.Vec2{X: x, Y: y},
		Valid: valid,
	}
}
