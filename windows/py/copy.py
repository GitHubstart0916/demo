import os
import time
from collections import deque


def getDirAndCopyFile(sourcePath, targetPath):
    if not os.path.exists(sourcePath):
        return
    if not os.path.exists(targetPath):
        os.makedirs(targetPath)

    # 遍历文件夹
    for fileName in os.listdir(sourcePath):
        # 拼接原文件或者文件夹的绝对路径
        absourcePath = os.path.join(sourcePath, fileName)
        # 拼接目标文件或者文件加的绝对路径
        abstargetPath = os.path.join(targetPath, fileName)
        # 判断原文件的绝对路径是目录还是文件
        if os.path.isdir(absourcePath):
            # 是目录就创建相应的目标目录
            os.makedirs(abstargetPath)
            # 递归调用getDirAndCopyFile()函数
            getDirAndCopyFile(absourcePath, abstargetPath)
        # 是文件就进行复制
        if os.path.isfile(absourcePath):
            rbf = open(absourcePath, "rb")
            wbf = open(abstargetPath, "wb")
            while True:
                content = rbf.readline(1024 * 1024)
                if len(content) == 0:
                    break
                wbf.write(content)
                wbf.flush()
            rbf.close()
            wbf.close()


if __name__ == '__main__':
    # startTime = time.clock()
    # sourcePath = r"H:\培训资料"
    # targetPath = r"H:\培训资料_备份"
    # getDirAndCopyFile(sourcePath, targetPath)
    # # 时间是用来计算复制总共消耗了多少时间
    # endTime = time.clock()
    # time_mi = endTime // 60
    # time_s = endTime // 1 % 60
    # time_ms = ((endTime * 100) // 1) % 100
    # print("总用时:%02.0f:%02.0f:%2.0f" % (time_mi, time_s, time_ms))
    
  sourcePath = r"F:\SE\team08-proj/backend" 
  targetPath = r"F:\SE\demo" 
  getDirAndCopyFile(sourcePath, targetPath) 
