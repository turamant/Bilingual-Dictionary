package main

import (
	"fmt"
	"os"
	"io"
)


func main(){
	
	filename:= "slovar.dat"
	
	
	var slovar map[string]string
	//var p_sl  *map[string]string
	
    	fmt.Println("Это программа изучения анг.языка по карточкам слов")
	slovar = make(map[string]string)
	
	
	//p_sl = &slovar
	
	var i int
	for i != 9 {
	    fmt.Println("Выберите <1> - для записи <2> - для просмотра")
	    fmt.Scan(&i)
            switch i {
		case 0:
		        createFileWords(filename)
		        break
		case 1:
			recordWords(slovar, filename)
			break
		case 2:
			viewWords(slovar, filename)
			break
		case 3:
		        randWords(filename)
		        break	
		default:
			fmt.Println("Default switch!")
		}
	}
	fmt.Println("Good bye!!!")

}

func createFileWords(filename string){
    file ,err := os.Create(filename)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	} 
	defer file.Close()
    
}

func randWords(filename string){
	var engword string
	var rusword string
	
	var slov_rand map[string]string
	slov_rand = make(map[string]string)
	
	
	file, err := os.Open(filename)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	defer file.Close()
	
	for{
	    _, err = fmt.Fscanf(file,"%s\t%s\n", &engword, &rusword )
	    if err != nil{
	        if err == io.EOF{
	            break
	            }else{
	            fmt.Println(err)
	            os.Exit(1)
	            }
	        }
	        slov_rand[engword]=rusword
	    }
	} 




func recordWords(sl map[string]string, filename string){
        var engword, rusword string

	fmt.Println("Введите слово на английском")
	fmt.Scan(&engword)

	fmt.Println("Введите перевод на русском")
	fmt.Scan(&rusword)

	sl[engword]=rusword
	
	file, err := os.OpenFile(filename, os.O_RDWR,0660)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	defer file.Close()
	
	for k, v := range sl{
	    fmt.Println("Пишу в файл")
	    fmt.Fprintf(file, "%s\t%s\n", k, v)
	}
}

func viewWords(sl map[string]string, filename string){
	var engword string
	var rusword string
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666) 
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	defer file.Close()
	
	for{
	    _, err = fmt.Fscanf(file,"%s\t%s\n", &engword, &rusword )
	    if err != nil{
	        if err == io.EOF{
	            break
	            }else{
	            fmt.Println(err)
	            os.Exit(1)
	            }
	        }
	        fmt.Printf("%s\t%s\n", engword, rusword)
	    }
	} 

