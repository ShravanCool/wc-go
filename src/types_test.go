package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_result_generateOutput(t *testing.T) {
	type fields struct {
		lineCount int
		wordCount int
		charCount int
		filename  string
		err       error
	}

	tests := []struct {
		name      string
		fields    fields
		testFlags FlagOptions
		want      string
		wantErr   bool
	}{
		{
			name: "success, all flags are set",
			fields: fields{
				lineCount: 3,
				wordCount: 6,
				charCount: 12,
				filename:  "test.txt",
				err:       nil,
			},
			testFlags: FlagOptions{
				LineFlag: true,
				WordFlag: true,
				CharFlag: true,
			},
			want:    "       3       6      12 test.txt\n",
			wantErr: false,
		},
		{
			name: "success, only lines flag is set",
			fields: fields{
				lineCount: 3,
				wordCount: 6,
				charCount: 12,
				filename:  "test.txt",
				err:       nil,
			},
			testFlags: FlagOptions{
				LineFlag: true,
			},
			want:    "       3 test.txt\n",
			wantErr: false,
		},
		{
			name: "success, only lines and words flag are set",
			fields: fields{
				lineCount: 3,
				wordCount: 6,
				charCount: 12,
				filename:  "test.txt",
				err:       nil,
			},
			testFlags: FlagOptions{
				LineFlag: true,
				WordFlag: true,
			},
			want:    "       3       6 test.txt\n",
			wantErr: false,
		},
		{
			name: "success, only words and chars flag is set",
			fields: fields{
				lineCount: 3,
				wordCount: 6,
				charCount: 12,
				filename:  "test.txt",
				err:       nil,
			},
			testFlags: FlagOptions{
				LineFlag: true,
				CharFlag: true,
			},
			want:    "       3      12 test.txt\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := result{
				lineCount: tt.fields.lineCount,
				wordCount: tt.fields.wordCount,
				charCount: tt.fields.charCount,
				fileName:  tt.fields.filename,
				err:       tt.fields.err,
			}

			flagSet := tt.testFlags

			got, err := r.generateOutput(flagSet)
			if tt.wantErr {
				assert.EqualError(t, err, err.Error(), tt.fields.err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
