# go-test-tool
golang单元测试工具

使用golang的过程中，执行命令总是出现文件进程被占用，而且写go test执行单元测试时，测试方法写好长。希望列出所有可执行的方法直接选择执行即可，不用自己输入一长串的东西，达到这种效果：

$ gm chdir
1 ---- logf_test.go----TestSysLog1 \n
2 ---- logf_test.go----TestSlice1 \n
3 ---- logfactory_test.go----TestSysLog
4 ---- logfactory_test.go----TestSlice
5 ---- regex_test.go----TestRegEx
6 ---- regex_test.go----TestRegEx1
7 ---- regex_test.go----Test111ifdska1
8 ---- regex_test.go----Testfff288f1
9 ---- regex_test.go----Test111ifdska1
10 ---- regex_test.go----Testfff288f1
请选择要执行的编号：
1
_----------------E:/go3/go-test-tool/logf_test.go----TestSysLog1
=== RUN   TestSysLog1
2019/07/03 20:19:06 sssssssssssssssssssssss
--- PASS: TestSysLog1 (0.09s)
PASS
ok      command-line-arguments  (cached)
