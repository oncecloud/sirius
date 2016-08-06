import random
import matplotlib.pyplot as plt

def haha():
    x=[]
    y=[]
    y1=[]
    for i in range(1,100):
        x.append(i)
        y.append(0.4+random.random())
        y1.append(1.5+3.0*random.random())
    y.sort()
    y1.sort()
    plt.plot(x,y)
    plt.plot(x,y1)
    plt.ylim(0,10)
    plt.show()
haha()

    