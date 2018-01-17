package main

import (
  "flag"
  "github.com/fighterlyt/git-clean/clean"
)

var (
  root="/"
  exclude="/Volumes"
)
func init(){
  flag.StringVar(&root,"root","","根目录")
  flag.StringVar(&exclude,"exclude","","排除目录")
}
func main(){
  flag.Parse()
  if root==""{
    panic("root参数不能为空")
  }
  clean.Clean(root,exclude)
}
