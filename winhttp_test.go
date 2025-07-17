//go:build windows

package winhttp

import (
	"testing"

	"golang.org/x/sys/windows"
)

func TestDetectAutoProxyConfigUrl(t *testing.T) {
	url, err := DetectAutoProxyConfigUrl(WINHTTP_AUTO_DETECT_TYPE_DHCP | WINHTTP_AUTO_DETECT_TYPE_DNS_A)
	if err != nil {
		if url != "" {
			t.Fatal("url should be empty")
		}
		if e, ok := err.(windows.Errno); !ok {
			t.Fatalf("err should be of type windows.Errno, not %T", err)
		} else if e != ERROR_WINHTTP_AUTODETECTION_FAILED {
			t.Fatalf("err should be ERROR_WINHTTP_AUTODETECTION_FAILED, not %v", e)
		}
	} else if url == "" {
		t.Fatal("url should not be empty")
	}
}
