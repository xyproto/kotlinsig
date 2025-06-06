package kotlinsig

import "testing"

func TestIsKotlinFunc(t *testing.T) {
	cases := []struct {
		line     string
		expected bool
	}{
		{"fun run() {", true},
		{"private fun <T> filter(x: List<T>): List<T> {", true},
		{"Something()", false},
		{"return x", false},
		{"super.foo()", false},
		{"@Override fun compareTo(o: Any): Int {", true},
		{"main(args)", false},
		{"abstract fun <T> make(): T", true},
		{"fun renew()", true},
		{"fun toString(): String", true},
		{"throw RuntimeException()", false},
		{"interface Foo {}", false},
		{"fun main(args: Array<String>) {", true},
		{"fun <K, V> createMap(): Map<K, V>", true},
		{"val rnd = Random()", false},
		{"} catch (e: InterruptedException) {", false},
		{"fun String.extension(): Int {", true},
		{"fun add(a: Int, b: Int) = a + b", true},
		{"data class User(val name: String)", false},
		{"object Singleton {", false},
		{"enum class Color {", false},
		{"sealed class Result {", false},
		{"class MyClass {", false},
		{"fun myFunction(): Unit = println(\"hello\")", true},
		{"suspend fun fetchData(): String {", true},
		{"inline fun <reified T> create(): T {", true},
		{"tailrec fun factorial(n: Int, acc: Int = 1): Int {", true},
		{"internal fun helper() {", true},
		{"protected fun validate(): Boolean {", true},
		{"init {", true},
		{"init { println(\"Initializing\") }", true},
		{"constructor() {", true},
		{"constructor(name: String) {", true},
		{"constructor(name: String, age: Int) : this(name) {", true},
		{"primary constructor(val name: String)", true},
		{"secondary constructor() : this(\"default\") {", true},
	}

	for _, c := range cases {
		actual := Is(c.line)
		if actual != c.expected {
			t.Errorf("Is(%q) = %v; want %v", c.line, actual, c.expected)
		}
	}
}
