package main

import (
	"fmt"
	"github.com/danyelkeddah/go-design-patterns/abstract_factory"
	"github.com/danyelkeddah/go-design-patterns/adapter"
	"github.com/danyelkeddah/go-design-patterns/bridge"
	"github.com/danyelkeddah/go-design-patterns/builder"
	"github.com/danyelkeddah/go-design-patterns/chain_of_responsibility"
	"github.com/danyelkeddah/go-design-patterns/command"
	"github.com/danyelkeddah/go-design-patterns/composite"
	"github.com/danyelkeddah/go-design-patterns/decorator"
	"github.com/danyelkeddah/go-design-patterns/facade"
	"github.com/danyelkeddah/go-design-patterns/iterator"
	"github.com/danyelkeddah/go-design-patterns/mediator"
	"github.com/danyelkeddah/go-design-patterns/memento"
	"github.com/danyelkeddah/go-design-patterns/observer"
	"github.com/danyelkeddah/go-design-patterns/prototype"
	"github.com/danyelkeddah/go-design-patterns/proxy"
	"github.com/danyelkeddah/go-design-patterns/singleton"
	"github.com/danyelkeddah/go-design-patterns/state"
	"github.com/danyelkeddah/go-design-patterns/strategy"
	"github.com/danyelkeddah/go-design-patterns/template_method"
	"github.com/danyelkeddah/go-design-patterns/visitor_pattern"
	"log"
)

func main() {
	visitorMain()
}

func abstractFactoryMain() {
	// brand should be eiter (adidas or nike)
	sportsFactory, err := abstract_factory.MakeSportsFactory("nike")
	if err != nil {
		fmt.Println(err)
		return
	}

	sportsFactory.MakeShoe()
}

func builderMain() {
	normalBuilder, _ := builder.MakeBuilder("normal")
	iglooBuilder, _ := builder.MakeBuilder("igloo")
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Println(normalHouse)
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Println(iglooHouse)
}

func prototypeMain() {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}
	folder1 := &prototype.Folder{
		Name:      "Folder1",
		Childrens: []prototype.Inode{file1},
	}
	folder2 := &prototype.Folder{
		Name:      "Folder2",
		Childrens: []prototype.Inode{folder1, file2, file3},
	}

	fmt.Println("printing hierarchy for folder 2")
	folder2.Print("    ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nprinting hierarchy for clone Folder")
	cloneFolder.Print("    ")
}

func singletonMain() {
	for i := 0; i < 100; i++ {
		go singleton.GetDatabaseInstance()
	}
	fmt.Scanln()
}

func adapterMain() {
	client := &adapter.Client{}
	mac := &adapter.Mac{}
	client.InsertSquareUsbInComputer(mac)
	windows := &adapter.Windows{}
	windowsAdapter := &adapter.WindowsAdapter{WindowsMachine: windows}
	client.InsertSquareUsbInComputer(windowsAdapter)
}

func bridgeMain() {
	hpPrinter := &bridge.HpPrinter{}
	epsonPrinter := &bridge.EpsonPrinter{}
	macComputer := &bridge.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	windowsComputer := &bridge.Windows{}
	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()
}

func compositeMain() {
	file1 := &composite.File{Name: "File1"}
	file2 := &composite.File{Name: "File2"}
	file3 := &composite.File{Name: "File3"}
	folder1 := &composite.Folder{Name: "Folder1"}
	folder1.Add(file1)
	folder2 := &composite.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)
	folder2.Search("rose")
}

