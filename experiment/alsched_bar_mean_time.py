# 
# Machines range from 0 to 5000
#

import matplotlib.pyplot as plt
import numpy as np

group = 2

random = (8.75, 10.77)
alsched = (4.96, 9.73)
sirius_alsched = (4.83, 7.25)


# fig, ax = plt.subplots()    
index = np.arange(group)    
bar_width = 0.23   
opacity = 0.6  

plt.bar(index, random, bar_width, alpha=opacity, color="blue", label="Random")
plt.bar(index + bar_width, alsched, bar_width, alpha=opacity, color="green", label="Swarm/alsched")
plt.bar(index + bar_width * 2, sirius_alsched,bar_width, alpha=opacity,  color="red", label="Sirius/alsched")

plt.legend(loc='upper left')
plt.ylabel('Mean Task Runtime(s)')
plt.xlabel("Task Numbers")
plt.xticks(index + bar_width, ('T=10', 'T=20'))  
plt.ylim(0,15)
plt.show()
