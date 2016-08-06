# Tasks = 240, machine=10
# cpu and mem usage of sirius,swarm,kubernetes

import matplotlib.pyplot as plt
import numpy as np

# plt.axhline(xmax=50)
plt.ylim(0, 100)

group = 4

random = (10.21, 17.43, 25.33, 40.26)
sirius = (19.88, 34.59, 46.22, 79.98)


# fig, ax = plt.subplots()    
index = np.arange(group)    
bar_width = 0.32    
opacity = 0.8    

plt.bar(index, random, bar_width, alpha=opacity, color="blue", label="Random")
plt.bar(index + bar_width, sirius, bar_width, alpha=opacity, color="red", label="Sirius")


plt.legend(loc='upper right')
plt.ylabel('Mem Usage(%)')

plt.xticks(index + bar_width, ('(T=10000,M=5000)','(T=20000,M=5000)','(T=20000,M=10000)','(T=40000, M=10000)' ),rotation=17  )

plt.show()
