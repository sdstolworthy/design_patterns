package strategy

import "fmt"

type EditBehavior interface {
	Edit()
}

type ShareBehavior interface {
	Share()
}

type AdvancedEdit struct{}

type BasicEdit struct{}

func (AdvancedEdit) Edit() {
	fmt.Println("Advanced Edit")
}

func (BasicEdit) Edit() {
	fmt.Println("Basic Edit")
}

type SocialShare struct{}

func (SocialShare) Share() {
	fmt.Println("Socially shared")
}

type PhoneCameraApp struct {
	EditBehavior  EditBehavior
	ShareBehavior ShareBehavior
}

func (a *PhoneCameraApp) SetEditBehavior(editBehavior EditBehavior) {
	a.EditBehavior = editBehavior
}
