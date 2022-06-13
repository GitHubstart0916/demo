import string

from PIL import Image
import os
from glob import glob
import sys

args = list(map(str, sys.argv))
print(args)

input_path = "/home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps/"
input_path += args[1]
print(args[0], args[1])
print(input_path)
path = input_path
print(path)
# 获取指定目录下的所有pgm文件列表
img_list = glob(os.path.join(path, "*.pgm"))
print(img_list)
# 遍历列表，重新保存为jpg格式
for img in img_list:
    filename = img[:-7] + args[1] + ".jpg"
    print(img)
    Image.open(img).save(filename)

    img = Image.open(filename)  ## 打开chess.png文件，并赋值给img
    region = img.crop((1850, 1850, 2150, 2150))  ## 0,0表示要裁剪的位置的左上角坐标，50,50表示右下角。
    region.save(filename)  ## 将裁剪下来的图片保存到 举例.png

