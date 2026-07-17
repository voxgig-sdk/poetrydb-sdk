-- Poetrydb SDK exists test

local sdk = require("poetrydb_sdk")

describe("PoetrydbSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
