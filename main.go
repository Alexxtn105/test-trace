package main

// По материалам:
// https://kovardin.ru/articles/go/go-tool-trace/
import (
	"os"
	"runtime/trace" // Необходимо для запуска трейсинга
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Чтобы можно было добраться до событий рантайма в вашем бинарнике, начинаем
	// Это позволить сохранят все события, произошедшие в программе, в бинарный файл trace.out.
	// После этого можно запускать go tool trace trace.out.
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// ... вся остальная логика вашей программы
}
