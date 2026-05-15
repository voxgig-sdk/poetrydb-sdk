# Poetrydb SDK feature factory

from feature.base_feature import PoetrydbBaseFeature
from feature.test_feature import PoetrydbTestFeature


def _make_feature(name):
    features = {
        "base": lambda: PoetrydbBaseFeature(),
        "test": lambda: PoetrydbTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
