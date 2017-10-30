from collections import OrderedDict
import json

class Model:
    def __init__(self, json_str):
        self._json = json.loads(json_str, object_pairs_hook=OrderedDict)
        self._id = self._json["compositeModelName"]
        self._complexity = self._json["complexity"]
        self._frequency = 0


a = "{\"complexity\":70701480,\"elementaryModels\":[{\"modelName\":\"Match with 1 Toggled Case letter(s):JohnTheRipper\",\"complexity\":28360,\"modelIndex\":0},{\"modelName\":\"Exact Match:ptPopular\",\"complexity\":2493,\"modelIndex\":1}],\"compositeModelName\":\"|Match with 1 Toggled Case letter(s):JohnTheRipper|Exact Match:ptPopular|\"}"

m = Model(a)

b = 1

