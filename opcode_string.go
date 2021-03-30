// Code generated by "stringer -type=PluginOpcode,HostOpcode -output=opcode_string.go"; DO NOT EDIT.

package vst2

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[plugOpen-0]
	_ = x[plugClose-1]
	_ = x[plugSetProgram-2]
	_ = x[plugGetProgram-3]
	_ = x[plugSetProgramName-4]
	_ = x[plugGetProgramName-5]
	_ = x[plugGetParamLabel-6]
	_ = x[plugGetParamDisplay-7]
	_ = x[plugGetParamName-8]
	_ = x[plugGetVu-9]
	_ = x[plugSetSampleRate-10]
	_ = x[plugSetBufferSize-11]
	_ = x[plugStateChanged-12]
	_ = x[PlugEditGetRect-13]
	_ = x[PlugEditOpen-14]
	_ = x[PlugEditClose-15]
	_ = x[plugEditDraw-16]
	_ = x[plugEditMouse-17]
	_ = x[plugEditKey-18]
	_ = x[PlugEditIdle-19]
	_ = x[plugEditTop-20]
	_ = x[plugEditSleep-21]
	_ = x[plugIdentify-22]
	_ = x[plugGetChunk-23]
	_ = x[plugSetChunk-24]
	_ = x[PlugProcessEvents-25]
	_ = x[PlugCanBeAutomated-26]
	_ = x[PlugString2Parameter-27]
	_ = x[plugGetNumProgramCategories-28]
	_ = x[plugGetProgramNameIndexed-29]
	_ = x[plugCopyProgram-30]
	_ = x[plugConnectInput-31]
	_ = x[plugConnectOutput-32]
	_ = x[PlugGetInputProperties-33]
	_ = x[PlugGetOutputProperties-34]
	_ = x[PlugGetPlugCategory-35]
	_ = x[plugGetCurrentPosition-36]
	_ = x[plugGetDestinationBuffer-37]
	_ = x[PlugOfflineNotify-38]
	_ = x[PlugOfflinePrepare-39]
	_ = x[PlugOfflineRun-40]
	_ = x[PlugProcessVarIo-41]
	_ = x[plugSetSpeakerArrangement-42]
	_ = x[plugSetBlockSizeAndSampleRate-43]
	_ = x[PlugSetBypass-44]
	_ = x[PlugGetPluginName-45]
	_ = x[plugGetErrorText-46]
	_ = x[PlugGetVendorString-47]
	_ = x[PlugGetProductString-48]
	_ = x[PlugGetVendorVersion-49]
	_ = x[PlugVendorSpecific-50]
	_ = x[PlugCanDo-51]
	_ = x[PlugGetTailSize-52]
	_ = x[plugIdle-53]
	_ = x[plugGetIcon-54]
	_ = x[plugSetViewPosition-55]
	_ = x[plugGetParameterProperties-56]
	_ = x[plugKeysRequired-57]
	_ = x[PlugGetVstVersion-58]
	_ = x[PlugEditKeyDown-59]
	_ = x[PlugEditKeyUp-60]
	_ = x[PlugSetEditKnobMode-61]
	_ = x[PlugGetMidiProgramName-62]
	_ = x[PlugGetCurrentMidiProgram-63]
	_ = x[PlugGetMidiProgramCategory-64]
	_ = x[PlugHasMidiProgramsChanged-65]
	_ = x[PlugGetMidiKeyName-66]
	_ = x[PlugBeginSetProgram-67]
	_ = x[PlugEndSetProgram-68]
	_ = x[PlugGetSpeakerArrangement-69]
	_ = x[PlugShellGetNextPlugin-70]
	_ = x[PlugStartProcess-71]
	_ = x[PlugStopProcess-72]
	_ = x[PlugSetTotalSampleToProcess-73]
	_ = x[PlugSetPanLaw-74]
	_ = x[PlugBeginLoadBank-75]
	_ = x[PlugBeginLoadProgram-76]
	_ = x[PlugSetProcessPrecision-77]
	_ = x[PlugGetNumMidiInputChannels-78]
	_ = x[PlugGetNumMidiOutputChannels-79]
}

const _PluginOpcode_name = "plugOpenplugCloseplugSetProgramplugGetProgramplugSetProgramNameplugGetProgramNameplugGetParamLabelplugGetParamDisplayplugGetParamNameplugGetVuplugSetSampleRateplugSetBufferSizeplugStateChangedPlugEditGetRectPlugEditOpenPlugEditCloseplugEditDrawplugEditMouseplugEditKeyPlugEditIdleplugEditTopplugEditSleepplugIdentifyplugGetChunkplugSetChunkPlugProcessEventsPlugCanBeAutomatedPlugString2ParameterplugGetNumProgramCategoriesplugGetProgramNameIndexedplugCopyProgramplugConnectInputplugConnectOutputPlugGetInputPropertiesPlugGetOutputPropertiesPlugGetPlugCategoryplugGetCurrentPositionplugGetDestinationBufferPlugOfflineNotifyPlugOfflinePreparePlugOfflineRunPlugProcessVarIoplugSetSpeakerArrangementplugSetBlockSizeAndSampleRatePlugSetBypassPlugGetPluginNameplugGetErrorTextPlugGetVendorStringPlugGetProductStringPlugGetVendorVersionPlugVendorSpecificPlugCanDoPlugGetTailSizeplugIdleplugGetIconplugSetViewPositionplugGetParameterPropertiesplugKeysRequiredPlugGetVstVersionPlugEditKeyDownPlugEditKeyUpPlugSetEditKnobModePlugGetMidiProgramNamePlugGetCurrentMidiProgramPlugGetMidiProgramCategoryPlugHasMidiProgramsChangedPlugGetMidiKeyNamePlugBeginSetProgramPlugEndSetProgramPlugGetSpeakerArrangementPlugShellGetNextPluginPlugStartProcessPlugStopProcessPlugSetTotalSampleToProcessPlugSetPanLawPlugBeginLoadBankPlugBeginLoadProgramPlugSetProcessPrecisionPlugGetNumMidiInputChannelsPlugGetNumMidiOutputChannels"

