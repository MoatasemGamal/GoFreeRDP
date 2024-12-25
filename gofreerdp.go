package gofreerdp

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
