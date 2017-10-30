import jpype
from model_recognition.model import Model
import json
from collections import OrderedDict

class ModelRecognition:
    def __init__(self):
        classpath = "/home/bernardo/PycharmProjects/corinda/model_recognition/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"
        jpype.startJVM(jpype.getDefaultJVMPath(), "-Djava.class.path=%s" % classpath)
        passfaultPkg = jpype.JPackage('org').owasp.passfault
        TextAnalysis = passfaultPkg.TextAnalysis

        self._analyzer = TextAnalysis()


    def analyze(self, password):
        return self._analyzer.passwordAnalysis(password)

    def analyzer_thread(self, password_list, model_dict):
        jpype.attachThreadToJVM()
        for index, row in password_list.iterrows():
            password = row["password"]
            freq = row["freq"]

            json_data = self._analyzer.passwordAnalysis(password)
            id = json.loads(json_data, object_pairs_hook=OrderedDict)["compositeModelName"]

            if (id in model_dict):
                model_dict[id].acquire_lock()
                model_dict[id].add_frequency(freq)
                model_dict[id].release_lock()

            else:
                model_dict[id] = Model(json_data)
                model_dict[id].acquire_lock()
                model_dict[id].add_frequency(freq)
                model_dict[id].release_lock()


def main():
    recog = ModelRecognition()
    print(recog.analyze("bernArdorodrigues"))

if __name__ == "__main__":
    main()
