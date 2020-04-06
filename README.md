# Tigger

![tigger](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585213570088&di=02cc1cd15b966373b38d8fd9a7457fab&imgtype=0&src=http%3A%2F%2Fpic.qiantucdn.com%2F58pic%2F15%2F53%2F01%2F21E58PICC4X_1024.jpg%2521%2Ffw%2F1024%2Fwatermark%2Furl%2FL2ltYWdlcy93YXRlcm1hcmsvZGF0dS5wbmc%3D%2Frepeat%2Ftrue%2Fcrop%2F0x1009a0a0)

`tigger` 主要是用与 安装 `pipeline` `ide` `istio` 等应用的工具

```go
/*
执行顺序

PersistentPreRun
PreRun
Run
PostRun
PersistentPostRun

*** 设置帮助文档


*/

cmd.SetHelpCommand(cmd *Command)
cmd.SetHelpFunc(f func(*Command, []string))
cmd.SetHelpTemplate(s string)
```