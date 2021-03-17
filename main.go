package main

import (
	"fmt"
	"os"
	"io"
)

/*   ====== This is main ============   */
func main(){
	filename:= "slovar.dat"

	var slovar map[string]string
	slovar = make(map[string]string)
	//var p_sl  *map[string]string
	//p_sl = &slovar

	fmt.Println("This is a program for learning English by word cards")

	var i int
	for i != 9 {
	    fmt.Println("Select <1> - to record <2> - to view")
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

/* File creation function*/
func createFileWords(filename string){
    file ,err := os.Create(filename)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	} 
	defer file.Close()
}

/*Random word generator function*/
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

	fmt.Println("Enter a word in English")
	fmt.Scan(&engword)

	fmt.Println("Enter translation in Russian")
	fmt.Scan(&rusword)

	sl[engword]=rusword
	
	file, err := os.OpenFile(filename, os.O_RDWR,0660)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	defer file.Close()
	
	for k, v := range sl{
	    fmt.Println("I write to file")
	    fmt.Fprintf(file, "%s\t%s\n", k, v)
	}
}
/* View all words from a file */
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

