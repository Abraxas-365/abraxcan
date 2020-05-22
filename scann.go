package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
	//"strconv"
	"os/exec"
	"sync"	
)

//CREAR FUNCIONES PARA LO QUE SE REPITE Y SEA MAS FACIL LEER
//USAR - SET PARA SETEAR LAS COSAS , REESCRIBIR LA PARTE DE SWITCH

func main(){


	var esperar sync.WaitGroup
	var esperar2 sync.WaitGroup
	ip := " "
	puert := 0
	puertos :=[]string{}
	mas_usados :=[]string{"21", "22", "80", "443" , "445", "3389"}

	cmd := bufio.NewScanner(os.Stdin)

	fmt.Print("abraxcan> ")


	for cmd.Scan(){
		imput := cmd.Text()
		componentes := strings.Split(imput, " ")
		comand := componentes[0]
		//cambiar esto para que joja el nombre del usuario
		workspace_path :="/home/abraxas"


		switch comand{





		case "help"	:
			fmt.Println("ip 	para set ip")
			fmt.Println("puertos   y separacio con comas para los puertos")
			fmt.Println("todos      para todos los puertos")
			fmt.Println("run 		para comenzar")
		case "ip":
				ip = componentes[1]
		case "puertos":
				puertos = strings.Split(componentes[1], ",")
				puert = 1
				
				
		case "run":
				if puert == 1 {
						for x :=0; x<len(puertos);x++ {
								direccion := ip + ":" + puertos[x]
								si, err := net.Dial("tcp", direccion )
								if err ==nil {
									fmt.Println(" ")
									fmt.Println("[+]" + puertos[x] + " abierto")
									si.Close()
							}else {
									
									continue
							}	
						}

				}//acaba el primer if
				if puert == 0 {     //no olvodarme de cambair esto a un numero(eficiencia)
						for x :=0; x<len(mas_usados);x++ {
							esperar.Add(1)
							go func(r int){
								defer esperar.Done()
								direccion := ip + ":" + mas_usados[r]
								si2, err := net.Dial("tcp", direccion )
								if err ==nil {
									fmt.Println("[+]" + mas_usados[r] + " abierto")
									si2.Close()
								}else {
									
									return
								}	
							}(x)
						}
						esperar.Wait()

				}

		case "todos":
			for x :=0; x<=1024;x++ {
				
				esperar2.Add(1)
				go func(j int){
					
					defer esperar2.Done()
					num :=fmt.Sprintf(ip + ":%d",j)
					
					si3, err := net.Dial("tcp", num )
					if err ==nil {
						si3.Close()
						fmt.Println("[+]" + num + " abierto")
						
					}else{
						
						return
					}
				}(x)
			}
			esperar2.Wait()


		case "exit":
			fmt.Println("[+]Volviendo a la normalidad...")
            os.Exit(0)

        



       	case "workspace":
       		carpeta := componentes[1]
        	_, err := exec.Command("mkdir",workspace_path,carpeta).CombinedOutput()
       		
        	if err != nil {
        		os.Stderr.WriteString(err.Error())
        	}
        	fmt.Println("[+]Workspace creado")


        default:
        	output, err := exec.Command(comand).CombinedOutput()
        	if err != nil {
        		os.Stderr.WriteString(err.Error())
        	}
        	fmt.Println(string(output))
	}


	fmt.Print("abraxcan> ")
	
	

	}
}
