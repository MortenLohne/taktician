// Code generated by "stringer -type=Feature"; DO NOT EDIT

package ai

import "fmt"

const _Feature_name = "TempoTopFlatStandingCapstoneHardTopCapCapMobilityFlatCaptives_SoftFlatCaptives_HardStandingCaptives_SoftStandingCaptives_HardCapstoneCaptives_SoftCapstoneCaptives_HardLibertiesGroupLibertiesGroupsGroups_1Groups_2Groups_3Groups_4Groups_5Groups_6Groups_7Groups_8PotentialThreatEmptyControlFlatControlCenterCenterControlThrowMineThrowTheirsThrowEmptyTerminal_PliesTerminal_FlatsTerminal_ReservesTerminal_OpponentReservesMaxFeature"

var _Feature_index = [...]uint16{0, 5, 12, 20, 28, 38, 49, 66, 83, 104, 125, 146, 167, 176, 190, 196, 204, 212, 220, 228, 236, 244, 252, 260, 269, 275, 287, 298, 304, 317, 326, 337, 347, 361, 375, 392, 417, 427}

func (i Feature) String() string {
	if i < 0 || i >= Feature(len(_Feature_index)-1) {
		return fmt.Sprintf("Feature(%d)", i)
	}
	return _Feature_name[_Feature_index[i]:_Feature_index[i+1]]
}
