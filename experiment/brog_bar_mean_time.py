# Tasks = 5000
# Machines range from 0 to 5000
#

import matplotlib.pyplot as plt
import numpy as np

group = 3
random = ()
brog = (9.8,11.8,4.73)
sirius = (10.0,11.6,4.22)

threshold=1.0
   
index = np.arange(group)    
bar_width = 0.33 
opacity  = 0.6    

plt.bar(index, brog, bar_width, alpha=opacity , color="blue",label="Brog")
plt.bar(index+bar_width, sirius, bar_width,alpha=opacity, color="red", label="Sirius/Brog")
plt.plot([0., 3], [threshold, threshold], "k--", color="black", linewidth="1.5")

plt.legend(loc='upper right')
plt.ylabel('Mean Task Runtime(s)')
plt.xticks(index + bar_width, ('MergeSort(p=5)', 'PageRank(p=5)', 'TPC-H(p=5)'))  
plt.ylim(0,15)
plt.show()