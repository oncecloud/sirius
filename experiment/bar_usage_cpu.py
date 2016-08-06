# Tasks = 240, machine=10
# cpu and mem usage of sirius,swarm,kubernetes

import matplotlib.pyplot as plt
import numpy as np

# plt.axhline(xmax=50)
plt.ylim(0, 100)

group = 4

random = (15.21, 17.32, 20.33, 23.88)
sirius = (31,32.22, 34.34, 39.68)


# fig, ax = plt.subplots()    
index = np.arange(group)    
bar_width = 0.32    
opacity = 0.8    

plt.bar(index, random, bar_width, alpha=opacity, color="blue", label="Random")
plt.bar(index + bar_width, sirius, bar_width, alpha=opacity, color="red", label="Sirius")


plt.legend(loc='upper right')
plt.ylabel('CPU Usage(%)')
#plt.xticks()
plt.xticks(index + bar_width, ('(T=10000,M=5000)','(T=20000,M=5000)','(T=20000,M=10000)','(T=40000, M=10000)' ),rotation=17 )

plt.show()
