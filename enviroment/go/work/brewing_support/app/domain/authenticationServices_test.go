package domain

import (
	"reflect"
	"testing"
)

func createExpectedStaff1Entity() *Staff {
	return &Staff{
		ID:          "staffid1",
		AccountID:   "12345",
		Name:        "name 1",
		Password:    "password 1",
		StaffGroups: createExpectedStaffGroupEntity1Slice(),
	}
}

func createExpectedStaff2Entity() *Staff {
	return &Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		StaffGroups: StaffGroups{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaffGroup1Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid1",
		Name: "staff group name 1",
	}
}

func createExpectedStaffGroup2Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid2",
		Name: "staff group name 2",
	}
}

func createExpectedStaffGroup3Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid3",
		Name: "staff group name 3",
	}
}

func createExpectedStaffGroupEntity1Slice() StaffGroups {
	return StaffGroups{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func createExpectedStaffGroupEntity2Slice() StaffGroups {
	return StaffGroups{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
		createExpectedStaffGroup3Entity(),
	}
}

func TestStaffGroups_Sort(t *testing.T) {
	tests := []struct {
		name string
		g    *StaffGroups
		want StaffGroups
	}{
		{
			name: "sort by staffGroup name",
			g: &StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
			want: StaffGroups{
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup3Entity(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaffGroups.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStaffGroups_FilterByString(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name       string
		g          *StaffGroups
		args       args
		wantResult StaffGroups
	}{
		{
			name: "filter without strings",
			g: &StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
			args: args{
				strSlice: []string{},
			},
			wantResult: StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
		},
		{
			name: "filter with ''",
			g: &StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
			args: args{
				strSlice: []string{},
			},
			wantResult: StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
		},
		{
			name: "filter with '1'",
			g: &StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup3Entity(),
			},
			args: args{
				strSlice: []string{"1"},
			},
			wantResult: StaffGroups{
				createExpectedStaffGroup1Entity(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.g.FilterByString(tt.args.strSlice...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("StaffGroups.FilterByString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
