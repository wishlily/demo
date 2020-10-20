import tensorflow as tf
# 载入并准备好 MNIST 数据集
mnist = tf.keras.datasets.mnist
# 将样本从整数转换为浮点数
(x_train, y_train), (x_test, y_test) = mnist.load_data()
x_train, x_test = x_train / 255.0, x_test / 255.0
# 将模型的各层堆叠起来，以搭建 tf.keras.Sequential 模型
model = tf.keras.models.Sequential([
  tf.keras.layers.Flatten(input_shape=(28, 28)),
  tf.keras.layers.Dense(128, activation='relu'),
  tf.keras.layers.Dropout(0.2),
  tf.keras.layers.Dense(10, activation='softmax')
])
# 为训练选择优化器和损失函数
model.compile(optimizer='adam',
              loss='sparse_categorical_crossentropy',
              metrics=['accuracy'])
# 训练
model.fit(x_train, y_train, epochs=5)
# 验证
model.evaluate(x_test, y_test, verbose=2)
'''
Epoch 1/5
60000/60000 [==============================] - 5s 77us/sample - loss: 0.2960 - acc: 0.9133
Epoch 2/5
60000/60000 [==============================] - 5s 85us/sample - loss: 0.1431 - acc: 0.9577
Epoch 3/5
60000/60000 [==============================] - 5s 87us/sample - loss: 0.1071 - acc: 0.9682
Epoch 4/5
60000/60000 [==============================] - 5s 84us/sample - loss: 0.0883 - acc: 0.9728
Epoch 5/5
60000/60000 [==============================] - 5s 86us/sample - loss: 0.0761 - acc: 0.9764
10000/10000 - 0s - loss: 0.0756 - acc: 0.9755
'''