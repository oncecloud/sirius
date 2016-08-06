"""Examples illustrating the use of plt.subplots().

This function creates a figure and a grid of subplots with a single call, while
providing reasonable control over how the individual plots are created.  For
very refined tuning of subplot creation, you can still use add_subplot()
directly on a new figure.
"""

import matplotlib.pyplot as plt
import numpy as np

# Simple data to display in various forms
x = np.linspace(0, 2 * np.pi, 400)
y = np.sin(x ** 2)
group = 4
index = np.arange(group)    
bar_width = 0.32    
opacity = 0.8  


random_cpu = (15.21, 24.21, 27.33, 39.28)
sirius_cpu = (27, 41, 46.22, 60.22)


plt.close('all')

# Two subplots, the axes array is 1-d
f, axarr = plt.subplots(2, sharex=True)
axarr[0].bar(index, random_cpu, bar_width, alpha=opacity, color="blue", label="Random")
axarr[0].bar(index + bar_width, sirius_cpu, bar_width, alpha=opacity, color="red", label="Sirius")
plt.ylabel('CPU Usage(%)')
axarr[1].bar(index, random_cpu, bar_width, alpha=opacity, color="blue", label="Random")
axarr[1].bar(index + bar_width, sirius_cpu, bar_width, alpha=opacity, color="red", label="Sirius")
plt.ylabel('CPU Usage(%)')
plt.xticks(index + bar_width, ('(T=10000, M=5000)','(T=20000, M=5000)','(T=20000, M=10000)','(T=40000, M=10000)' ) )








plt.show()