func decoratorMain() {
	veggiePizza := &decorator.VeggeMania{}
	// add cheese
	veggiePizzaWithCheese := &decorator.CheeseTopping{Pizza: veggiePizza}
	veggiePizzaWithCheeseAndTomato := &decorator.TomatoTopping{Pizza: veggiePizzaWithCheese}
	fmt.Printf("Prize for veggie mania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.GetPrice())

	peppyPaneerPizza := &decorator.PeppyPaneer{}

	//Add cheese topping
	peppyPaneerPizzaWithCheese := &decorator.CheeseTopping{
		Pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.GetPrice())
}

func facadeMain() {
	wallet := facade.NewWallet("abc123", 1234)
	fmt.Println()
	err := wallet.AddMoneyToWallet("abc123", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	err = wallet.DeductMoneyFromWallet("ab", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

func flyweightMain() {

}

func proxyMain() {
	nginxServer := proxy.NewNginxServer()
	appStatusUrl := "/app/status"
	createUserUrl := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)
}

func chainOfResponsibilityMain() {
	patient := &chain_of_responsibility.Patient{
		Name: "Jack",
	}
	reception := &chain_of_responsibility.Reception{}
	doctor := &chain_of_responsibility.Doctor{}
	medical := &chain_of_responsibility.Medical{}
	cashier := &chain_of_responsibility.Cashier{}
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)
	reception.Execute(patient)
}

func commandMain() {
	tv := &command.Tv{}
	onCommand := &command.OnCommand{Device: tv}   // action
	offCommand := &command.OffCommand{Device: tv} // action

	onButton := &command.Button{Command: onCommand}
	onButton.Press()

	offButton := &command.Button{Command: offCommand}
	offButton.Press()
}

func iteratorMain() {
	user1 := &iterator.User{
		Name: "a",
		Age:  15,
	}
	user2 := &iterator.User{
		Name: "b",
		Age:  10,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}
	iteratorObj := userCollection.CreateIterator()
	for iteratorObj.HasNext() {
		user := iteratorObj.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

func mediatorMain() {
	stationManager := mediator.NewStationManager()
	passengerTrain := &mediator.PassengerTrain{Mediator: stationManager}
	goodsTrain := &mediator.GoodsTrain{Mediator: stationManager}
	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}

func mementoMain() {
	careTaker := &memento.Caretaker{}
	originator := &memento.Originator{State: "initial editor text"}
	fmt.Printf("state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())
	originator.SetState("change 1")
	fmt.Printf("state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())
	originator.SetState("chagne 2")
	fmt.Printf("state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())
	fmt.Println("++++++ undo everything ++++++")
	originator.RestoreMemento(careTaker.GetMemento(2))
	fmt.Printf("state: %s\n", originator.GetState())
	originator.RestoreMemento(careTaker.GetMemento(1))
	fmt.Printf("state: %s\n", originator.GetState())
	originator.RestoreMemento(careTaker.GetMemento(0))
	fmt.Printf("state: %s\n", originator.GetState())
}
func observerMain() {
	shirtItem := observer.NewItem("Nike Shirt")
	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}

func stateMain() {
	vendingMachine := state.NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("")
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
}

func strategyMain() {
	leastFrequentlyUsed := &strategy.LeastFrequentlyUsed{}
	leastRecentlyUsed := &strategy.LeastRecentlyUsed{}
	firstInFirstOut := &strategy.FirstInFirstOut{}
	cache := strategy.NewCache(leastFrequentlyUsed)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")
	cache.SetDeletionAlgorithm(leastRecentlyUsed)
	cache.Add("d", "4")
	cache.SetDeletionAlgorithm(firstInFirstOut)
	cache.Add("e", "5")
}

func templateMethodMain() {
	smsOtp := &template_method.Sms{}
	o := template_method.Otp{Otp: smsOtp}
	err := o.GenerateAndSendOtp(4)
	if err != nil {
		fmt.Println(err)
	}
	emailOtp := &template_method.Email{}
	o = template_method.Otp{Otp: emailOtp}
	err = o.GenerateAndSendOtp(4)
	if err != nil {
		fmt.Println(err)
	}
}

func visitorMain() {
	square := &visitor_pattern.Square{}
	circle := &visitor_pattern.Circle{}
	rectangle := &visitor_pattern.Rectangle{}

	areaCalculator := &visitor_pattern.AreaCalculator{}
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
	fmt.Println()
	middleCoordinates := &visitor_pattern.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
