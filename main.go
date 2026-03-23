package main
import (
  "flag"; "fmt"; "os/exec"
)
func main(){ ns:=flag.String("namespace","default","namespace"); name:=flag.String("name","","deployment"); image:=flag.String("image","","image"); dry:=flag.Bool("dry-run",true,"dry run"); flag.Parse(); if *name==""||*image=="" { fmt.Println("name and image required"); return }; cmd:=[]string{"set","image","deployment/"+*name,*name+"="+*image,"-n",*ns}; fmt.Println("kubectl", cmd); if *dry { return }; out,err:=exec.Command("kubectl", cmd...).CombinedOutput(); fmt.Print(string(out)); if err!=nil { panic(err) } }
