package main

import strategy "github.com/sdstolworthy/design_patterns/strategy"

func runStrategyDuck() {
	duck := strategy.Duck{QuackBehavior: new(strategy.LoudQuack), FlyBehavior: new(strategy.DiveFly)}
	duck.Quack()
	duck.Fly()
}

func runStrategyPhoto() {
	basicCamerApp := strategy.PhoneCameraApp{EditBehavior: new(strategy.BasicEdit), ShareBehavior: new(strategy.SocialShare)}
	basicCamerApp.EditBehavior.Edit()
	basicCamerApp.ShareBehavior.Share()
	basicCamerApp.SetEditBehavior(new(strategy.AdvancedEdit))
	basicCamerApp.EditBehavior.Edit()
}

func main() {
	runStrategyDuck()
	runStrategyPhoto()
}
