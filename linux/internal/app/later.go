package app

import "github.com/rabden/nanolight/linux/internal/apps"

// Filled in by later phases.

func (a *App) startSettings()            {}
func (a *App) settingsChanged()          {}
func (a *App) startWelcome()             {}
func (a *App) welcomeShowing()           {}
func (a *App) welcomePressed(bool)       {}
func (a *App) startTray()                {}
func (a *App) startUpdates()             {}
func (a *App) updateAway()               {}
func (a *App) startAccount()             {}
func (a *App) maybePull()                {}
func (a *App) touched(bool)              {}
func (a *App) raycastRows() []apps.Entry { return nil }
func (a *App) openRaycast(apps.Entry) bool {
	return false
}
