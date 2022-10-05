package warehouse

import (
	"reflect"
	"testing"
	"time"

	"github.com/fullstorydev/hauser/testing/testutils"
)

var legacyColumns = []string{
	"EventCustomName",
	"EventStart",
	"EventType",
	"EventTargetText",
	"EventTargetSelectorTok",
	"EventModFrustrated",
	"EventModDead",
	"EventModError",
	"EventModSuspicious",
	"IndvId",
	"PageClusterId",
	"PageUrl",
	"PageDuration",
	"PageActiveDuration",
	"PageRefererUrl",
	"PageLatLong",
	"PageAgent",
	"PageIp",
	"PageBrowser",
	"PageDevice",
	"PageOperatingSystem",
	"PageNumInfos",
	"PageNumWarnings",
	"PageNumErrors",
	"SessionId",
	"PageId",
	"UserAppKey",
	"UserEmail",
	"UserDisplayName",
	"UserId",
	"CustomVars",
	"LoadDomContentTime",
	"LoadFirstPaintTime",
	"LoadEventTime",
}

var (
	stringType  = reflect.TypeOf("")
	int64Type   = reflect.TypeOf(int64(0))
	float64Type = reflect.TypeOf(float64(0))
	timeType    = reflect.TypeOf(time.Time{})
)

