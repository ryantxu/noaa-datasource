package plugin

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNextTide(t *testing.T) {
	frame, err := GetNextHighLow(int32(9414750), "feet")
	require.Nil(t, err)

	rows, err := frame.RowLen()
	assert.Nil(t, err)
	assert.Equal(t, 1, rows)

	next := time.Until(frame.Fields[0].At(0).(time.Time))
	mins := int(next / time.Minute)

	high := frame.Fields[2].At(0).(string) == "High"

	fmt.Printf("NEXT: %d, high=%v", mins, high)

	assert.GreaterOrEqual(t, mins, 0)
}
