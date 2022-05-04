package main

func main() {

	/*

		Go має лише одну циклічну конструкцію — цикл for.

		Основний цикл for складається з трьох компонентів, розділених крапкою з комою:

		оператор init: виконується перед першою ітерацією
		вираз умови: оцінюється перед кожною ітерацією
		оператор post: виконується в кінці кожної ітерації
		Інструкція init часто є короткою декларацією змінної, а змінні, оголошені там, видимі лише в області дії оператора for.

		Цикл припинить ітерацію, коли логічна умова стане хибною.

		Примітка: на відміну від інших мов, таких як C, Java або JavaScript,
		три компоненти оператора for не містять круглих дужок, а фігурні дужки { } завжди потрібні.

		Перший і останній оператори є необов'язковими.
		У цей момент ви можете відкинути крапку з комою: C while пишеться як for у Go.

	*/

	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//sum := 1
	//for ; sum < 1000; {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//for {
	//}

	/*
		if оператори в Go подібні до циклів for; вираз необов'язково оточувати круглими дужками ( ), але дужки { } є обов'язковими.
	*/

	//fmt.Println(sqrt(2), sqrt(-4))

	/*

		if з коротким оператором
		Як і у випадку for, оператор if може починатися з короткого оператора, який слід виконати перед умовою.

		Змінні, оголошені оператором, знаходяться лише в області дії до кінця if.
	*/

	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//)

	/*

		Оператор switch — це коротший спосіб написати послідовність операторів if - else.
		Він запускає перший випадок, значення якого дорівнює виразу умови.

		Перемикач Go подібний до перемикача в C, C++, Java, JavaScript і PHP, за винятком того,
		що Go запускає лише вибраний регістр, а не всі наступні.
		Фактично, оператор break, який необхідний у кінці кожного випадку цими мовами, в Go додається автоматично.

	*/

	//fmt.Print("Go runs on ")
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.\n", os)
	//}

	//fmt.Println("When's Saturday?")
	//today := time.Now().Weekday()
	//
	//switch time.Wednesday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow.")
	//case today + 2:
	//	fmt.Println("In two days.")
	//default:
	//	fmt.Println("Too far away.")
	//}

	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Good morning!")
	//case t.Hour() < 17:
	//	fmt.Println("Good afternoon.")
	//default:
	//	fmt.Println("Good evening.")
	//}

	/*
		Оператор defer відкладає виконання функції, до виходу з функції через return.

		Аргументи defer виклику оцінюються негайно, але виклик функції не виконується до return.

		https://gobyexample.com/defer#:~:text=Defer%20is%20used%20to%20ensure,be%20used%20in%20other%20languages.&text=Suppose%20we%20wanted%20to%20create,close%20when%20we're%20done.

		https://go.dev/blog/defer-panic-and-recover

		Відкладені виклики функцій поміщаються в стек. Коли функція повертається,
		її відкладені виклики виконуються в порядку last-in-first-out.
	*/

	//defer fmt.Println("world")
	//
	//fmt.Println("hello")

	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")

	/*

		Go має вказівники. Вказівник містить адресу пам'яті значення.

		Тип *T є вказівником на значення T. Його нульове значення дорівнює nil.

		var p *int

		Оператор & створює вказівник на свій операнд.

		i := 42
		p = &i

		Оператор * позначає значення вказівника.

		fmt.Println(*p) // читаємо i через вказівник p
		*p = 21 // встановлюємо i через вказівник p
	*/

	//i, j := 42, 2701
	//
	//p := &i         // point to i
	//fmt.Println(*p) // read i through the pointer
	//*p = 21         // set i through the pointer
	//fmt.Println(i)  // see the new value of i
	//
	//p = &j         // point to j
	//*p = *p / 37   // divide j through the pointer
	//fmt.Println(j) // see the new value of j

	//fmt.Println(Vertex{1, 2})

	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)

	/*

		Доступ до полів структури можна отримати через вказівник структури.

		Щоб отримати доступ до поля X структури, коли у нас є вказівник на структуру p,
		ми можемо написати (*p).X. Однак це позначення є громіздким,
		тому мова дозволяє нам замість цього писати лише p.X, без явного розіменування.

	*/

	//v := Vertex{1, 2}
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)

	//fmt.Println(v1, p, v2, v3)

	/*

		Масиви
		Тип [n]T — це масив із n значень типу T.

		Вираз

		var a [10]int

		оголошує змінну a як масив із десяти цілих чисел.

		Розміри масивів не можна змінювати, але не хвилюйтеся;
		Go забезпечує зручний спосіб роботи з масивами.

	*/

	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)

	/*
		Слайс
		Масив має фіксований розмір. З іншого боку, слайс — це гнучкий перегляд
		елементів масиву. На практиці слайси зустрічаються набагато частіше, ніж масиви.

		Тип []T є слайсом з елементами типу T.

		Слайс формується шляхом вказуванням двох індексів, нижньої та високої меж, розділених двокрапкою:

		a[low : high]
		Це визначає напіввідкритий діапазон, який включає перший елемент, але виключає останній.

		Наступний вираз створює фрагмент, який включає елементи з 1 по 3:

		a[1:4]

	*/

	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = primes[1:4]
	//fmt.Println(s)

	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)

	/*

		Літерал слайсу подібний літералу масиву без вказання довжини.

		Це літерал масиву:

		[3]bool{true, true, false}

		А цей створює той самий масив, що й вище, а потім створює слайс,
		який посилається на нього:

		[]bool{true, true, false}

	*/

	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(s)

	//s := []int{2, 3, 5, 7, 11, 13}
	//
	//fmt.Println(s[1:4])
	//fmt.Println(s[:2])
	//fmt.Println(s[1:])

	/*
		Слайс має довжину і місткість.

		Довжина слайсу — це кількість елементів, які він містить.

		Місткість слайсу — це кількість елементів у базовому масиві, рахуючи від першого
		елемента в слайсі.

		Довжину та місткість слайсу s можна отримати за допомогою виразів len(s) і cap(s).
	*/

	//s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	//
	//// Slice the slice to give it zero length.
	//s = s[:0]
	//printSlice(s)
	//
	//// Extend its length.
	//s = s[:4]
	//printSlice(s)
	//
	//// Drop its first two values.
	//s = s[2:]
	//printSlice(s)

	//Нульовий слайс
	//var s []int
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!")
	//}

	// Слайси можуть включати будь-який інший тип, в тому числі інші слайси

	// Create a tic-tac-toe board.
	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//
	//// The players take turns.
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	/*
		https://go.dev/blog/slices-intro
	*/

	//var s []int
	//printSlice(s)
	//
	//// append works on nil slices.
	//s = append(s, 0)
	//printSlice(s)
	//
	//// The slice grows as needed.
	//s = append(s, 1)
	//printSlice(s)
	//
	//// We can add more than one element at a time.
	//s = append(s, 2, 3, 4)
	//printSlice(s)

	// Створення слайсу з make

	//a := make([]int, 5)
	//printSlice("a", a)
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//
	//c := b[:2]
	//printSlice("c", c)
	//
	//d := c[2:5]
	//printSlice("d", d)

	// Range
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i) // == 2**i
	//}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}

}

//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}

//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	}
//	return lim
//}

//type Vertex struct {
//	X int
//	Y int
//}

//var (
//	v1 = Vertex{1, 2}  // has type Vertex
//	v2 = Vertex{X: 1}  // Y:0 is implicit
//	v3 = Vertex{}      // X:0 and Y:0
//	p  = &Vertex{1, 2} // has type *Vertex
//)

//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
//}
