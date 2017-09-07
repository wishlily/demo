# [MathJax 基本用法](http://waterbolik.github.io/jekyll/2015/10/14/MathJax%E8%AF%AD%E6%B3%95%E8%AF%A6%E8%A7%A3)

## 上标与下标

运算符      | 说明   | 代码                             | 示----例
:------- | :--- | :----------------------------- | :------------------------------:
^        | 上标   | x ^ y                          | $x ^ y$
\_       | 下标   | x _ y                          | $x _ y$
\sideset | 四周标记 | \sideset{^1_2}{^3_4}\bigotimes | $\sideset{^1_2}{^3_4}\bigotimes$

1. 表示排列使用{n+1 \choose 2k} 或 \binom{n+1}{2k}。
2. ${n+1 \choose 2k}$ $\binom{n+1}{2k}$

## 括号

运算符             | 说明  | 代码                            | 示-------例
:-------------- | :-- | :---------------------------- | :-----------------------------:
(和)             | 小括号 | (2+3)                         | $(2+3)$
[和]             | 中括号 | [4+4]                         | $[4+4]$
{和}             | 大括号 | \{a*b\}                       | $\{a*b\}$
<和>             | 尖括号 | \langle x+y \rangle           | $\langle x+y \rangle$
\lceil和\rceil   | 上取整 | \lceil\frac{1}{2}\rceil = 1   | $\lceil\frac{1}{2}\rceil = 1$
\lfloor和\rfloor | 下取整 | \lfloor\frac{1}{2}\rfloor = 0 | $\lfloor\frac{1}{2}\rfloor = 0$

## 分式与根式

运算符   | 说明       | 代码          | 示例
:---- | :------- | :---------- | :-----------:
\frac | -分式      | \frac{x}{y} | $\frac{x}{y}$
\over | -分式      | {x}\over{y} | ${x}\over{y}$
\sqrt | x√‾开二次方  | \sqrt x     | $\sqrt x$
\sqrt | x√n‾开n次方 | \sqrt[n]{x} | $\sqrt[n]{x}$

## 表格

$$
  \begin{array}
  {c | l c r}
  n & \text{Left} & \text{Center} & \text{Right} \newline
  \hline
  1 & 0.24 & 1 & 125 \newline
  2 & -1 & 189 & -8 \newline
  3 & -20 & 2000 & 1+10i \newline
  \end{array}
$$

### 对齐的公式

$$
  \begin{align}
  \sqrt{37}
  & = \sqrt{\frac{73^2-1}{12^2}} \newline
  & = \sqrt{\frac{73^2}{12^2}\cdot\frac{73^2-1}{73^2}} \newline
  & = \sqrt{\frac{73^2}{12^2}}\sqrt{\frac{73^2-1}{73^2}} \newline
  & = \frac{73}{12}\sqrt{1 - \frac{1}{73^2}} \newline
  & \approx \frac{73}{12}\left(1 - \frac{1}{2\cdot73^2}\right)
  \end{align}
$$


### 分类表达式

$$
  f(n) =
  \begin{cases}
  n/2, & \text{if $n$ is even} \newline
  [5ex] 3n+1, & \text{if $n$ is odd} \newline
  \end{cases}
$$

$$
  \left.
  \begin{array}
  {l}
  \text{if $n$ is even:} & n/2 \\
  \text{if $n$ is odd:} & 3n+1
  \end{array}
  \right\}
  =f(n)
$$

### 使用 `\mid` 代替 | 作为分隔符

$$
  \begin{array}
  {c | c}
  \mathrm{Bad} & \mathrm{Better} \newline
  \hline \
  { x | x^2 \in \Bbb Z } & { x \mid x^2 \in \Bbb Z } \newline
  \end{array}
$$

## 方程组

$$
  \left.
  \begin{array}
  {c}
  a_1x+b_1y+c_1z=d_1 \\
  a_2x+b_2y+c_2z=d_2 \\
  a_3x+b_3y+c_3z=d_3
  \end{array}
  \right\}
$$

