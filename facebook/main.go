package main

import (
	"fmt"
	fb "github.com/huandu/facebook"
)

func main() {
	getInvitableFriends()
}

func getInvitableFriends() {
	if res, err := fb.Get("/109129899790046/invitable_friends", fb.Params{
		"access_token": "EAAURjSwOTLoBAPCrfegcI4Pe0Bn3K9E9zZBEar7fumF6iPsteivMc4gdRihqkymqHtfIAR83MXxymP2bEUnhXRsZAQO4cq3xqgMuGKrvC39fPuZBjJp94ZAigsL9ZCZAMfwpEdLBi99vevt6YXmn2qnW7lPLaqQOPgFBi4W7TBEqrwGPxQXIu3KaHiX9o2vSTjXK18JPzSi0UiYlCGSYsTpcwUb8DuQ2sZD",
	}); err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("res:", res)
	}
}

func getPermissions() {
	if res, err := fb.Get("/109129899790046/permissions", fb.Params{
		"access_token": "EAAURjSwOTLoBAPCrfegcI4Pe0Bn3K9E9zZBEar7fumF6iPsteivMc4gdRihqkymqHtfIAR83MXxymP2bEUnhXRsZAQO4cq3xqgMuGKrvC39fPuZBjJp94ZAigsL9ZCZAMfwpEdLBi99vevt6YXmn2qnW7lPLaqQOPgFBi4W7TBEqrwGPxQXIu3KaHiX9o2vSTjXK18JPzSi0UiYlCGSYsTpcwUb8DuQ2sZD",
	}); err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("res:", res)
	}
}

func getProfile() {
	if res, err := fb.Get("/109129899790046", fb.Params{
		"fields":       "first_name,gender,picture.width(100).height(100)",
		"access_token": "EAAURjSwOTLoBAPCrfegcI4Pe0Bn3K9E9zZBEar7fumF6iPsteivMc4gdRihqkymqHtfIAR83MXxymP2bEUnhXRsZAQO4cq3xqgMuGKrvC39fPuZBjJp94ZAigsL9ZCZAMfwpEdLBi99vevt6YXmn2qnW7lPLaqQOPgFBi4W7TBEqrwGPxQXIu3KaHiX9o2vSTjXK18JPzSi0UiYlCGSYsTpcwUb8DuQ2sZD",
	}); err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("here is my facebook info:", res["first_name"])
	}

}
