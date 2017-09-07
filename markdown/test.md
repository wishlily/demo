---
date: 2015-01-10 19:55
layout: 'post'
title: 'Blogging Like a Hacker'
tags:
    - "hello"
    - 'ruby'
status: 'public'
categories : [lessons, beginner]
---
# 目录
[TOC]

---
<!-- 这是一个注释 -->

# 文章信息
如开始所示，首行以`---`开始，内部使用yaml格式，并最终以`---`单独一行结束的，会最终解析为文章的元信息(描述性信息)，一般Jekyll、Hexo等静态博客用得比较多。

# 常用语法
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

粗体：**用两个星号标记起来，表示加粗**
斜体：*一个星号，表示斜体*
删除：~~这样子表示删除~~
下角标：H~2~O
上角标：2^10^

段落：
Line 1
Line 2

引用：
> 你
> 一会看我
> 一会看云

> 我觉得
> 你看我时很远
> 你看云时很近

这首诗的作者是顾城[^gucheng]的《远和近》。
其中`[^gucheng]`就是脚注的声明，你可以在文末对这个关键标注进行说明。除此之外，在标注内容上点击鼠标右键，可以让光标在标注、最终补充的(脚注)内容进行快速跳转。

[^gucheng]: 顾城，童话诗人，以及坏掉了的人。

# 表格

Item     | Value
-------- | ---
Computer | $1600
Phone    | $12
Pipe     | $1

| Item     | Value | Qty   |
| :------- | ----: | :---: |
| Computer | $1600 |  5    |
| Phone    | $12   |  12   |
| Pipe     | $1    |  234  |

|                  | ASCII                        | HTML              |
 ----------------- | ---------------------------- | ------------------
| Single backticks | `'Isn't this fun?'`            | 'Isn't this fun?' |
| Quotes           | `"Isn't this fun?"`            | "Isn't this fun?" |
| Dashes           | `-- is en-dash, --- is em-dash` | -- is en-dash, --- is em-dash |

# 链接
**插入链接:**
这是一个 [链接](http://url.com/)

**快速链接:**
只需要在网址头尾用尖括号包裹即可，比如<http://url.com>

**另外一种插链接的语法:**
这是一个 [链接][link_key]，然后，在文档的任何一行中(通常在文档末尾比较合适)，对这个链接进行定义，比如:
[link_key]: http://the_url.com/

**邮箱链接:**
这是一个 <myname@example.com> 邮箱的链接。

**常见的插图语法:**
![图片的alt信息，可空)](图片的url)

**另一种插图语法:**
![alt text][image_id]
[image_id]: 图片的url

# 列表
**无序列表:**

- Red
- Green
- Blue

**有序列表则使用数字接着一个英文句点：**

1. Bird
2. McHale
3. Parish

**也可以混合在一起使用:**

-   Bird
    - blue bird
-   McHale
    1.  a man
    2.  HoustonRockets
-   Parish

# 代码高亮

``` html
<snippet>
	<content><![CDATA[
[^${1:Footnote}]$3
[^$1]:${2:Footnote_Text}
]]></content>
	<tabTrigger>mdn</tabTrigger>
	<scope>text.html.markdown.multimarkdown, text.html.markdown</scope>
	<description>Insert Footnote</description>
</snippet>
```

# html
<div style="text-align:center; color:red">
hello world
</div>

# 扩展
## MathJax

The *Gamma function* satisfying $\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$ is via the Euler integral

$$
\Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,.
$$

## UML diagrams

You can also render sequence diagrams like this:

```sequence
Alice->Bob: Hello Bob, how are you?
Note right of Bob: Bob thinks
Bob-->Alice: I am good thanks!
```

And flow charts like this:

```flow
st=>start: Start
e=>end
op=>operation: My Operation
cond=>condition: Yes or No?

st->op->cond
cond(yes)->e
cond(no)->op
```