$$
  \begin{cases}
  a_1x+b_1y+c_1z=d_1 \\
  a_2x+b_2y+c_2z=d_2 \\
  a_3x+b_3y+c_3z=d_3
  \end{cases}
$$

$$
  \left\{
  \begin{aligned}
  a_1x+b_1y+c_1z & = d_1+e_1 \\
  a_2x+b_2y & = d_2 \\
  a_3x+b_3y+c_3z & = d_3
  \end{aligned}
  \right.
$$

$$
  \left\{
  \begin{array}
  {l l}
  a_1x+b_1y+c_1z & = d_1+e_1 \\
  a_2x+b_2y & = d_2 \\
  a_3x+b_3y+c_3z & = d_3
  \end{array}
  \right.
$$

## 矩阵

### 基本用法

$$
  \begin{matrix}
  1 & x & x^2 \\
  1 & y & y^2 \\
  1 & z & z^2 \\
  \end{matrix}
$$

### 加括号

$$
  \begin{pmatrix}
  1 & 2 \\
  3 & 4 \\
  \end{pmatrix}
$$

$$
  \begin{bmatrix}
  1 & 2 \\
  3 & 4 \\
  \end{bmatrix}
$$

$$
  \begin{Bmatrix}
  1 & 2 \\
  3 & 4 \\
  \end{Bmatrix}
$$

$$
  \begin{vmatrix}
  1 & 2 \\
  3 & 4 \\
  \end{vmatrix}
$$

$$
  \begin{Vmatrix}
  1 & 2 \\
  3 & 4 \\
  \end{Vmatrix}
$$

### 省略元素

$$
  \begin{pmatrix}
  1 & a_1 & a_1^2 & \cdots & a_1^n \\
  1 & a_2 & a_2^2 & \cdots & a_2^n \\
  \vdots & \vdots& \vdots & \ddots & \vdots \\
  1 & a_m & a_m^2 & \cdots & a_m^n
  \end{pmatrix}
$$

### 增广矩阵

$$
  \left[
  \begin{array}
  {c c | c}
  1 & 2 & 3 \\
  4 & 5 & 6 \\
  \end{array}
  \right]
$$

## 算术运算符

运算符        | 说明    | 代码                    | 示-------------------例
:--------- | :---- | :-------------------------- | :-----------------------------:
\+         | +加    | x + y                       | $x + y$
\-         | -减    | x - y                       | $x - y$
\times     | ×乘    | x \times y                  | $x \times y$
\cdot      | ⋅乘    | x \cdot y                   | $x \cdot y$
\ast       | ∗乘    | x \ast y                    | $x \ast y$
\div       | ÷除    | x \div y                    | $x \div y$
\pmode     | mod取模 | a \equiv b \pmod n          | $a \equiv b \pmod n$
\pm        | ±加减   | x \pm y                     | $x \pm y$
\mp        | ∓减加   | x \mp y                     | $x \mp y$
\=         | =等于   | x = y                       | $x = y$
\mid       | ∣是除数|                             | $x \mid y$
\nmid      | ∤不是除数|                            | $x \nmid y$
\sum       | ∑连加求和 | sum_{i=0}^n frac{1}{i^2}    | $\sum_{i=0}^n \frac{1}{i^2}$
\prod      | ∏连乘求积 | prod_{i=0}^n frac{1}{i^2}   | $\prod_{i=0}^n \frac{1}{i^2}$
\coprod    | ∐     | coprod_{i=0}^n frac{1}{i^2} | $\coprod_{i=0}^n \frac{1}{i^2}$
\oplus     | ⊕圆加   | x \oplus y                  | $x \oplus y$
\odot      | ⨀圆点   | x \odot y                   | $x \odot y$
\otimes    | ⨂圆乘   | x \otimes y                 | $x \otimes y$
\bigoplus  | ⨁圆加   | x \bigoplus y               | $x \bigoplus y$
\bigodot   | ⨀圆点   | x \bigodot y                | $x \bigodot y$
\bigotimes | ⨂圆乘   | x \bigotimes y              | $x \bigotimes y$

## 比较运算符

运算符     | 说明     | 代码                   | 示----------例
:------ | :----- | :------------------- | :--------------------:
\=      | =等于    | x = y                | $x = y$
\neq    | ≠不等于   | x \neq y \not= z     | $x \neq y \not= z$
\<\lt   | <小于    | x < y \lt z          | $x < y \lt z$
\not\<  | ≮不小于   | x \not\< y \not\lt z | $x \not< y \not\lt z$
\leq    | ≤小于等于  | x \leq y             | $x \leq y$
\nleq   | ≰不小于等于 | x \nleq y \not\leq z | $x \nleq y \not\leq z$
\>      | >大于    | x > y \gt z          | $x > y\gt z$
\not\>  | ≯不大于   | x \not> y \not\gt z  | $x \not\> y \not\gt z$
\geq    | ≥大于等于  | x \geq y             | $x \geq y$
\ngeq   | ≱不大于等于 | x \ngeq y \not\geq z | $x \ngeq y \not\geq z$
\approx | ≈约等于   | x \approx y          | $x \approx y$
\equiv  | ≡恒等于   | x \equiv y           | $x \equiv y$
\sim    | ∼      | \sim                 | $\sim$
\cong   | ≅      | \cong                | $\cong$
\prec   | ≺      | \prec                | $\prec$

## 集合运算符

运算符           | 说明                     | 代码                | 示----例
:------------ | :--------------------- | :---------------- | :-----------------:
\emptyset     | ∅空集                    | \emptyset         | $\emptyset$
\varnothing   | ∅空集                    | \varnothing       | $\varnothing$
\in           | ∈属于                    | x \in y           | $x \in y$
\notin        | ∉不属于                   | x \notin y        | $x \notin y$
\subset       | ⊂子集                    | x \subset y       | $x \subset y$
\not\subset   | ⊄非子集                   | x \not\subset y   | $x \not\subset y$
\subseteq     | ⊆子等集                   | x \subseteq y     | $x \subseteq y$
\not\subseteq | ⊈非子等集                  | x \not\subseteq y | $x \not\subseteq y$
\supset       | ⊃超集                    | x \supset y       | $x \supset y$
\not\supset   | ⊅非超集                   | x \not\supset y   | $x \not\supset y$
\supseteq     | ⊇超等集                   | x \supseteq y     | $x \supseteq y$
\not\supseteq | ⊉非超等集                  | x \not\supseteq y | $x \not\supseteq y$
\cup          | ∪并                     | x \cup y          | $x \cup y$
\not\cup      | ∪̸非并                   | x \not\cup y      | $x \not\cup y$
\cap          | ∩交                     | x \cap y          | $x \cap y$
\not\cap      | ∩̸非交                   | x \not\cap y      | $x \not\cap y$
\vee          | ∨合取                    | x \vee y          | $x \vee y$
\not\vee      | ∨̸非合取                  | x \not\vee y      | $x \not\vee y$
\wedge        | ∧析取                    | x \wedge y        | $x \wedge y$
\not\wedge    | ∧̸非析取                  | x \not\wedge y    | $x \not\wedge y$
\uplus        | ⊎                      | x \uplus y        | $x \uplus y$
\not\uplus    | ⊎̸                     | x \not\uplus y    | $x \not\uplus y$
\sqcup        | ⊔                      | x \sqcup y        | $x \sqcup y$
\not\sqcup    | ⊔̸                     | x \not\sqcup y    | $x \not\sqcup y$
\bigcup       | ⋃大并                    | x \bigcup y       | $x \bigcup y$
\not\bigcup   | ⧸⋃大非并                  | x \not\bigcup y   | $x \not\bigcup y$
\bigcap       | ⋂大交                    | x \bigcap y       | $x \bigcap y$
\not\bigcap   | ⧸⋂大非交                  | x \not\bigcap y   | $x \not\bigcap y$
\bigvee       | ⋁命题的“合取”（“与”）运算        | x \bigvee y       | $x \bigvee y$
\not\bigvee   | ⧸⋁命题的“合取”（“与”）运算       | x \not\bigvee y   | $x \not\bigvee y$
\bigwedge     | ⋀命题的“析取”（“或”，“可兼或”）运算  | x \bigwedge y     | $x \bigwedge y$
\not\bigwedge | ⧸⋀命题的“析取”（“或”，“可兼或”）运算 | x \not\bigwedge y | $x \not\bigwedge y$
\biguplus     | ⨄                      | x \biguplus y     | $x \biguplus y$
\not\biguplus | ⧸⨄                     | x \not\biguplus y | $x \not\biguplus y$
\bigsqcup     | ⨆                      | x \bigsqcup y     | $x \bigsqcup y$
\not\bigsqcup | ⧸⨆                     | x \not\bigsqcup y | $x \not\bigsqcup y$

## 对数运算符

运算符  | 说明    | 代码      | 示----例
:--- | :---- | :------ | :-------:
\log | log对数 | \log(x) | $\log(x)$
\lg  | lg对数  | \lg(x)  | $\lg(x)$
\ln  | ln对数  | \ln(x)  | $\ln(x)$

## 三角运算符

运算符     | 说明        | 代码         | 示---------例
:------ | :-------- | :--------- | :----------:
\bot    | ⊥垂直       | A \bot B   | $A \bot B$
\angle  | ∠角        | \angle 45  | $\angle 45$
\circ   | ∘度        | 45^\circ   | $45^\circ$
\sin    | sin正弦     | \sin 45    | $\sin 45$
\cos    | cos余弦     | \cos 45    | $\cos 45$
\tan    | tan正切     | \tan 45    | $\tan 45$
\arcsin | arcsin反正弦 | \arcsin 45 | $\arcsin 45$
\arccos | arccos反余弦 | \arccos 45 | $\arccos 45$
\arctan | arctan反正切 | \arctan 45 | $\arctan 45$
\cot    | cot       | \cot       | $\cot$
\sec    | sec       | \sec       | $\sec$
\csc    | csc       | \csc       | $\csc$

## 微积分运算符

运算符           | 说明          | 代码                              | 示------------------例
:------------ | :---------- | :------------------------------ | :-------------------------------:
\prime        | ′           |                                 | $\prime$
\int          | ∫积分         | \int_0^1 x^2 {\rm d}x           | $\int_0^1 x^2 {\rm d}x$
\iint         | ∬二重积分       | \iint_D f(x,y)d\sigma           | $\iint_D f(x,y)d\sigma$
\iiint        | ∭三重积分       | \iiint_D f(x,y)d\sigma          | $\iiint_D f(x,y)d\sigma$
\iiiint       | ⨌四重积分       | \iiiint_D f(x,y)d\sigma         | $\iiiint_D f(x,y)d\sigma$
\oint         | ∮闭合曲面（曲线）积分 | \oint e^{x+y} ds                | $\oint e^{x+y} ds$
\lim          | lim极限       | \lim*{x\to\infty}               | $\lim*{x\to\infty}$
\infty        | ∞极限         | \sum_{i=0}^\infty i^2           | $\sum_{i=0}^\infty i^2$
\nabla        | ∇           |                                 | $\nabla$
\partial      | ∂部分         | \frac{\partial x}{\partial y}   | $\frac{\partial x}{\partial y}$
\displaystyle | 块公式格式       | \displaystyle \lim*{x\to\infty} | $\displaystyle \lim*{x\to\infty}$


$$
  \begin{array}
  {c | c}
  \mathrm{Bad} & \mathrm{Better} \\
  \hline \\
  \int\int_S f(x) \, dy \, dx & \iint_S f(x) \, dy \, dx \\
  \int\int\int_V f(x) \, dz \, dy \, dx & \iiint_V f(x) \, dz \, dy \, dx \\
  \int\int\int\int_V f(x) \, dz \, dy \, dx \, dt & \iiiint_V f(x) \, dz \, dy \, dx \, dt
  \end{array}
$$

$$
  \begin{array}
  {c | c}
  \mathrm{Bad} & \mathrm{Better} \\
  \hline \\
  \iiint_V f(x)dz dy dx & \iiint_V f(x) \, dz \, dy \, dx
  \end{array}
$$

## 逻辑运算符

运算符        | 说明    | 代码         | 示例
:--------- | :---- | :--------- | :----------:
\because   | ∵因为   | \because   | $\because$
\therefore | ∴所以   | \therefore | $\therefore$
\land      | ∧     | \land      | $\land$
\lor       | ∨     | \lor       | $\lor$
\lnot      | ¬     | \lnot      | $\lnot$
\forall    | ∀全称量词 | \forall    | $\forall$
\exists    | ∃存在量词 | \exists    | $\exists$
\top       | ⊤     | \top       | $\top$
\bot       | ⊥     | \bot       | $\bot$
\vdash     | ⊢     | \vdash     | $\vdash$
\vDash     | ⊨     | \vDash     | $\vDash$

## 顶部符号 `&` 连线符 `.` 号

运算符             | 说明      | 代码                                    | 示------------------例
:-------------- | :------ | :------------------------------------------- | :-------------------------------------:
\hat            | ŷ       | \hat{xyz}                                    | $\hat{xyz}$
\hat            | Ŷ拟合值    | \hat Y = \hat \beta_0 + \hat \beta_1X        | $\hat Y = \hat \beta_0 + \hat \beta_1X$
\vec            | a→向量   | \vec a + \vec b = \vec c                     | $\vec a + \vec b = \vec c$
\vec            | abc→向量  | \vec {abc}                                   | $\vec{abc}$
\widehat        | xyzˆ    | \widehat{xyz}                                | $\widehat{xyz}$
\check          | yˇ      | \check{xyz}                                  | $\check{xyz}$
\breve          | y˘      | \breve{xyz}                                  | $\breve{xyz}$
\overline       | ⎯⎯⎯平均数  | \overline{x}                                 | $\overline{x}$
\overline       | ⎯⎯⎯连线符号 | \overline{a+b+c} +d                          | $\overline{a+b+c}+d$
\underline      | ⎯⎯⎯下划线  | a+\underline{b+c}+d                          | $a+\underline{b+c}+d$
\overrightarrow | y→      | \overrightarrow{y}                           | $\overrightarrow{y}$
\overleftarrow  | y←      | \overleftarrow{y}                            | $\overleftarrow{y}$
\dot            | y˙      | \dot{xyz}                                    | $\dot{xyz}$
\ddot           | y¨      | \ddot{xyz}                                   | $\ddot{xyz}$
\overbrace      | ⏞上大括号   | \overbrace{a+\underbrace{b+c}*{1.5}+d}^{2.0} | $\overbrace{a+d}^{2.0}$
\underbrace     | ⏟下大括号   | \underbrace{b+c}*{1.5}                       | $\underbrace{b+c}*{1.5}$

## 箭头符号

运算符             | 说明     | 代码              | 示---例
:-------------- | :----- | :-------------- | :---------------:
\to             | →右箭头   | \to             | $\to$
\mapsto         | ↦左顶右箭头 | \mapsto         | $\mapsto$
\uparrow        | ↑上箭头   | \uparrow        | $\uparrow$
\Uparrow        | ⇑上箭头   | \Uparrow        | $\Uparrow$
\downarrow      | ↓下箭头   | \downarrow      | $\downarrow$
\Downarrow      | ⇓下箭头   | \Downarrow      | $\Downarrow$
\leftarrow      | ←左箭头   | \leftarrow      | $\leftarrow$
\Leftarrow      | ⇐左箭头   | \Leftarrow      | $\Leftarrow$
\longleftarrow  | ⟵左箭头   | \longleftarrow  | $\longleftarrow$
\Longleftarrow  | ⟸左箭头   | \Longleftarrow  | $\Longleftarrow$
\rightarrow     | →右箭头   | \rightarrow     | $\rightarrow$
\Rightarrow     | ⇒右箭头   | \Rightarrow     | $\Rightarrow$
\longrightarrow | ⟶右箭头   | \longrightarrow | $\longrightarrow$
\Longrightarrow | ⟹右箭头   | \Longrightarrow | $\Longrightarrow$

## 其他

运算符      | 说明       | 代码                      | 示-----------------------例
:------- | :------- | :----------------------------- | :------------------------------:
\ldots   | 底端对齐的省略号 | 1,2,\ldots,n                   | $1,2,\ldots,n$
\cdots   | 中线对齐的省略号 | x_1^2 + x_2^2 + \cdots + x_n^2 | $x_1^2 + x_2^2 + \cdots + x_n^2$
\vdots   | 竖对齐的省略号  | 1,2,\vdots,n                   | $1,2,\vdots,n$
\ddots   | 矩阵对齐的省略号 | 1,2,\ddots,n                   | $1,2,\ddots,n$
\star    | ⋆五角星     | \star                          | $\star$
\ast     | ∗雪花      | \ast                           | $\ast$
\circ    | ∘圆点      | \circ                          | $\circ$
\bullet  | ∙实着重号    | \bullet                        | $\bullet$
\bigstar | ⋆五角星     | \bigstar                       | $\bigstar$
\bigcirc | ∘圆点      | \bigcirc                       | $\bigcirc$
\aleph   | ℵ        | \aleph                         | $\aleph$
\Im      | ℑ        | \Im                            | $\Im$
\Re      | ℜ        | \Re                            | $\Re$

## 附录

### 希腊字母

序号 | 希腊字母大写 | 语法       | 希腊字母小写 | 语法          | 中文名称   | 示例                         | 意义
-: | :----: | :------- | :----: | :---------: | :----: | :------------------------: | :--------------------:
 1 | A      | A        | α      | \alhpa      | 阿尔法    | $A \qquad \alpha$          | 角度，系数，角加速度
 2 | B      | B        | β      | \beta       | 贝塔     | $B \qquad \beta$           | 磁通系数，角度，系数
 3 | Γ      | \Gamma   | γ      | \gamma      | 伽马     | $\Gamma \qquad \gamma$     | 电导系数，角度，比热容比
 4 | Δ      | \Delta   | δ      | \delta      | 德尔塔    | $\Delta \qquad \delta$     | 变化量，屈光度，一元二次方程中的判别式
 5 | E      | E        | ϵ      | \epsilon    | 伊普西隆   | $E \qquad \epsilon$        | 对数之基数，介电常数
 6 | Z      | Z        | ζ      | \zeta       | 泽塔     | $Z \qquad \zeta$           | 系数，方位角，阻抗，相对粘度
 7 | H      | H        | η      | \eta        | 伊塔     | $H \qquad \eta $           | 迟滞系数，效率
 8 | Θ      | \Theta   | θ      | \theta      | 西塔     | $\Theta \qquad \theta$     | 温度，角度
 9 | I      | I        | ι      | \iota       | 约塔     | $I \qquad \iota$           | 微小，一点
10 | K      | K        | κ      | \kappa      | 卡帕     | $K \qquad \kappa$          | 介质常数，绝热指数
11 | Λ      | \Lambda  | λ      | \lambda     | 兰姆达    | $\Lambda \qquad \lambda$   | 波长，体积，导热系数
12 | M      | M        | μ      | \mu         | 谬      | $M \qquad \mu$             | 磁导系数，微，动摩擦系（因）数，流体动力粘度
13 | N      | N        | ν      | \nu         | 纽      | $N \qquad \nu$             | 磁阻系数，流体运动粘度,光子频率
14 | Ξ      | \Xi      | ξ      | \xi         | 克西     | $\Xi \qquad \xi$           | 随机数，（小）区间内的一个未知特定值
15 | O      | O        | ο      | \omicron    | 欧米克隆   | $O \qquad \omicron$        | 高阶无穷小函数
16 | Π      | \Pi      | π      | \pi         | 派      | $\Pi \qquad \pi$           | 圆周率，π(n)表示不大于n的质数个数
17 | R      | R        | ρ      | \rho        | 柔      | $P \qquad \rho$            | 电阻系数，柱坐标和极坐标中的极径，密度
18 | Σ      | \Sigma   | σ      | \sigma      | 西格玛    | $\Sigma \qquad \sigma$     | 总和，表面密度，跨导，正应力
19 | T      | T        | τ      | \tau        | 陶      | $T \qquad \tau$            | 时间常数，切应力
20 | Υ      | \Upsilon | υ      | \upsilon    | 宇普西隆   | $\Upsilon \qquad \upsilon$ | 位移
21 | Φ      | \Phi     | ϕ      | \phi        | 弗爱     | $\Phi \qquad \phi$         | 磁通，角，透镜焦度，热流量
22 | X      | X        | χ      | \chi        | 卡      | $X \qquad \chi$            | 统计学中有卡方(χ^2)分布
23 | Ψ      | \Psi     | ψ      | \psi        | 普赛     | $\Psi \qquad \psi$         | 角速，介质电通量
24 | Ω      | \Omega   | ω      | \omega      | 欧米伽    | $\Omega \qquad \omega$     | 欧姆，角速度，交流电的电角度
异体 | E      | E        | ε      | \varepsilon | 异体伊普西隆 | $E \qquad \varepsilon$     |
异体 | K      | K        | ϰ      | \varkappa   | 异体卡帕   | $K \qquad \varkappa$       |
异体 | Θ      | \Theta   | ϑ      | \vartheta   | 异体西塔   | $\Theta \qquad \vartheta$  |
异体 | Π      | \Pi      | ϖ      | \varpi      | 异体派    | $P \qquad \varpi$          |
异体 | R      | R        | ϱ      | \varrho     | 异体柔    | $R \qquad \varrho$         |
异体 | Σ      | \Sigma   | ς      | \varsigma   | 异体西格玛  | $\Sigma \qquad \varsigma$  |
异体 | Φ      | \Phi     | φ      | \varphi     | 异体弗爱   | $\Phi \qquad \varphi$      |

### 字体：

语法      | 字体                | 例子                      | 效果
:------ | :---------------- | :---------------------- | --------------------------------------------:
\rm     | 罗马体               | \{\rm 你好，World，123}     |     {$\rm 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\mathrm | 罗马体               | \{\mathrm 你好，World，123} | {$\mathrm 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\bf     | 黑体                | \{\bf 你好，World，123}     |     {$\bf 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\Bbb    | 黑板粗体字             | \{\Bbb 你好，World，123}    |    {$\Bbb 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\mit    | 数学斜体              | \{\mit 你好，World，123}    |    {$\mit 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\scr    | 小体大写字母            | \{\scr 你好，World，123}    |    {$\scr 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\it     | 意大利体              | \{\it 你好，World，123}     |     {$\it 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\cal    | 花体                | \{\cal 你好，World，123}    |    {$\cal 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\sf     | 等线体               | \{\sf 你好，World，123}     |     {$\sf 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\tt     | 打字机字体             | \{\tt 你好，World，123}     |     {$\tt 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}
\frak   | Fraktur字母（一种德国字体） | \{\frak 你好，World，123}   |   {$\frak 你好，ABCDEFGHIJKLMNOPQRSTUVWXYZ，123$}

### 颜色：

代码                            | 效---------------果
:---------------------------- | :------------------------------
\color{black}{Hello World!}   | $\color{black}{Hello World!}$
\color{gray}{Hello World!}    | $\color{gray}{Hello World!}$
\color{silver}{Hello World!}  | $\color{silver}{Hello World!}$
\color{white}{Hello World!}   | $\color{white}{Hello World!}$
\color{maroon}{Hello World!}  | $\color{maroon}{Hello World!}$
\color{red}{Hello World!}     | $\color{red}{Hello World!}$
\color{yellow}{Hello World!}  | $\color{yellow}{Hello World!}$
\color{lime}{Hello World!}    | $\color{lime}{Hello World!}$
\color{olive}{Hello World!}   | $\color{olive}{Hello World!}$
\color{green}{Hello World!}   | $\color{green}{Hello World!}$
\color{teal}{Hello World!}    | $\color{teal}{Hello World!}$
\color{aqua}{Hello World!}    | $\color{aqua}{Hello World!}$
\color{blue}{Hello World!}    | $\color{blue}{Hello World!}$
\color{navy}{Hello World!}    | $\color{navy}{Hello World!}$
\color{purple}{Hello World!}  | $\color{purple}{Hello World!}$
\color{fuchsia}{Hello World!} | $\color{fuchsia}{Hello World!}$
