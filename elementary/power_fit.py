import powerlaw
import matplotlib.pyplot as plt
import math

complexity = 10000

f = open("/home/bernardo/go/src/github.com/bernardoaraujor/corinda/elementary/export.txt", 'r')
lines = f.readlines()
f.close()

freq = []
freq_i = []
i = 0
for line in lines:
    n= int(line.replace('\n', ''))
    freq.append(n)
    freq_i.append(i)
    i = i+1

# estimate alpha
fit = powerlaw.Fit(freq, discrete=True)
alpha = fit.power_law.alpha

# generate power law
power = []
power_i = range(0, complexity)
for i in power_i:
    power.append(math.pow(i + 1, -alpha))

# find power law constant
last_freq = freq[-1]
intersect_power = power[len(freq)]  #where the lines meet
c = last_freq/intersect_power

# multiply power law by constant
#power = [c*p for p in power]

plt.loglog(freq_i, freq)
plt.loglog(power_i, power)
plt.show()