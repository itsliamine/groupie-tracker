package utils

import "groupie-tracker/datatypes"

func GetAllLocations(json []datatypes.Location) []string {
	locs := []string{}
	for _, e := range json {
		locs = append(locs, e.Locations...)
	}
	return locs
}
