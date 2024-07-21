package tui

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestControl(t *testing.T) {
	data := "ChtbMzg7NTsyNTJtG1swbRtbMzg7NTsyNTJtG1swbSAgG1szODs1OzI1Mm1XYWl0aW5nIGZvciBtb2RlbBtbMG0bWzM4OzU7MjUybSByZXNwb25zZS4uLhtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbRtbMzg7NTsyNTJtIBtbMG0bWzM4OzU7MjUybSAbWzBtG1szODs1OzI1Mm0gG1swbQoK"
	s, err := base64.StdEncoding.DecodeString(data)
	require.NoError(t, err)

	str := string(s)
	x := stripControl.ReplaceAllString(str, "")
	assert.NotEqual(t, str, x)
	assert.True(t, strings.HasPrefix(str, x))

	suffix, ok := strings.CutPrefix(str, x)
	assert.True(t, ok)
	assert.True(t, len(suffix) > 0)
	assert.Len(t, suffix, 2438)
	assert.Equal(t, "\n\x1b[38;5;252m\x1b[0m\x1b[38;5;252m\x1b[0m  \x1b[38;5;252mWaiting for model\x1b[0m\x1b[38;5;252m response...", x)
}
