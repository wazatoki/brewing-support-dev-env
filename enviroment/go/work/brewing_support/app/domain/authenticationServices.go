package domain

import (
	"sort"
	"strings"
)

type StaffGroups []*StaffGroup

/*
FilterByString is filter staffGroup slice by string
*/
func (g *StaffGroups) FilterByString(strSlice ...string) (result StaffGroups) {
	staffGroups := *g

	if len(strSlice) == 0 {

		result = staffGroups

		return
	}

	isContains := func(staffGroup *StaffGroup, str string) bool {

		if str == "" {
			return true
		}

		return strings.Contains(staffGroup.Name, str)
	}

	for _, s := range strSlice {
		result = StaffGroups{}
		for _, group := range staffGroups {

			if isContains(group, s) {
				result = append(result, group)
			}
		}
		staffGroups = result
	}

	return

}

/*
Sort is sort staffGroup slice by name
*/
func (g *StaffGroups) Sort() StaffGroups {
	staffGroups := *g
	sort.Slice(staffGroups, func(i int, j int) bool {

		return staffGroups[i].Name < staffGroups[j].Name

	})
	return staffGroups
}