var _PluginOpcode_index = [...]uint16{0, 8, 17, 31, 45, 63, 81, 98, 117, 133, 142, 159, 176, 192, 207, 219, 232, 244, 257, 268, 280, 291, 304, 316, 328, 340, 357, 375, 395, 422, 447, 462, 478, 495, 517, 540, 559, 581, 605, 622, 640, 654, 670, 695, 724, 737, 754, 770, 789, 809, 829, 847, 856, 871, 879, 890, 909, 935, 951, 968, 983, 996, 1015, 1037, 1062, 1088, 1114, 1132, 1151, 1168, 1193, 1215, 1231, 1246, 1273, 1286, 1303, 1323, 1346, 1373, 1401}

func (i PluginOpcode) String() string {
	if i >= PluginOpcode(len(_PluginOpcode_index)-1) {
		return "PluginOpcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PluginOpcode_name[_PluginOpcode_index[i]:_PluginOpcode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HostAutomate-0]
	_ = x[HostVersion-1]
	_ = x[HostCurrentID-2]
	_ = x[HostIdle-3]
	_ = x[hostPinConnected-4]
	_ = x[hostWantMidi-6]
	_ = x[HostGetTime-7]
	_ = x[HostProcessEvents-8]
	_ = x[hostSetTime-9]
	_ = x[hostTempoAt-10]
	_ = x[hostGetNumAutomatableParameters-11]
	_ = x[hostGetParameterQuantization-12]
	_ = x[HostIOChanged-13]
	_ = x[hostNeedIdle-14]
	_ = x[HostSizeWindow-15]
	_ = x[HostGetSampleRate-16]
	_ = x[HostGetBufferSize-17]
	_ = x[HostGetInputLatency-18]
	_ = x[HostGetOutputLatency-19]
	_ = x[hostGetPreviousPlug-20]
	_ = x[hostGetNextPlug-21]
	_ = x[hostWillReplaceOrAccumulate-22]
	_ = x[HostGetCurrentProcessLevel-23]
	_ = x[HostGetAutomationState-24]
	_ = x[HostOfflineStart-25]
	_ = x[HostOfflineRead-26]
	_ = x[HostOfflineWrite-27]
	_ = x[HostOfflineGetCurrentPass-28]
	_ = x[HostOfflineGetCurrentMetaPass-29]
	_ = x[hostSetOutputSampleRate-30]
	_ = x[hostGetOutputSpeakerArrangement-31]
	_ = x[HostGetVendorString-32]
	_ = x[HostGetProductString-33]
	_ = x[HostGetVendorVersion-34]
	_ = x[HostVendorSpecific-35]
	_ = x[hostSetIcon-36]
	_ = x[HostCanDo-37]
	_ = x[HostGetLanguage-38]
	_ = x[hostOpenWindow-39]
	_ = x[hostCloseWindow-40]
	_ = x[HostGetDirectory-41]
	_ = x[HostUpdateDisplay-42]
	_ = x[HostBeginEdit-43]
	_ = x[HostEndEdit-44]
	_ = x[HostOpenFileSelector-45]
	_ = x[HostCloseFileSelector-46]
	_ = x[hostEditFile-47]
	_ = x[hostGetChunkFile-48]
	_ = x[hostGetInputSpeakerArrangement-49]
}

const (
	_HostOpcode_name_0 = "HostAutomateHostVersionHostCurrentIDHostIdlehostPinConnected"
	_HostOpcode_name_1 = "hostWantMidiHostGetTimeHostProcessEventshostSetTimehostTempoAthostGetNumAutomatableParametershostGetParameterQuantizationHostIOChangedhostNeedIdleHostSizeWindowHostGetSampleRateHostGetBufferSizeHostGetInputLatencyHostGetOutputLatencyhostGetPreviousPlughostGetNextPlughostWillReplaceOrAccumulateHostGetCurrentProcessLevelHostGetAutomationStateHostOfflineStartHostOfflineReadHostOfflineWriteHostOfflineGetCurrentPassHostOfflineGetCurrentMetaPasshostSetOutputSampleRatehostGetOutputSpeakerArrangementHostGetVendorStringHostGetProductStringHostGetVendorVersionHostVendorSpecifichostSetIconHostCanDoHostGetLanguagehostOpenWindowhostCloseWindowHostGetDirectoryHostUpdateDisplayHostBeginEditHostEndEditHostOpenFileSelectorHostCloseFileSelectorhostEditFilehostGetChunkFilehostGetInputSpeakerArrangement"
)

var (
	_HostOpcode_index_0 = [...]uint8{0, 12, 23, 36, 44, 60}
	_HostOpcode_index_1 = [...]uint16{0, 12, 23, 40, 51, 62, 93, 121, 134, 146, 160, 177, 194, 213, 233, 252, 267, 294, 320, 342, 358, 373, 389, 414, 443, 466, 497, 516, 536, 556, 574, 585, 594, 609, 623, 638, 654, 671, 684, 695, 715, 736, 748, 764, 794}
)

func (i HostOpcode) String() string {
	switch {
	case i <= 4:
		return _HostOpcode_name_0[_HostOpcode_index_0[i]:_HostOpcode_index_0[i+1]]
	case 6 <= i && i <= 49:
		i -= 6
		return _HostOpcode_name_1[_HostOpcode_index_1[i]:_HostOpcode_index_1[i+1]]
	default:
		return "HostOpcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
