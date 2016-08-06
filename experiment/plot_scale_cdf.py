# Tasks = 30000
# Machines range from 0 to 5000
# 

import matplotlib.pyplot as plt
import numpy as np

x = [0,2,4,6,8,10]
ics = [0.44, 0.5 ,0.533, 0.6, 0.68,0.78, 0.86]
quincy = [1.32, 1.78, 2.05, 2.63, 3.5 ,4.4, 5.8]
random = [0.24, 0.33, 0.44, 0.69, 0.89, 1.13,1.3]

plt.plot(x, random, color="green",  linestyle="-",marker="8", label="Random")
plt.plot(x, quincy, color="blue",  linestyle="-", marker="^",label="Sirius/CS")
plt.plot(x, ics, color="red",  linestyle="-", marker="s",label="Sirius/ICS")


plt.legend(loc='upper left')
plt.xlabel('Numbers of task(Machine=3000)')
plt.ylabel('Runtime(s)')
plt.show()