func TestSchema_ReconcileWithExisting(t *testing.T) {
	testCases := []struct {
		name   string
		cols   []string
		fields []interface{}
		expect Schema
	}{
		{
			name:   "Legacy with new columns",
			cols:   legacyColumns,
			fields: []interface{}{BaseExportFields{}},
			expect: []WarehouseField{
				{"EventCustomName", "EventCustomName", stringType},
				{"EventStart", "EventStart", timeType},
				{"EventType", "EventType", stringType},
				{"EventTargetText", "EventTargetText", stringType},
				{"EventTargetSelectorTok", "EventTargetSelectorTok", stringType},
				{"EventModFrustrated", "EventModFrustrated", int64Type},
				{"EventModDead", "EventModDead", int64Type},
				{"EventModError", "EventModError", int64Type},
				{"EventModSuspicious", "EventModSuspicious", int64Type},
				{"IndvId", "IndvId", int64Type},
				{"PageClusterId", "PageClusterId", int64Type},
				{"PageUrl", "PageUrl", stringType},
				{"PageDuration", "PageDuration", int64Type},
				{"PageActiveDuration", "PageActiveDuration", int64Type},
				{"PageRefererUrl", "PageRefererUrl", stringType},
				{"PageLatLong", "PageLatLong", stringType},
				{"PageAgent", "PageUserAgent", stringType},
				{"PageIp", "PageIp", stringType},
				{"PageBrowser", "PageBrowser", stringType},
				{"PageDevice", "PageDevice", stringType},
				{"PageOperatingSystem", "PageOperatingSystem", stringType},
				{"PageNumInfos", "PageNumInfos", int64Type},
				{"PageNumWarnings", "PageNumWarnings", int64Type},
				{"PageNumErrors", "PageNumErrors", int64Type},
				{"SessionId", "SessionId", int64Type},
				{"PageId", "PageId", int64Type},
				{"UserAppKey", "UserAppKey", stringType},
				{"UserEmail", "UserEmail", stringType},
				{"UserDisplayName", "UserDisplayName", stringType},
				{"UserId", "UserId", int64Type},
				{"CustomVars", "CustomVars", stringType},
				{"LoadDomContentTime", "LoadDomContentTime", int64Type},
				{"LoadFirstPaintTime", "LoadFirstPaintTime", int64Type},
				{"LoadEventTime", "LoadEventTime", int64Type},
				{"UserCreated", "UserCreated", timeType},
				{"EventSubType", "EventSubType", stringType},
				{"EventTargetSelector", "EventTargetSelector", stringType},
				{"EventPageOffset", "EventPageOffset", int64Type},
				{"EventSessionOffset", "EventSessionOffset", int64Type},
				{"EventWebSourceFileUrl", "EventWebSourceFileUrl", stringType},
				{"EventFirstInputDelay", "EventFirstInputDelay", int64Type},
				{"EventCumulativeLayoutShift", "EventCumulativeLayoutShift", float64Type},
				{"SessionStart", "SessionStart", timeType},
				{"PageStart", "PageStart", timeType},
				{"PageBrowserVersion", "PageBrowserVersion", stringType},
				{"PagePlatform", "PagePlatform", stringType},
				{"PageScreenWidth", "PageScreenWidth", int64Type},
				{"PageScreenHeight", "PageScreenHeight", int64Type},
				{"PageViewportWidth", "PageViewportWidth", int64Type},
				{"PageViewportHeight", "PageViewportHeight", int64Type},
				{"PageMaxScrollDepthPercent", "PageMaxScrollDepthPercent", int64Type},
				{"LoadLargestPaintTime", "LoadLargestPaintTime", int64Type},
				{"ReqUrl", "ReqUrl", stringType},
				{"ReqMethod", "ReqMethod", stringType},
				{"ReqStatus", "ReqStatus", int64Type},
			},
		},
		{
			name:   "brand new schema, with mobile apps",
			cols:   []string{},
			fields: []interface{}{BaseExportFields{}, MobileFields{}},
			expect: []WarehouseField{
				{"IndvId", "IndvId", int64Type},
				{"UserId", "UserId", int64Type},
				{"SessionId", "SessionId", int64Type},
				{"PageId", "PageId", int64Type},
				{"UserCreated", "UserCreated", timeType},
				{"UserAppKey", "UserAppKey", stringType},
				{"UserDisplayName", "UserDisplayName", stringType},
				{"UserEmail", "UserEmail", stringType},
				{"EventStart", "EventStart", timeType},
				{"EventType", "EventType", stringType},
				{"EventSubType", "EventSubType", stringType},
				{"EventCustomName", "EventCustomName", stringType},
				{"EventTargetText", "EventTargetText", stringType},
				{"EventTargetSelector", "EventTargetSelector", stringType},
				{"EventPageOffset", "EventPageOffset", int64Type},
				{"EventSessionOffset", "EventSessionOffset", int64Type},
				{"EventModFrustrated", "EventModFrustrated", int64Type},
				{"EventModDead", "EventModDead", int64Type},
				{"EventModError", "EventModError", int64Type},
				{"EventModSuspicious", "EventModSuspicious", int64Type},
				{"EventWebSourceFileUrl", "EventWebSourceFileUrl", stringType},
				{"EventFirstInputDelay", "EventFirstInputDelay", int64Type},
				{"EventCumulativeLayoutShift", "EventCumulativeLayoutShift", float64Type},
				{"SessionStart", "SessionStart", timeType},
				{"PageStart", "PageStart", timeType},
				{"PageDuration", "PageDuration", int64Type},
				{"PageActiveDuration", "PageActiveDuration", int64Type},
				{"PageUrl", "PageUrl", stringType},
				{"PageRefererUrl", "PageRefererUrl", stringType},
				{"PageIp", "PageIp", stringType},
				{"PageLatLong", "PageLatLong", stringType},
				{"PageUserAgent", "PageUserAgent", stringType},
				{"PageBrowser", "PageBrowser", stringType},
				{"PageBrowserVersion", "PageBrowserVersion", stringType},
				{"PageDevice", "PageDevice", stringType},
				{"PagePlatform", "PagePlatform", stringType},
				{"PageOperatingSystem", "PageOperatingSystem", stringType},
				{"PageScreenWidth", "PageScreenWidth", int64Type},
				{"PageScreenHeight", "PageScreenHeight", int64Type},
				{"PageViewportWidth", "PageViewportWidth", int64Type},
				{"PageViewportHeight", "PageViewportHeight", int64Type},
				{"PageNumInfos", "PageNumInfos", int64Type},
				{"PageNumWarnings", "PageNumWarnings", int64Type},
				{"PageNumErrors", "PageNumErrors", int64Type},
				{"PageClusterId", "PageClusterId", int64Type},
				{"PageMaxScrollDepthPercent", "PageMaxScrollDepthPercent", int64Type},
				{"LoadDomContentTime", "LoadDomContentTime", int64Type},
				{"LoadEventTime", "LoadEventTime", int64Type},
				{"LoadFirstPaintTime", "LoadFirstPaintTime", int64Type},
				{"LoadLargestPaintTime", "LoadLargestPaintTime", int64Type},
				{"ReqUrl", "ReqUrl", stringType},
				{"ReqMethod", "ReqMethod", stringType},
				{"ReqStatus", "ReqStatus", int64Type},
				{"CustomVars", "CustomVars", stringType},
				{"AppName", "AppName", stringType},
				{"AppPackageName", "AppPackageName", stringType},
				{"AppDeviceModel", "AppDeviceModel", stringType},
				{"AppDeviceVendor", "AppDeviceVendor", stringType},
				{"AppVersion", "AppVersion", stringType},
				{"AppOsVersion", "AppOsVersion", stringType},
				{"AppViewName", "AppViewName", stringType},
				{"EventMobileSourceFile", "EventMobileSourceFile", stringType},
			},
		},
		{
			name:   "someone added some columns",
			fields: []interface{}{BaseExportFields{}},
			cols:   []string{"preexisting", "columns", "userid"},
			expect: []WarehouseField{
				{"preexisting", "", nil},
				{"columns", "", nil},
				{"UserId", "UserId", int64Type},
				{"IndvId", "IndvId", int64Type},
				{"SessionId", "SessionId", int64Type},
				{"PageId", "PageId", int64Type},
				{"UserCreated", "UserCreated", timeType},
				{"UserAppKey", "UserAppKey", stringType},
				{"UserDisplayName", "UserDisplayName", stringType},
				{"UserEmail", "UserEmail", stringType},
				{"EventStart", "EventStart", timeType},
				{"EventType", "EventType", stringType},
				{"EventSubType", "EventSubType", stringType},
				{"EventCustomName", "EventCustomName", stringType},
				{"EventTargetText", "EventTargetText", stringType},
				{"EventTargetSelector", "EventTargetSelector", stringType},
				{"EventPageOffset", "EventPageOffset", int64Type},
				{"EventSessionOffset", "EventSessionOffset", int64Type},
				{"EventModFrustrated", "EventModFrustrated", int64Type},
				{"EventModDead", "EventModDead", int64Type},
				{"EventModError", "EventModError", int64Type},
				{"EventModSuspicious", "EventModSuspicious", int64Type},
				{"EventWebSourceFileUrl", "EventWebSourceFileUrl", stringType},
				{"EventFirstInputDelay", "EventFirstInputDelay", int64Type},
				{"EventCumulativeLayoutShift", "EventCumulativeLayoutShift", float64Type},
				{"SessionStart", "SessionStart", timeType},
				{"PageStart", "PageStart", timeType},
				{"PageDuration", "PageDuration", int64Type},
				{"PageActiveDuration", "PageActiveDuration", int64Type},
				{"PageUrl", "PageUrl", stringType},
				{"PageRefererUrl", "PageRefererUrl", stringType},
				{"PageIp", "PageIp", stringType},
				{"PageLatLong", "PageLatLong", stringType},
				{"PageUserAgent", "PageUserAgent", stringType},
				{"PageBrowser", "PageBrowser", stringType},
				{"PageBrowserVersion", "PageBrowserVersion", stringType},
				{"PageDevice", "PageDevice", stringType},
				{"PagePlatform", "PagePlatform", stringType},
				{"PageOperatingSystem", "PageOperatingSystem", stringType},
				{"PageScreenWidth", "PageScreenWidth", int64Type},
				{"PageScreenHeight", "PageScreenHeight", int64Type},
				{"PageViewportWidth", "PageViewportWidth", int64Type},
				{"PageViewportHeight", "PageViewportHeight", int64Type},
				{"PageNumInfos", "PageNumInfos", int64Type},
				{"PageNumWarnings", "PageNumWarnings", int64Type},
				{"PageNumErrors", "PageNumErrors", int64Type},
				{"PageClusterId", "PageClusterId", int64Type},
				{"PageMaxScrollDepthPercent", "PageMaxScrollDepthPercent", int64Type},
				{"LoadDomContentTime", "LoadDomContentTime", int64Type},
				{"LoadEventTime", "LoadEventTime", int64Type},
				{"LoadFirstPaintTime", "LoadFirstPaintTime", int64Type},
				{"LoadLargestPaintTime", "LoadLargestPaintTime", int64Type},
				{"ReqUrl", "ReqUrl", stringType},
				{"ReqMethod", "ReqMethod", stringType},
				{"ReqStatus", "ReqStatus", int64Type},
				{"CustomVars", "CustomVars", stringType},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			baseSchema := MakeSchema(tc.fields...)
			updatedSchema := baseSchema.ReconcileWithExisting(tc.cols)
			testutils.Assert(t, tc.expect.Equals(updatedSchema), "wrong schema:\nwant %#v,\ngot  %#v", tc.expect, updatedSchema)
		})
	}
}
