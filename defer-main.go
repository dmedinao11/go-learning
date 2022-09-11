package main

/*
Defer es un keyword que ejecuta su instrucción posterior después de que la función finalice.

Si existen múltiples defer en una sola función, se ejecutaran primero los que hayan sido llamados de último en la función.
En otras palabras, los defer se almacenan en una pila para definir su orden de ejecución

# Así, cuando finalice la ejecucón de la función el orden será

println("Printing counter in defer", 9)
println("Printing counter in defer", 8)
...
println("Execute me first after end")
println("--END--")
*/
func main() {
	defer println("--END--")
	defer println("Execute me first after end")
	println("Executeme first")

	counter := 0
	for counter < 10 {
		defer println("Printing counter in defer", counter)
		println("Printing counter", counter)
		counter++
	}
	println("Last line in main func")
}
