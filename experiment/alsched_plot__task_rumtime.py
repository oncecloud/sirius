# Machines = 5000
# Tasks ranges from 1000 to 5000
# 

import matplotlib.pyplot as plt
import numpy as np

x = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23]
kubernetes = [47, 37, 23, 22.5, 70, 80, 73, 53, 24, 43, 45, 56, 57,
              67, 54, 32, 54, 67, 89, 32, 12, 34, 56, 78]
sirius = [30, 45, 50, 55, 60, 65, 70, 80, 73, 53, 24, 43, 45, 67, 54, 32, 54, 67, 89, 32, 12, 34, 56, 78]

plt.plot(x, kubernetes, color="green", linestyle="-", marker='8', label="Kubernetes")
plt.plot(x, sirius, color="blue", linestyle="-", marker="^", label="Sirius")


plt.legend(loc='upper left')
plt.xlabel('Numbers of Tasks')
plt.ylabel('Mean Task Wait time(s)')

plt.show()
