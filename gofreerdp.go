package gofreerdp

import (
	"errors"
	"net"
	"os/exec"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	Option_aero                  = "aero"
	Option_asyncChannels         = "async-channels"
	Option_asyncUpdate           = "async-update"
	Option_authOnly              = "auth-only"
	Option_authentication        = "authentication"
	Option_autoReconnect         = "auto-reconnect"
	Option_compression           = "compression"
	Option_credentialsDelegation = "credentials-delegation"
	Option_decorations           = "decorations"
	Option_drives                = "drives"
	Option_encryption            = "encryption"
	Option_fipsMode              = "fipsmode"
	Option_fonts                 = "fonts"
	Option_forceConsoleCallbacks = "force-console-callbacks"
	Option_gestures              = "gestures"
	Option_grabKeyboard          = "grab-keyboard"
	Option_grabMouse             = "grab-mouse"
	Option_heartbeat             = "heartbeat"
	Option_homeDrive             = "home-drive"
	Option_menuAnims             = "menu-anims"
	Option_mouseMotion           = "mouse-motion"
	Option_mouseRelative         = "mouse-relative"
	Option_multiTouch            = "multitouch"
	Option_multiTransport        = "multitransport"
	Option_nego                  = "nego"
	Option_oldLicense            = "old-license"
	Option_suppressOutput        = "suppress-output"
	Option_printReconnectCookie  = "print-reconnect-cookie"
	Option_themes                = "themes"
	Option_toggleFullscreen      = "toggle-fullscreen"
	Option_unmapButtons          = "unmap-buttons"
	Option_wallpaper             = "wallpaper"
	Option_windowDrag            = "window-drag"
)

var optionsMap = map[string]bool{
	Option_aero:                  false,
	Option_asyncChannels:         false,
	Option_asyncUpdate:           false,
	Option_authOnly:              false,
	Option_authentication:        true,
	Option_autoReconnect:         false,
	Option_compression:           true,
	Option_credentialsDelegation: false,
	Option_decorations:           true,
	Option_drives:                false,
	Option_encryption:            true,
	Option_fipsMode:              false,
	Option_fonts:                 true,
	Option_forceConsoleCallbacks: false,
	Option_gestures:              false,
	Option_grabKeyboard:          true,
	Option_grabMouse:             true,
	Option_heartbeat:             true,
	Option_homeDrive:             false,
	Option_menuAnims:             false,
	Option_mouseMotion:           true,
	Option_mouseRelative:         false,
	Option_multiTouch:            false,
	Option_multiTransport:        true,
	Option_nego:                  true,
	Option_oldLicense:            false,
	Option_suppressOutput:        true,
	Option_printReconnectCookie:  false,
	Option_themes:                true,
	Option_toggleFullscreen:      true,
	Option_unmapButtons:          false,
	Option_wallpaper:             true,
	Option_windowDrag:            false,
}

type RDPConfig struct {
	Addr     string `validate:"required,hostname|ip"` // Addr is required, and must be a valid hostname or IP address
	Username string `validate:"required"`             // Username is required
	Password string `validate:"required"`             // Password is required
}

// freeRDP struct and methods
type freeRDP struct {
	freeRDP   string // it may be xfreerdp or xfreerdp3 based on your system
	config    *RDPConfig
	options   map[string]bool
	arguments []string
}

// Declare singleton instance
var (
	instance *freeRDP
	once     sync.Once
)

func Init() (*freeRDP, error) {
	xfreerdp, err := checkDependencies()
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		// Initialize the freerdp struct when first accessed
		instance = &freeRDP{
			freeRDP: xfreerdp,
			options: optionsMap,
		}
	})
	return instance, nil
}

// Methods
func (freerdp *freeRDP) SetConfig(rdpConfig *RDPConfig) error {
	validate := validator.New()
	err := validate.Struct(rdpConfig)
	if err != nil {
		return err
	}

	// Set the config if all validations pass
	freerdp.config = rdpConfig
	return nil
}

func (freerdp *freeRDP) CheckServerAvailability(timeout time.Duration) error {
	if timeout == 0 {
		timeout = 2 * time.Second
	}
	address := freerdp.config.Addr

	// Attempt to dial the address with the specified timeout
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return errors.New("port is not available")
	}
	conn.Close()

	return nil
}

// Arguments methods
func (freerdp *freeRDP) AddArg(argName, argValue string) *freeRDP {
	if argValue != "" {
		freerdp.arguments = append(freerdp.arguments, "/"+argName+":"+argValue)
		return freerdp
	}
	freerdp.arguments = append(freerdp.arguments, "/"+argName)
	return freerdp
}

func (freerdp *freeRDP) DynamicResolution() *freeRDP {
	return freerdp.AddArg("dynamic-resolution", "")
}

func (freerdp *freeRDP) MultiMonitor() *freeRDP {
	return freerdp.AddArg("multimon", "force")
}

func (freerdp *freeRDP) FullScreenMode() *freeRDP {
	return freerdp.AddArg("f", "")
}

func (freerdp *freeRDP) App(app *SeamlessApp) *freeRDP {
	return freerdp.AddArg(app.toString(), "")
}

// helpers
func checkDependencies() (string, error) {
	commands := []string{"xfreerdp3", "xfreerdp"}

	for _, cmdName := range commands {
		cmd := exec.Command("which", cmdName)
		if err := cmd.Run(); err == nil {
			return cmdName, nil
		}
	}
	return "", errors.New("freerdp is not installed")
}
