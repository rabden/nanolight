package app

import (
	"github.com/rabden/nanolight/linux/internal/apps"
	"github.com/rabden/nanolight/linux/internal/view"
)

// Filled in by later phases.

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

func (a *App) trayChanged()                    {}
func (a *App) showTray(bool)                   {}
func (a *App) accountState() view.AccountState { return view.SignedOut }
func (a *App) accountEmail() string            { return "" }
func (a *App) signIn()                         {}
func (a *App) signOut()                        {}
func (a *App) updateReady() string             { return "" }
func (a *App) checkUpdates()                   {}
func (a *App) restartToUpdate()                {}
func (a *App) showWelcome()                    {}
