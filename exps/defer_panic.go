package exps

func TestDeferPanic(){
	defer panic("error in defer")
}
