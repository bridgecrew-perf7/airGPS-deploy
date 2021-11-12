package main

import (
	//	"encoding/json"
	"fmt"

	//"os"

	"github.com/NaySoftware/go-fcm"
)

const (
	//serverKey = "AAAAw1KydoI:APA91bETThtGh_291f_GtiCOMi6nCLiCoReIjpn1iY_600_6KikghsxKcIWb08YTQCDiZkIdY9WLUHhYSvjLzlWMYI1KU-U4LDTiiMMsihxujZ-qv1h-AL5K9ADyk9VHAT_t0e3ViMVH"
	//serverKey = "AAAAFTARtdc:APA91bERX9az700ogAIMKJTQldPd1AzFI42phDxc8j5ltJDENXRNxV3VKrQ45BT8-8avpLcV_DA4CHtNlum_n682N_ZhtPVIhtBOZRUDY6RQgqr1kMzALR59tyo-PAeqBKY48wVF0rfO"
	serverKey = "AAAA8n5S7sY:APA91bE7gUN7GoVyEUlh6sckgY9Ivaywl6BXpQHqSkQqKj4wWGE6EVESkkIrs82Y1NOw-OsTGqZVQH5dHMrmXXZoT4WK7ue558_f4XCX-ZCAhy5662nSZ4N_tuboxm7Wky8iNHneQWVB"
)

type SoundMessage struct {
	critical int
	name     string
	volume   float32
}

func checkDroid() {

	var NP fcm.NotificationPayload
	NP.Title = "Alarm"
	NP.Body = "world"

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	ids := []string{
		"c2EKbZK9TvGOcw5aD4cdqs:APA91bHxH0gZszE3e03B4Be3iK0T8v8vCOjgci5ULlSPlTl-N5siJ7np95PZzlwcqRkr5hSXD5FdgL8f5SHIFPfQNsPR1JVW_ltx5OBrxpsTzZGriaJ-wN_A5312Nl2n9K4MxLJsJdgo",
	}

	xds := []string{
		"c2EKbZK9TvGOcw5aD4cdqs:APA91bHxH0gZszE3e03B4Be3iK0T8v8vCOjgci5ULlSPlTl-N5siJ7np95PZzlwcqRkr5hSXD5FdgL8f5SHIFPfQNsPR1JVW_ltx5OBrxpsTzZGriaJ-wN_A5312Nl2n9K4MxLJsJdgo",
		"c2EKbZK9TvGOcw5aD4cdqs:APA91bHxH0gZszE3e03B4Be3iK0T8v8vCOjgci5ULlSPlTl-N5siJ7np95PZzlwcqRkr5hSXD5FdgL8f5SHIFPfQNsPR1JVW_ltx5OBrxpsTzZGriaJ-wN_A5312Nl2n9K4MxLJsJdgo",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)
	c.SetNotificationPayload(&NP)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(getDT(), "Get push status Error", err)
	}

}

func upgradeSoundTag(sound_file string) string {

	s_json := "{"
	s_json = s_json + getQuatedJSON("critical", "1", 0) + ","
	s_json = s_json + getQuatedJSON("name", sound_file, 1)
	s_json = s_json + "}"
	return sound_file
}

//https://github.com/NaySoftware/go-fcm
//https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/generating_a_remote_notification
func getTokenList(fmc_Token string, fmc_Title, fmc_category int) bool {
	var (
		send_Result bool
		ids         []string
		data        map[string]string
	)
	/// str_tokens []string
	send_Result = false
	s_json := getObjGeneral(fmc_Title)

	if fmc_category == 0 {
		data = map[string]string{
			"msg": "ALARM_ALL",
			"sum": "0",
			"key": s_json,
		}
	} else if fmc_category == 1 {
		data = map[string]string{
			"msg": "ALARM_GROUP",
			"sum": "1",
			"key": s_json,
		}

	} else if fmc_category == 2 {
		data = map[string]string{
			"msg": "ALARM_CANCEL",
			"sum": "2",
			"key": s_json,
		}
	} else if fmc_category == 3 {
		data = map[string]string{
			"msg": "CANCEL",
			"sum": "3",
			"key": s_json,
		}
	} else if fmc_category == 4 {
		data = map[string]string{
			"msg": "TROUBLE",
			"sum": "4",
			"key": s_json,
		}
	} else {
		data = map[string]string{
			"msg": "INFO",
			"sum": "5",
			"key": s_json,
		}
	}

	xds := []string{
		//	"fex7Fn4N5NA:APA91bEyFfdDxKfeLu6sXjYfuKPozsq8e6HdXJMNzeRQjuQI4CF9ub4sjBVMxEqMDp1da24itrgYFEAJjvaDCdO8s9SBFRFzSj1ogU7cDJAD_2Jc9O3xGWmSTRD_v3M8jSfENke7FyvO",
		//	"fex7Fn4N5NA:APA91bEyFfdDxKfeLu6sXjYfuKPozsq8e6HdXJMNzeRQjuQI4CF9ub4sjBVMxEqMDp1da24itrgYFEAJjvaDCdO8s9SBFRFzSj1ogU7cDJAD_2Jc9O3xGWmSTRD_v3M8jSfENke7FyvO",
	}

	ids = append(ids, fmc_Token)
	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)

	if fmc_category == 1 {
		c.SetPriority("high")
	} else {
		c.SetPriority("normal")
	}
	//c.SetDryRun(true)

	//c.SubscribeToTopic("sound", "alert")

	status, err := c.Send()

	if err == nil {
		fmt.Println(getDT(), "Send Push to tocken", fmc_Token)
		if status.Success == 1 {
			send_Result = true
			status.PrintResults()
		}
	} else {
		fmt.Println(getDT(), "Send push error", err)
		recoveryAppFunction()
		/*
			s_result = err.Error()
			if len(s_result) > 0 {
				os.Exit(0)

			}
		*/
	}

	return send_Result

}
