from collections import OrderedDict
import json
import threading

class Model:
    def __init__(self, json_str):
        self._json = json.loads(json_str, object_pairs_hook=OrderedDict)
        self._id = self._json["compositeModelName"]
        self._complexity = self._json["complexity"]
        self._frequency = 0
        self._lock = threading.Lock()

    def get_json(self):
        return self._json

    def get_id(self):
        return self._id

    def get_complexity(self):
        return self._complexity

    def get_frequency(self):
        return self._frequency

    def add_frequency(self, freq):
        self._frequency += freq

    def acquire_lock(self):
        self._lock.acquire()

    def release_lock(self):
        self._lock.release()
#a = "{\"complexity\":70701480,\"elementaryModels\":[{\"modelName\":\"Match with 1 Toggled Case letter(s):JohnTheRipper\",\"complexity\":28360,\"modelIndex\":0},{\"modelName\":\"Exact Match:ptPopular\",\"complexity\":2493,\"modelIndex\":1}],\"compositeModelName\":\"|Match with 1 Toggled Case letter(s):JohnTheRipper|Exact Match:ptPopular|\"}"

#m = Model(a)

#b = 1

