import pandas as pd

class Result:
    def __init__(self, trained, target, shaV):
        csvPath = "csv/" + target + ".csv.gz"
        resultsPath = "results/" + trained + "_" + target + "_" + shaV + ".csv.gz"

        csv = pd.read_csv(csvPath, compression='gzip', error_bad_lines=False, names=["freq", "pass"])
        results = pd.read_csv(resultsPath, compression='gzip', error_bad_lines=False, names=["pass", "hash"])

        self.original = dict(zip(list(csv["pass"]), list(csv["freq"])))
        self.results = dict(zip(list(results["pass"]), list(results["hash"])))

        self.name = trained + "_" + target + "_" + shaV

    def getFraction(self):
        originalN = len(self.original)
        resultsN = len(self.results)

        return float(resultsN)/originalN

    def totalFreqs(self):
        sumResults = 0
        for p, _ in self.results.items():
            if p in self.original:
                freq = int(self.original[p])
                sumResults += freq

        sumOriginal = 0
        for _, freq in self.original.items():
            sumOriginal += int(freq)

        return float(sumResults)/sumOriginal

    def genResults(self):
        print("analyzing " + self.name)
        f = open("output/results/" + self.name + ".csv", 'w')
        f.write("simple: " + str(self.getFraction()) + "\n")
        f.write("freqs: " + str(self.totalFreqs()) + "\n")
        f.close()

r = Result("rockyou_1M", "linkedin", "sha1")
r.genResults()
del r

r = Result("rockyou_1M", "linkedin", "sha256")
r.genResults()
del r

r = Result("rockyou_1M", "rockyou", "sha1")
r.genResults()
del r

r = Result("rockyou_1M", "rockyou", "sha256")
r.genResults()
del r

r = Result("rockyou", "antipublic", "sha1")
r.genResults()
del r

r = Result("rockyou", "antipublic", "sha256")
r.genResults()
del r

r = Result("rockyou", "linkedin", "sha1")
r.genResults()
del r

r = Result("rockyou", "linkedin", "sha256")
r.genResults()
del r

r = Result("rockyou", "rockyou", "sha1")
r.genResults()
del r

r = Result("rockyou", "rockyou", "sha256")
r.genResults()
del r