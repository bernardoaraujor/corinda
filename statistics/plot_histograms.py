import matplotlib.pyplot as plt
import numpy as np

#def hist_complexity(list):
#    f = open(list + '_comp.txt', 'r')
#    lines = f.readlines()
#    f.close()

#    complexities = []
#    sum = 0
#    i = 0
#    for line in lines:
#        complexity = float(line.replace('\n', ''))
#        complexities.append(complexity)
#        sum += complexity
#        i += 1

#    avgComplexity = sum / i
#    print(list + ' average complexity (log10): ' + str(np.log10(avgComplexity)))

#    MIN, MAX = np.power(10, 0), np.power(10, 15)
#    plt.hist(complexities, bins=10 ** np.linspace(np.log10(MIN), np.log10(MAX), 200), facecolor='gray')
#    plt.axvline(avgComplexity, color='deeppink', linestyle='dashed', linewidth=4, label="complexidade media")
#    plt.gca().set_xscale("log")
#    plt.xlabel('Complexidade')
#    plt.ylabel('Frequencia')
#    plt.title('histograma de complexidades: ' + list)
#    plt.legend()
#    plt.grid(True)
#    plt.show()

#def hist_entropy(list):
#    f = open(list + '_entr.txt', 'r')
#    lines = f.readlines()
#    f.close()

#    entropies = []
#    sum = 0
#    i = 0
#    for line in lines:
#        entropy = float(line.replace('\n', ''))
#        sum += entropy
#        i += 1
#        entropies.append(entropy)

#    avgEntropy = sum / i

#    plt.hist(entropies, 200, facecolor='gray')
#    plt.xlabel('Entropia')
#    plt.ylabel('Frequencia')
#    plt.axvline(avgEntropy, color='deeppink', linestyle='dashed', linewidth=4, label="entropia media")
#    plt.title('histograma de entropias: ' + list)
#    plt.legend()
#    plt.grid(True)
#    plt.show()

def hist_prob(list):
    f = open(list + '_prob.txt', 'r')
    lines = f.readlines()
    f.close()

    probabilities = []
    sum = 0
    for line in lines:
        prob = float(line.replace('\n', ''))
        sum += prob
        probabilities.append(prob)

    probabilities.sort(reverse=True)

    sum80 = 0
    for prob in probabilities:
        sum80 += prob
        if sum80 > 0.8*sum:
            maxProb = prob
            print(list + ' 80% probability: ' + str(maxProb))
            break

    plt.hist(probabilities, 1000, facecolor='gray')
    plt.xlim(0, 0.01)
    plt.xlabel('Probabilidade')
    plt.ylabel('Frequencia')
    plt.yscale('log', nonposy='clip')
    plt.axvline(maxProb, color='deeppink', linestyle='dashed', linewidth=4, label="80% das ocorrencias")
    plt.title('histograma de probabilidades: ' + list)
    plt.legend()
    plt.grid(True)
    plt.show()

hist_prob('rockyou')
#hist_entropy('rockyou')
#hist_complexity('rockyou')

hist_prob('rockyou_1M')
#hist_entropy('rockyou_1M')
#hist_complexity('rockyou_1M')

hist_prob('linkedin')
#hist_entropy('linkedin')
#hist_complexity('linkedin')

hist_prob('linkedin_1M')
#hist_entropy('linkedin_1M')
#hist_complexity('linkedin_1M')

hist_prob('antipublic')
#hist_entropy('antipublic')
#hist_complexity('antipublic')

hist_prob('antipublic_1M')
#hist_entropy('antipublic_1M')
#hist_complexity('antipublic_1M')