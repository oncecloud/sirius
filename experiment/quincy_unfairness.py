# 
# Machines range from 0 to 5000
#

import matplotlib.pyplot as plt
import numpy as np

group = 3

mix = (0.65, 0.34, 0.32)
batch = (0.52, 0.28, 0.30)
service = (0.67, 0.41, 0.38)


# fig, ax = plt.subplots()    
index = np.arange(group)    
bar_width = 0.23   
opacity = 0.8  

plt.bar(index, mix, bar_width, alpha=opacity, color="blue", label="Mix workloads")
plt.bar(index + bar_width, batch, bar_width, alpha=opacity, color="green", label="Batch workloads")
plt.bar(index + bar_width * 2, service, bar_width, alpha=opacity, color="red", label="Service workloads")

plt.legend(loc='upper left')
plt.ylabel('Unfairness')
plt.xticks(index + bar_width, ('Random', 'Firmament/quincy','Sirius/quincy'), ha="center")  
plt.ylim(0, 1)
plt.show()
