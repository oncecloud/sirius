# Machines = 5000
# Tasks ranges from 1000 to 5000
# 

import matplotlib.pyplot as plt
import numpy as np

x=[10,40,60,80,100]
random = [10,10,3,8.77]
alsched = [5, 5, 2.64,5.22, 7.24]
sirius_alsched = [5, 1.2, 2.22, 3.45, 4.23]

plt.plot(x, random, color="green",  linestyle="-",marker='8', label="Random")
plt.plot(x, alsched, color="blue",  linestyle="-",marker="^", label="Alsched")
plt.plot(x, sirius_alsched, color="red",  linestyle="-", marker="s",label="Sirius/alsched")


plt.legend(loc='upper left')
plt.xlabel('Numbers of Tasks')
plt.ylabel('Mean Task Wait time(s)')

plt.show()