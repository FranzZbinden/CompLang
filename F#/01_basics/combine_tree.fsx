(*
 * File: combine_three.fsx
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/10/2025
 * Purpose: This program computes and displays the result of combining three
 *          numbers using the given two-argument function.
 *)

let sum num1 num2 = num1 + num2

let minimum num1 num2 = if num1 < num2 then num1 else num2  // si las letras inferidas de tipo son iglaues, tienen que ser de distinto tipo

let combineThree num1 num2 num3 combineTwo =
    let temp = combineTwo num1 num2
    combineTwo temp num3

let number1, number2, number3 = 4, 2, 6
printfn "The numbers are %d, %d, and %d" number1 number2 number3

printfn "\nTheir sum is %d" (combineThree number1 number2 number3 sum)
printfn "\nTheir minimun is %d" (combineThree number1 number2 number3 minimum)

// Snonimous functions in f#
let maximum = fun n1 n2 -> if n1 > n2 then n1 else n2

printfn "\nTheir maximum is %d" (combineThree number1 number2 number3 maximum)
printfn "\nTheir product is %d" (combineThree number1 number2 number3 (fun n1 n2 -> n1 * n2))

printfn "\nUsing built-in operators/functions..."
printfn "\nTheir sum is %d" (combineThree number1 number2 number3 (+))
printfn "\nTheir product is %d" (combineThree number1 number2 number3 ( * ))
printfn "\nTheir product is %d" (combineThree number1 number2 number3 min)
printfn "\nTheir product is %d" (combineThree number1 number2 number3 max)


(*
// Combines three numbers using the given two-argument function.
func combineThree(num1, num2, num3 int, combineTwo func(int, int) int) int {
	temp := combineTwo(num1, num2)
	return combineTwo(temp, num3)
}

// Starts the execution of the program.
func main() {
	number1, number2, number3 := 4, 2, 6
	fmt.Printf("The numbers are %d, %d, and %d\n", number1, number2, number3)

	fmt.Println("\nTheir sum is", combineThree(number1, number2, number3, sum))
	fmt.Println("Their minimum is", combineThree(number1, number2, number3, minimum))

	// Computes the minimum of two numbers.
	maximum := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}
	fmt.Println("Their maximum is", combineThree(number1, number2, number3, maximum))

	fmt.Println("Their product is", combineThree(number1, number2, number3,
		func(n1, n2 int) int {
			return n1 * n2
		}))
}
*)