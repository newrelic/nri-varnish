package args

import (
	"testing"
)

func Test_ArgumentList_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		args    *ArgumentList
		wantErr bool
	}{
		{
			"Valid",
			&ArgumentList{
				InstanceName: "my_instance",
			},
			false,
		},
		{
			"Invalid",
			&ArgumentList{},
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.args.Validate()
		if tc.wantErr && err == nil {
			t.Fatal("Expected error")
		}

		if !tc.wantErr && err != nil {
			t.Fatalf("Unexpected error %s", err.Error())
		}
	}
}
