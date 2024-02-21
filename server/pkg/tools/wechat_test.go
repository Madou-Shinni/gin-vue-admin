package tools

import "testing"

func TestGetPhoneNumber(t *testing.T) {
	result, err := GetPhoneNumber("76_Li6mmEkwFoWOAsWdkOkLrvpXY0UwUJKnRHFJwA3COHEwAoKe8GQhWtyNVqaVu2pSIOYKZbonuDlcJSsWhqawPc4CTq9bwJBioBU8syyaZaL1MdWgr3zL7glCOQoSJEiADANEZ", "123")
	if err != nil {
		t.Error(err)
		return
	}

	if result.ErrCode != 0 {
		t.Errorf("GetPhoneNumber error: %s", result.ErrMsg)
		return
	}

	t.Log(result)
}

func TestGetQRCode(t *testing.T) {
	query := map[string]interface{}{
		"question_id": 22579,
	}
	result, err := GetQRCode(
		"76_FGo32_KOd2pQgP_-j_rGZ6fuHWGxkF4fsaEIlXiYXpGK7E2bElq8OLNXATtU8X-B12AyzCL3SNiZLFdokDNlWwBbRjNa2pOxRHawxHGBzAIZR9MTF4UyO1PzT6UQVKeAFAHMB",
		"funpackage/questionsWall/questionInfo",
		query,
		0,
		false,
		nil,
		false,
		"develop", // 开发版 默认正式版
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
