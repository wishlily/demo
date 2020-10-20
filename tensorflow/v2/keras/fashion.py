# TensorFlow and tf.keras
import tensorflow as tf
from tensorflow import keras

# Helper libraries
import numpy as np
import matplotlib.pyplot as plt

fashion_mnist = keras.datasets.fashion_mnist
(train_images, train_labels), (test_images, test_labels) = fashion_mnist.load_data()

class_names = ['T-shirt/top', 'Trouser', 'Pullover', 'Dress', 'Coat',
               'Sandal', 'Shirt', 'Sneaker', 'Bag', 'Ankle boot']

# test --- show one pic.
# plt.figure()
# plt.imshow(train_images[0])
# plt.colorbar()
# plt.grid(False)
# plt.show()

# convert float
train_images = train_images / 255.0
test_images = test_images / 255.0

# test --- show 25 pic & class
# plt.figure(figsize=(10,10))
# for i in range(25):
#     plt.subplot(5,5,i+1)
#     plt.xticks([])
#     plt.yticks([])
#     plt.grid(False)
#     plt.imshow(train_images[i], cmap=plt.cm.binary)
#     plt.xlabel(class_names[train_labels[i]])
# plt.show()

model = keras.Sequential([
    keras.layers.Flatten(input_shape=(28, 28)), # 将图像格式从二维数组（28 x 28 像素）转换成一维数组（28 x 28 = 784 像素）
    keras.layers.Dense(128, activation='relu'), # 密集连接或全连接神经层，第一个 Dense 层有 128 个节点
    keras.layers.Dense(10)                      # 最后一层会返回一个长度为 10 的 logits 数组
])

model.compile(optimizer='adam',                                                     # 优化器
              loss=tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True), # 损失函数
              metrics=['accuracy'])                                                 # 指标，本例使用了准确率，即被正确分类的图像的比率

model.fit(train_images, train_labels, epochs=10)

test_loss, test_acc = model.evaluate(test_images,  test_labels, verbose=2)
print('\nTest accuracy:', test_acc)