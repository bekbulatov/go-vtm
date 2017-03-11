package vtm

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/go-querystring/query"
)

type atomicSwitch int64

func (r *atomicSwitch) IsSwitched() bool {
	return atomic.LoadInt64((*int64)(r)) != 0
}

func (r *atomicSwitch) SwitchOn() {
	atomic.StoreInt64((*int64)(r), 1)
}

func (r *atomicSwitch) SwitchedOff() {
	atomic.StoreInt64((*int64)(r), 0)
}

func validateID(id string) string {
	if !strings.HasPrefix(id, "/") {
		return fmt.Sprintf("/%s", id)
	}
	return id
}

func trimRootPath(id string) string {
	if strings.HasPrefix(id, "/") {
		return strings.TrimPrefix(id, "/")
	}
	return id
}

func deadline(timeout time.Duration, work func(chan bool) error) error {
	result := make(chan error)
	timer := time.After(timeout)
	stopChannel := make(chan bool, 1)

	// allow the method to attempt
	go func() {
		result <- work(stopChannel)
	}()
	for {
		select {
		case err := <-result:
			return err
		case <-timer:
			stopChannel <- true
			return ErrTimeoutError
		}
	}
}

func getInterfaceAddress(name string) (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		// step: get only the interface we're interested in
		if iface.Name == name {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}
			// step: return the first address
			if len(addrs) > 0 {
				return parseIPAddr(addrs[0]), nil
			}
		}
	}

	return "", errors.New("Unable to determine or find the interface")
}

func contains(elements []string, value string) bool {
	for _, element := range elements {
		if element == value {
			return true
		}
	}
	return false
}

func parseIPAddr(addr net.Addr) string {
	return strings.SplitN(addr.String(), "/", 2)[0]
}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
