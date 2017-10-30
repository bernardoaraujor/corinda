import json
import threading
from collections import OrderedDict

import jpype

from model_recognition.model import Model


class ModelRecognition:
    def __init__(self):
        classpath = "/home/bernardo/PycharmProjects/corinda/model_recognition/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"
        jpype.startJVM(jpype.getDefaultJVMPath(), "-Djava.class.path=%s" % classpath)
        passfaultPkg = jpype.JPackage('org').owasp.passfault
        TextAnalysis = passfaultPkg.TextAnalysis

        self._analyzer = TextAnalysis()
        self._count = 0
        self._count_lock= threading.Lock()


    def analyze(self, password):
        return self._analyzer.passwordAnalysis(password)

    def increase_count(self):
        self._count += 1

    def get_count(self):
        return self._count

    def acquire_count_lock(self):
        self._count_lock.acquire()

    def release_count_lock(self):
        self._count_lock.release()

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

            self.acquire_count_lock()
            self.increase_count()
            self.release_count_lock()




def main():
    recog = ModelRecognition()
    print(recog.analyze("bernArdorodrigues"))

if __name__ == "__main__":
    main()
