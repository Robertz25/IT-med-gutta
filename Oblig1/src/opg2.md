**Oppgave 2a** 

	func Bubble_sort_modified(list []int) {
	n := len(list)
	swapped := true
	for swapped {
	swapped = false
	for i := 1; i < n-1; i++ {
	if list[i-1] > list[i] {

	list[i], list[i-1] = list[i-1], list[i]
	swapped = true
	}
	}
	}
	}

**Oppgave 2b**
  
  **Forstå algoritmer og utføre “benchmark”-tester på koden**
  
  ```
  func BenchmarkBSortModified100(b *testing.B) {
	benchmarkBSortModified(100, b)
}

func BenchmarkBSortModified1000(b *testing.B) {
	benchmarkBSortModified(1000, b)
}

func BenchmarkBSortModified10000(b *testing.B) {
	benchmarkBSortModified(10000, b)
}

func benchmarkBSortModified(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort_modified(values)
    }
 
 ```
**Oppgave 2c**

```
BenchmarkBSortModified100-4                30000             41181 ns/op
BenchmarkBSortModified1000-4                 500           2656742 ns/op
BenchmarkBSortModified10000-4                  3         340763500 ns/op
BenchmarkBSort100-4                        30000             37841 ns/op
BenchmarkBSort1000-4                        1000           2077786 ns/op
BenchmarkBSort10000-4                          5         298664220 ns/op
BenchmarkQSort100-4                       200000              8549 ns/op
BenchmarkQSort1000-4                       10000            115897 ns/op
BenchmarkQSort10000-4                       1000           1473063 ns/op

27.752s
```
Det vi kan si om big-O for alle 3 algoritmene som vi har testet er at space complexity er større eller verre for BubbleSort og BubbleSortModified, enn QuickSort. For de tre forskjellige testene med ulike elementer ser vi at QuickSort kjører gjennom langt flere looper enn BubbleSortene. Samme gjelder for time complexity der QuickSort går gjennom hver loop mye raskere, med for eksempel 8549 millisekund per loop med 100 elementer, mens Bubblesort bruker 37841 millisekund per loop.

Hvis vi ser på tabellen over Array Sorting Algorithms på: http://bigocheatsheet.com kan vi se at generelt presterer QuickSort algoritmen bedre enn BubbleSorten. Jo flere operasjoner og elementer vi får vil begge prestere dårligere, men QuickSort ser ut til å være bedre sammenlignet med BubbleSort.
