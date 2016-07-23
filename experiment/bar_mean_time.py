# Tasks = 5000
# Machines range from 0 to 5000
#

import matplotlib.pyplot as plt
import numpy as np

group = 3

brog = (98,298,10)
sirius = (100,286,2)


#fig, ax = plt.subplots()    
index = np.arange(group)    
bar_width = 0.2    
#opacity   = 0.4    

plt.bar(index, brog, bar_width,  color="blue",label="Kubernetes")
plt.bar(index+bar_width, sirius, bar_width,color="red", label="Sirius/Kubernetes")

plt.legend(loc='upper right')
plt.ylabel('Mean Task Runtime(s)')
plt.xticks(index + bar_width, ('MergeSort', 'PageRank', 'TPC-H'))  

plt.show()