import jpype
from model_recognition.model import Model

class ModelRecognition:
    def __init__(self):
        classpath = "/home/bernardo/PycharmProjects/corinda/model_recognition/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"
        jpype.startJVM(jpype.getDefaultJVMPath(), "-Djava.class.path=%s" % classpath)
        passfaultPkg = jpype.JPackage('org').owasp.passfault
        TextAnalysis = passfaultPkg.TextAnalysis

        self._analyzer = TextAnalysis()


    def analyze(self, password):
        return self._analyzer.passwordAnalysis(password)

    def analyzeThread(self, password, vals):
        jpype.attachThreadToJVM()
        vals.append( self._analyzer.passwordAnalysis(password))

def main():
    recog = ModelRecognition()
    print(recog.analyze("bernArdorodrigues"))

if __name__ == "__main__":
    main()
