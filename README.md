# 牛客网 OJ 常见输入输出编程方法

官方解释：<https://www.nowcoder.com/discuss/276>

练习网站：<https://ac.nowcoder.com/acm/contest/5657/F>

## 详细信息

| 编程语言         | 版本    |
| ---------------- | ------- |
| Go               | 1.14.4  |
| JavaScript(Node) | 12.18.2 |

## 感悟

GO 海量输入处理:

- 留意样例 E，本样例中同时使用了 `bufio.NewScanner(os.Stdin)` `sc.Split(bufio.ScanWords)`，可适与大量输入和解析，具体代码示例留意对应 solution.go 文件
- 使用 `text/scanner` 来处理 token 的样例 G
- 直接读取 bytes 数组解析为整数速度最快：<https://stackoverflow.com/a/63023672/11457223>
