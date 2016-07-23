# Tasks = 3000
# Machines range from 0 to 5000
# 

import matplotlib.pyplot as plt
import numpy as np

x = [0, 1000, 2000, 3000, 4000, 5000]
ics = [0, 85.3, 125.2 ,178.6, 201, 225]
quincy = [0, 200.1, 400 ,632.1, 800, 1001]
random = [0, 60, 80, 132, 160, 189]

plt.plot(x, random, color="green",  linestyle="-", label="Random")
plt.plot(x, quincy, color="blue",  linestyle="-", label="Quincy")
plt.plot(x, ics, color="red",  linestyle="-", label="Sirius/ICS")


plt.legend(loc='upper left')
plt.xlabel('Numbers of machine')
plt.ylabel('Runtime(ms)')

plt.show()