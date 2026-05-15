package = "voxgig-sdk-poetrydb"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/poetrydb-sdk.git"
}
description = {
  summary = "Poetrydb SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["poetrydb_sdk"] = "poetrydb_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
