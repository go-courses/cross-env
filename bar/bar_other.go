// +build !linux,!darwin

package bar

func (t *Twix) Name() string {
	return "twix other"
}
