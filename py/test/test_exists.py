# ProjectName SDK exists test

import pytest
from poetrydb_sdk import PoetrydbSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = PoetrydbSDK.test(None, None)
        assert testsdk is not None
