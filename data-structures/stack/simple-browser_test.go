package stack

import "testing"

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.PushBackStack("www.qq.com")
	b.PushBackStack("www.baidu.com")
	b.PushBackStack("www.sina.com")
	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.taobao.com")
	if b.CanForward() {
		b.Forward()
	}
}
