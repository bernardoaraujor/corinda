import json
import gzip
import matplotlib.pyplot as plt
import numpy as np
import csv

class Elementary:
    def __init__(self, name, entropy, complexity):
        self.name = name
        self.entropy = entropy
        self.complexity = complexity

    def getName(self):
        return self.name

    def getEntropy(self):
        return self.entropy

    def getComplexity(self):
        return self.complexity

class Composite:
    def __init__(self, name, entropy, prob, complexity):
        self.name = name
        self.entropy = entropy
        self.probability = prob
        self.complexity = complexity

    def getName(self):
        return self.name

    def getEntropy(self):
        return self.entropy

    def getProbability(self):
        return self.probability

    def getComplexity(self):
        return self.complexity

class Trained:
    def __init__(self, name):
        self.name = name
        self.elementaries = self.loadElementaries()
        self.composites = self.loadComposites()

    def loadElementaries(self):
        with gzip.open("maps/" + self.name + "Elementaries.json.gz", "rb") as f:
            list = json.loads(f.read().decode())

        elementaries = dict()
        for line in list:
            name = line["Name"]
            entropy = line["Entropy"]
            complexity = len(line['TokenFreqs'])
            elementaries[name] = Elementary(name, entropy, complexity)

        return elementaries

    def loadComposites(self):
        with gzip.open("maps/" + self.name + "Composites.json.gz", "rb") as f:
            list = json.loads(f.read().decode())

        composites = dict()
        for line in list:
            name = line["Name"]
            entropy = line["Entropy"]
            prob = line["Prob"]
            elementaries = line["Models"]
            complexity = 1
            for elementary in elementaries:
                complexity *= self.elementaries[elementary].complexity
            composites[name] = Composite(name, entropy, prob, complexity)

        return composites

    def getElementaryEntropies(self):
        entropies = np.zeros((len(self.elementaries), 1))
        i = 0
        for elementary in self.elementaries:
            entropies[i] = self.elementaries[elementary].getEntropy()
            i = i + 1

        return entropies

    def getCompositeEntropies(self):
        entropies = np.zeros((len(self.composites), 1))
        i = 0
        for composite in self.composites:
            entropies[i] = self.composites[composite].getEntropy()
            i = i + 1

        return entropies

    def getElementaryComplexities(self):
        complexities = np.zeros((len(self.elementaries), 1))
        i = 0
        for elementary in self.elementaries:
            complexities[i] = self.elementaries[elementary].getComplexity()
            i = i + 1

        return complexities

    def getCompositeComplexities(self):
        complexities = np.zeros((len(self.composites), 1))
        i = 0
        for composite in self.composites:
            complexities[i] = self.composites[composite].getComplexity()
            i = i + 1

        return complexities

    def getCompositeProbabilities(self):
        probabilities = np.zeros((len(self.composites), 1))
        i = 0
        for composite in self.composites:
            probabilities[i] = self.composites[composite].getProbability()
            i = i + 1

        return probabilities

    def histCompositeEntropy(self):
        plt.hist(self.getCompositeEntropies(), bins='auto')
        plt.title(self.name)
        plt.xlabel("H($\hat m$(s))")
        plt.ylabel("Número de ocorrências em $\Gamma$")
        plt.savefig("output/" + self.name + "_composite_entropy")
        plt.clf()

    def histElementaryEntropy(self):
        bins = np.arange(0, self.getElementaryEntropies().max(), 0.1)
        plt.hist(self.getElementaryEntropies(), bins=bins)
        plt.title(self.name)
        plt.xlabel("H(m(s))")
        plt.ylabel("Número de ocorrências em $\Gamma$")
        plt.savefig("output/" + self.name + "_elementary_entropy")
        plt.clf()

    def histCompositeProbability(self):
        probabilities = self.getCompositeProbabilities()
        bins = 10**np.arange(-5, 0, 0.1)
       
        plt.hist(probabilities, bins=bins, log=True)
        plt.title(self.name)
        plt.xscale('log')
        plt.xlabel("$\Theta$($\hat m$(s))")
        plt.ylabel("Número de ocorrências em $\Gamma$")
        plt.savefig("output/" + self.name + "_composite_probability")
        plt.clf()

    def frequentModelsList(self):
        probabilities = self.getCompositeProbabilities().tolist()
        probabilities.sort(reverse=True)

        topProbabilities = probabilities[0:19]

        f = open("output/" + self.name + "_frequent_models.csv", 'w')
        for prob in topProbabilities:
            prob = prob[0]
            for _, composite in self.composites.items():
                if prob == composite.getProbability():
                    f.write(str(prob) + ", " + composite.getName() + "\n")
        f.close()

    def genOutput(self):
        self.histCompositeProbability()
        self.histCompositeEntropy()
        self.histElementaryEntropy()
        self.frequentModelsList()

rockyou = Trained("rockyou")
rockyou.genOutput()
del rockyou

rockyou_1M = Trained("rockyou_1M")
rockyou_1M.genOutput()
del rockyou_1M

linkedin = Trained("linkedin")
linkedin.genOutput()
del linkedin

linkedin_1M = Trained("linkedin_1M")
linkedin_1M.genOutput()
del linkedin_1M

antipublic = Trained("antipublic")
antipublic.genOutput()
del antipublic

antipublic_1M = Trained("antipublic_1M")
antipublic_1M.genOutput()
del antipublic_1M
