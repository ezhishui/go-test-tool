# go-test-tool
golang单元测试工具

使用golang的过程中，执行命令总是出现文件进程被占用，而且写go test执行单元测试时，测试方法写好长。希望列出所有可执行的方法直接选择执行即可，不用自己输入一长串的东西，达到这种效果：

$ gm chdir<br>
1 ---- logf_test.go----TestSysLog1<br>
2 ---- logf_test.go----TestSlice1<br>
3 ---- logfactory_test.go----TestSysLog<br>
4 ---- logfactory_test.go----TestSlice<br>
5 ---- regex_test.go----TestRegEx<br>
6 ---- regex_test.go----TestRegEx1<br><br>
7 ---- regex_test.go----Test111ifdska1<br>
8 ---- regex_test.go----Testfff288f1<br>
9 ---- regex_test.go----Test111ifdska1<br>
10 ---- regex_test.go----Testfff288f1<br>
请选择要执行的编号：<br>
1<br>
_----------------E:/go3/go-test-tool/logf_test.go----TestSysLog1<br>
=== RUN   TestSysLog1<br>
2019/07/03 20:19:06 sssssssssssssssssssssss<br>
--- PASS: TestSysLog1 (0.09s)<br>
PASS<br>
ok      command-line-arguments  (cached)<br>
