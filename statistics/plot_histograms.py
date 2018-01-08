import matplotlib.pyplot as plt

def hist_entropy(list):
    f = open(list + '_entr.txt', 'r')
    lines = f.readlines()
    f.close()

    entropies = []
    sum = 0
    i = 0
    for line in lines:
        entropy = float(line.replace('\n', ''))
        sum += entropy
        i += 1
        entropies.append(entropy)

    avgEntropy = sum / i

    plt.hist(entropies, 100, facecolor='lightgray')
    plt.xlabel('Entropia')
    plt.ylabel('Frequencia')
    plt.axvline(avgEntropy, color='deeppink', linestyle='dashed', linewidth=4, label="entropia media")
    plt.title('histograma de entropias: ' + list)
    plt.legend()
    plt.grid(True)
    plt.show()

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

    plt.hist(probabilities, 500, facecolor='lightgray')
    plt.xlim(0, 0.01)
    plt.xlabel('Probabilidade')
    plt.ylabel('Frequencia')
    plt.yscale('log', nonposy='clip')
    plt.axvline(maxProb, color='deeppink', linestyle='dashed', linewidth=4, label="probabilidade media")
    plt.title('histograma de probabilidades: ' + list)
    plt.legend()
    plt.grid(True)
    plt.show()

hist_prob('rockyou')
hist_entropy('rockyou')

hist_prob('linkedin')
hist_entropy('linkedin')

hist_prob('antipublic')
hist_entropy('antipublic')


