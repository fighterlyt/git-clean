package clean

import (
  "path/filepath"
  "os"
  "fmt"
  "os/exec"
)
var (
  gits=[]string{}
  exclude=""
)
func sClean(root,e string) {
  exclude=e
  filepath.Walk(root,Record)
  for _,git:=range gits{
    fmt.Printf("git gc %s\n",git)
    if err:=gc(git);err!=nil{
      fmt.Printf("gc出错[%s] [%s]\n",git,err.Error())
    }
    fmt.Printf("git prune %s\n",git)

    if err:=prune(git);err!=nil{
      fmt.Printf("prune出错[%s] [%s]\n",git,err.Error())
    }
  }
}

func Record(path string,info os.FileInfo,err error) error{
  if path==exclude{
    return filepath.SkipDir
  }
  if filepath.Base(path)==".git"&&info.IsDir(){
    fmt.Printf("找到git目录%s\n",filepath.Dir(path))
    gits=append(gits,filepath.Dir(path))
    return filepath.SkipDir
  }
  return nil
}

func gc(path string) error{
  cmd:=exec.Command("git","gc")
  cmd.Dir=path
  cmd.Stdout=os.Stdout
  return cmd.Run()
}
func prune(path string) error{
  cmd:=exec.Command("git","prune")
  cmd.Dir=path
  cmd.Stdout=os.Stdout
  return cmd.Run()
}
