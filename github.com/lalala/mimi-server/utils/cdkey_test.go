package utils

import "testing"

func TestCdkey(t *testing.T) {
	batch := 10000
	num := 50000
	uniqueKeys := map[string]bool{}
	keys := GenerateCdkey(batch, num)
	for _, v := range keys {
		if uniqueKeys[v] {
			t.Errorf("重复key: %s", v)
		}
		uniqueKeys[v] = true
	}
	for k, _ := range uniqueKeys {
		_batch, _no, _valid := VerifyCdkey(k)
		if !_valid {
			t.Errorf("验证失败: %s", k)
		}
		if _batch != batch {
			t.Errorf("batch不一致: %d, %d", _batch, batch)
		}
		if _no < 1 || _no > num {
			t.Errorf("no无效: %d, %d", _no, num)
		}
	}
}

func TestInvitationCode(t *testing.T) {
	t.Log("测试生成1KW邀请码是否有重复:")
	uniqueCodes := map[string]int64{}
	N := 10000 * 1000
	for i := 0; i < N; i++ {
		accountId := int64(i + 1)
		code := GenerateInvitationCode(accountId)
		v, ok := uniqueCodes[code]
		if ok {
			t.Errorf("重复code: %s, %d, %d", code, v, accountId)
		}
		uniqueCodes[code] = accountId
	}
}
