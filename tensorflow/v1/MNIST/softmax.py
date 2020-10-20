# 导入 MNIST 数据
import input_data
mnist = input_data.read_data_sets("data/", one_hot=True)

import tensorflow as tf

# 变量，784 列数组，行不定
# 此函数可以理解为形参，用于定义过程，在执行的时候再赋具体的值
x = tf.compat.v1.placeholder(tf.float32, [None, 784])
W = tf.Variable(tf.zeros([784,10]))
b = tf.Variable(tf.zeros([10]))
# 我们把向量化后的图片 x 和权重矩阵 W 相乘，加上偏置 b，然后计算每个分类的 softmax 概率值
y = tf.nn.softmax(tf.matmul(x,W) + b)
y_ = tf.compat.v1.placeholder("float", [None,10])
# 可以很容易的为训练过程指定最小化误差用的损失函数，我们的损失函数是目标类别和预测类别之间的交叉熵
# tf.reduce_sum 把 minibatch 里的每张图片的交叉熵值都加起来了。我们计算的交叉熵是指整个 minibatch 的
cross_entropy = -tf.reduce_sum(y_*tf.math.log(y))
# 最速下降法让交叉熵下降，步长为 0.01
train_step = tf.compat.v1.train.GradientDescentOptimizer(0.01).minimize(cross_entropy)
# 初始化
init = tf.compat.v1.global_variables_initializer()
sess = tf.compat.v1.Session()
sess.run(init)
# 迭代训练
for i in range(1000):
  # 每一步迭代，我们都会加载 100 个训练样本
  batch_xs, batch_ys = mnist.train.next_batch(100)
  # 执行一次 train_step，并通过 feed_dict 将 x 和 y_ 张量占位符用训练训练数据替代
  sess.run(train_step, feed_dict={x: batch_xs, y_: batch_ys})

# tf.argmax 是一个非常有用的函数，它能给出某个 tensor 对象在某一维上的其数据最大值所在的索引值
# tf.argmax(y,1) 返回的是模型对于任一输入 x 预测到的标签值
# tf.argmax(y_,1) 代表正确的标签
correct_prediction = tf.equal(tf.argmax(y,1), tf.argmax(y_,1))
# 布尔值转换为浮点数来代表对、错，然后取平均值
# [True, False, True, True] 变为 [1,0,1,1]，计算出平均值为 0.75
accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))
# 评价大概 91%
print (sess.run(accuracy, feed_dict={x: mnist.test.images, y_: mnist.test.labels}))