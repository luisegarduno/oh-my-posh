package segments

import (
	"strings"
	"testing"
	"time"

	"github.com/jandedobbeleer/oh-my-posh/src/properties"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime/mock"

	"github.com/stretchr/testify/assert"
)

func TestTimeSegmentTemplate(t *testing.T) {
	// set date for unit test
	currentDate := time.Now()
	cases := []struct {
		Case            string
		ExpectedString  string
		Template        string
		ExpectedEnabled bool
	}{
		{
			Case:            "no template",
			Template:        "",
			ExpectedString:  currentDate.Format("15:04:05"),
			ExpectedEnabled: true,
		},
		{
			Case:            "time only",
			Template:        "{{.CurrentDate | date \"15:04:05\"}}",
			ExpectedString:  currentDate.Format("15:04:05"),
			ExpectedEnabled: true,
		},
		{
			Case:            "lowercase",
			Template:        "{{.CurrentDate | date \"January 02, 2006 15:04:05\" | lower }}",
			ExpectedString:  strings.ToLower(currentDate.Format("January 02, 2006 15:04:05")),
			ExpectedEnabled: true,
		},
	}

	for _, tc := range cases {
		env := new(mock.Environment)

		tempus := &Time{
			CurrentDate: currentDate,
		}
		tempus.Init(properties.Map{}, env)

		assert.Equal(t, tc.ExpectedEnabled, tempus.Enabled())
		if tc.Template == "" {
			tc.Template = tempus.Template()
		}
		if tc.ExpectedEnabled {
			assert.Equal(t, tc.ExpectedString, renderTemplate(env, tc.Template, tempus), tc.Case)
		}
	}
}
