# Tasks = 30000
# Machines range from 0 to 5000
# 

import matplotlib.pyplot as plt
import numpy as np

x = [10000, 12000, 14000, 16000, 18000, 20000]
ics = [0.68 ,0.733, 0.8, 0.88,0.98, 1.16]
quincy = [4.00, 4.53, 5.53, 6.83 ,8.27, 9.75]
random = [0.45, 0.53, 0.65, 1.00, 1.43,1.97]

plt.plot(x, random, color="green",  linestyle="-",marker="8", label="Random")
plt.plot(x, quincy, color="blue",  linestyle="-", marker="^",label="Sirius/CS")
plt.plot(x, ics, color="red",  linestyle="-", marker="s",label="Sirius/ICS")


plt.legend(loc='upper left')
plt.xlabel('Numbers of task(Machine=10000)')
plt.ylabel('Runtime(s)')

plt.show()